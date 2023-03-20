package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Account struct {
	ID       int       `sql:"id"`
	Owner    string    `sql:"owner"`
	Balance  int       `sql:"balance"`
	Currency string    `sql:"currency"`
	Time     time.Time `sql:"created_at"`
}

// const (
// 	DATABASE = "postgres"
// 	USERNAME = "root"
// 	PASSWORD = "vinay"
// 	ADDRESS  = "localhost"
// 	DATABASE = "simple_bank"
// 	SSL      = " sslmode=disable"
// )

func main() {
	// sslmode=disable
	conn := "postgres://postgres:vinay@localhost/simple_bank?sslmode=disable"
	// db, err := sql.Open("postgres", "postgresql://simple_bank:vinay@localhost:5432/validation?sslmode=disable")
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	query := "SELECT * FROM accounts Where Currency=$1 and Balance=$2 "
	rows, err := db.Query(query, "USD", 100)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var accounts []Account
	var accmap = make(map[int]Account)
	for rows.Next() {
		var account Account
		err := rows.Scan(&account.ID, &account.Owner, &account.Balance, &account.Currency, &account.Time)
		if err != nil {
			panic(err)
		}
		accmap[account.ID] = account
		accounts = append(accounts, account)
	}
	fmt.Println(accmap)
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	res, err2 := json.Marshal(&accounts)
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(string(res))
	// fmt.Println(accounts)
}
