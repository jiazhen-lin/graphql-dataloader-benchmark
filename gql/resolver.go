package gql

import (
	"context"
)

type Resolver struct{}

func (r *Resolver) Users(context context.Context) ([]*userResolver, error) {
	// get all user from db
	type user struct {
		ID   int    `json:"int"`
		Name string `json:"name"`
	}
	var us []*user
	if err := db.Select(&us, "SELECT id, name FROM TestUser"); err != nil {
		return nil, err
	}

	rs := []*userResolver{}
	for _, u := range us {
		rs = append(rs, &userResolver{
			id:   u.ID,
			name: u.Name,
		})
	}

	return rs, nil
}

type userResolver struct {
	id   int
	name string
}

func (r *userResolver) Name() string {
	return r.name
}
func (r *userResolver) Posts(context context.Context) ([]*postResolver, error) {
	ps, err := LoadPost(context, r.id)
	if err != nil {
		return nil, err
	}

	rs := []*postResolver{}
	for _, p := range ps {
		rs = append(rs, &postResolver{
			text: p.Text,
		})
	}
	return rs, nil
}

type postResolver struct {
	text string
}

func (r *postResolver) Text(context context.Context) string {
	return r.text
}
