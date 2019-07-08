package main

import (
	"fmt"
	"github.com/go-pg/pg"
)

func insert(value *PageData) {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "123456",
		Database: "mydb",
	})
	defer db.Close()

	err := db.RunInTransaction(func(tx *pg.Tx) error {
		fmt.Println("================")
		var n int
		_, err := tx.Query(pg.Scan(&n), "insert into company (company_number,nzbn) values (?company_number,?nzbn) RETURNING id", value)
		return err
	})
	if err != nil {
		panic(err)
	}
}
