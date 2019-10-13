package gql

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/graph-gophers/dataloader"
	"github.com/jmoiron/sqlx"
)

type loaderKey string

const (
	postLoaderKey loaderKey = "post"
)

var (
	connectCount    int
	connectTimeNano int64
)

func Analyze() {
	fmt.Println(fmt.Sprintf("Connect count: %d, timeNano: %d, ave time: %f",
		connectCount,
		connectTimeNano,
		float64(connectTimeNano)/float64(connectCount),
	))
}

func Attach(ctx context.Context, loaderEnable bool) context.Context {
	ctx = context.WithValue(ctx, "loaderEnable", loaderEnable)
	ctx = context.WithValue(ctx, postLoaderKey, dataloader.NewBatchedLoader(postBatchFunc))
	return ctx
}

func extract(ctx context.Context, key loaderKey) (*dataloader.Loader, error) {
	ld, ok := ctx.Value(key).(*dataloader.Loader)
	if !ok {
		return nil, fmt.Errorf("invalid key")
	}
	return ld, nil
}

type Post struct {
	UserID int    `db:"userID"`
	Text   string `db:"text"`
}

type userIDKey int

func (k userIDKey) String() string {
	return strconv.Itoa(int(k))
}
func (k userIDKey) Raw() interface{} { return k }

func LoadPost(ctx context.Context, userID int) ([]*Post, error) {
	// get loader from ctx
	ld, err := extract(ctx, postLoaderKey)
	if err != nil {
		return nil, err
	}

	// check enable
	enable, ok := ctx.Value("loaderEnable").(bool)
	if !enable || !ok {
		return getPostByUser(userID)
	}

	data, err := ld.Load(ctx, userIDKey(userID))()
	if err != nil {
		return nil, err
	}
	return data.([]*Post), nil
}

func postBatchFunc(context context.Context, keys dataloader.Keys) []*dataloader.Result {
	ids := []int{}
	for _, k := range keys {
		id := k.Raw().(userIDKey)
		ids = append(ids, int(id))
	}

	results := make([]*dataloader.Result, len(keys))
	ps, err := getPostByUsers(ids...)
	if err != nil {
		for i := range keys {
			results[i] = &dataloader.Result{Error: err}
		}
		return results
	}
	userPosts := map[int][]*Post{}
	for _, p := range ps {
		if _, ok := userPosts[p.UserID]; !ok {
			userPosts[p.UserID] = []*Post{}
		}
		userPosts[p.UserID] = append(userPosts[p.UserID], p)
	}
	for i, id := range ids {
		results[i] = &dataloader.Result{Data: userPosts[id]}
	}

	return results
}

func getPostByUser(userID int) ([]*Post, error) {
	ps, err := getPostByUsers(userID)
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func getPostByUsers(userIDs ...int) ([]*Post, error) {
	startTime := time.Now()
	var ps []*Post
	q, args, err := sqlx.In("SELECT userID, text FROM TestPost WHERE userID IN (?)", userIDs)
	if err != nil {
		return nil, err
	}
	if err := db.Select(&ps, q, args...); err != nil {
		fmt.Println("err: ", err)
		return nil, err
	}

	connectCount++
	endTime := time.Now()
	connectTimeNano += (endTime.UnixNano() - startTime.UnixNano())
	return ps, nil
}
