// 使用官方 database/sql 包进行数据库操作
package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"strconv"
	"math/rand"
)

var (
	db *sql.DB
)

func main() {
	var err error
	db, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/test?charset=utf8mb4")
	// 支持下面几种DSN写法，具体看MySQL服务端配置，常见为第2种
	// user@unix(/path/to/socket)/dbname?charset=utf8
	// user:password@tcp(localhost:5555)/dbname?charset=utf8
	// user:password@/dbname
	// user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname
	if err != nil {
		panic(err)
	}
	defer db.Close()

	insertData("user1", "male", "1@qq.com")
	insertData("user2", "female", "2@qq.com")
	insertData("user3", "male", "3@qq.com")

	deleteData(2)
	updateData("user2", 1)
	queryData(1)

	transaction()
}

func insertData(username, sex, email string) {
	fmt.Println("insertData...")
	stmt, err := db.Prepare(`INSERT INTO person (username, sex, email) VALUES (?, ?, ?)`)
	if err != nil {
		fmt.Printf("prepare data err: %s\n", err)
		return
	}
	defer stmt.Close()

	res, err := stmt.Exec(username, sex, email)
	// 通过返回的res可以进一步查询本次插入数据影响的行数
	// RowsAffected和最后插入的Id(如果数据库支持查询最后插入Id)
	if err != nil {
		fmt.Printf("insert data error: %v\n",  err)
		return
	}
	if lastInsertId, err := res.LastInsertId(); err == nil {
		fmt.Println("LastInsertId:",  lastInsertId)
	}
	if rowAffected, err := res.RowsAffected(); err == nil {
		fmt.Println("RowsAffected:",  rowAffected)
	}
}

func deleteData(user_id int) {
	fmt.Println("deleteData...")
	stmt, err := db.Prepare(`DELETE FROM person WHERE user_id=?`)
	if err != nil {
		fmt.Printf("prepare data err: %s\n", err)
		return
	}
	defer stmt.Close()
	res, err := stmt.Exec(user_id)
	if err != nil {
		fmt.Printf("delete data error: %v\n",  err)
		return
	}
	if rowAffected, err := res.RowsAffected(); err == nil {
		fmt.Println("RowsAffected:",  rowAffected)
	}
}

func updateData(username string, user_id int) {
	fmt.Println("updateData...")
	stmt, err := db.Prepare(`UPDATE person SET username=? WHERE user_id=?`)
	if err != nil {
		fmt.Printf("prepare data err: %s\n", err)
		return
	}
	defer stmt.Close()
	res, err := stmt.Exec(username, user_id)
	if err != nil {
		fmt.Printf("update data error: %v\n",  err)
		return
	}
	if rowsAffected,  err := res.RowsAffected(); nil == err {
		fmt.Println("RowsAffected:",  rowsAffected)
	}
}

func queryData(user_id int) {
	stmt, err := db.Prepare(`SELECT * FROM person WHERE user_id=?`)
	if err != nil {
		fmt.Printf("prepare data err: %s\n", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(user_id)
	if err != nil {
		fmt.Printf("query data error: %v\n",  err)
		return
	}
	defer rows.Close()

	// 构造scanArgs、values两个slice，
	// scanArgs的每个值指向values相应值的地址
	columns,  _ := rows.Columns()
	fmt.Println(columns)
	rowMaps := make([]map[string]string,  9)
	values := make([]sql.RawBytes,  len(columns))
	scans := make([]interface{},  len(columns))
	for i := range values {
		scans[i] = &values[i]
		scans[i] = &values[i]
	}
	i := 0
	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scans...)

		each := make(map[string]string,  4)
		// 由于是map引用，放在上层for时，rowMaps最终返回值是最后一条。
		for i,  col := range values {
			each[columns[i]] = string(col)
		}

		// 切片追加数据，索引位置有意思。不这样写就不是希望的样子。
		rowMaps = append(rowMaps[:i],  each)
		fmt.Println(each)
		i++
	}
	fmt.Println(rowMaps)

	for i,  col := range rowMaps {
		fmt.Println(i,  col)
	}

	err = rows.Err()
	if err != nil {
		fmt.Printf(err.Error())
	}
}

func transaction() {
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("transaction begin err: %s\n", err)
		return
	}
	//defer tx.Rollback()
	stmt, err := tx.Prepare(`INSERT INTO person (username, sex, email) VALUES (?, ?, ?)`)
	if err != nil {
		fmt.Printf("prepare data err: %s\n", err)
		return
	}
	defer stmt.Close()

	for i := 0; i < 10; i++ {
		username := "user" + strconv.Itoa(rand.Intn(1000) + 10)
		sex := "male"
		if i % 2 == 0 {
			sex = "female"
		}
		email := username + "@qq.com"
		_, err := stmt.Exec(username, sex, email)
		if err != nil {
			fmt.Printf("insert data err: %s\n", err)
			return
		}
	}

	err = tx.Commit()
	if err != nil {
		fmt.Printf("commit transaction error: %v\n",  err)
		return
	}
}