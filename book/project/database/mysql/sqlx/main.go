package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"fmt"
)

var (
	db *sqlx.DB
)

type Person struct {
	UserId		int    `db:"user_id"`
	Username	string `db:"username"`
	Sex			string `db:"sex"`
	Email		string `db:"email"`
}

func init() {
	var err error
	db, err = sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("open mysql failed , err : ", err)
		return
	}
}

func main() {
	// Insert
	res, err := db.Exec(`INSERT INTO person(username, sex, email) VALUES(?, ?, ?)`,
		"user1", "male", "user1@qq.com")
	if err != nil {
		fmt.Println("exec failed, err : ",err)
		return
	}
	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, err : ", err)
		return
	}
	fmt.Println("insert succ:", id)

	// Select
	persons := []Person{}
	err = db.Select(&persons, `SELECT user_id, username, sex, email FROM person WHERE user_id=?`, 1)
	if err != nil {
		fmt.Printf("select err: %s\n", err)
		return
	}
	fmt.Printf("select Person: %#v\n", persons)

	// Update
	_, err = db.Exec(`UPDATE person SET username=? WHERE user_id=?`, "user1231231", 3)

	// Delete
	_, err = db.Exec(`DELETE FROM person WHERE user_id=?`, 5)
}
