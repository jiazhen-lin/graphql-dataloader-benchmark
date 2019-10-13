package gql

import (
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql" // go mysql-driver
	"github.com/jmoiron/sqlx"
)

const (
	url  = "" // <account>:<password>@tcp(<ip>:<port>)
	name = ""
)

var db *sqlx.DB

func CreateTestData() {
	var err error
	db, err = sqlx.Connect("mysql", fmt.Sprintf("%s/%s", url, name))
	if err != nil {
		panic(err)
	}

	// drop table if exists
	if _, err = db.Exec(`DROP TABLE IF EXISTS TestUser;`); err != nil {
		panic(err)
	}
	if _, err = db.Exec(`DROP TABLE IF EXISTS TestPost;`); err != nil {
		panic(err)
	}

	// create table
	if _, err = db.Exec(`CREATE TABLE IF NOT EXISTS TestUser (
		id INT NOT NULL,
		name VARCHAR(10) NOT NULL,
		PRIMARY KEY (id)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`); err != nil {
		panic(err)
	}
	if _, err = db.Exec(`CREATE TABLE IF NOT EXISTS TestPost (
		userID INT NOT NULL,
		text VARCHAR(20) NOT NULL,
		KEY (userID)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`); err != nil {
		panic(err)
	}

	// insert user
	userCount := 500
	paramStr := []string{}
	paramValues := []interface{}{}
	for i := 0; i < userCount; i++ {
		paramStr = append(paramStr, "(?,?)")
		paramValues = append(paramValues, i, fmt.Sprintf("user-%d", i))
	}
	insertUsers := fmt.Sprintf("INSERT INTO TestUser (id, name) VALUES %s", strings.Join(paramStr, ","))
	// fmt.Println(paramValues)
	if _, err = db.Exec(insertUsers, paramValues...); err != nil {
		panic(err)
	}

	// insert posts for every users
	paramStr = []string{}
	paramValues = []interface{}{}
	for i := 0; i < userCount; i++ {
		for j := 0; j < 1; j++ {
			paramStr = append(paramStr, "(?,?)")
			paramValues = append(paramValues, i, fmt.Sprintf("text-%d-%d", i, j))
		}
	}
	insertPost := fmt.Sprintf("INSERT INTO TestPost (userID, text) VALUES %s", strings.Join(paramStr, ","))
	if _, err = db.Exec(insertPost, paramValues...); err != nil {
		panic(err)
	}

}
