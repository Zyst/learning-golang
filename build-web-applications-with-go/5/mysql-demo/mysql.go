package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	db, err := sql.Open("mysql", "root@/sys?charset=utf8")
	defer db.Close()

	checkErr(err)

	// Insert
	statement, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	checkErr(err)

	res, err := statement.Exec("Erick", "Romero", "1992-07-03")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	// Update
	statement, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = statement.Exec("erickXupdate", id)
	checkErr(err)

	affected, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affected)

	// Query
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string

		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)

		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	// Delete
	statement, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = statement.Exec(id)
	checkErr(err)

	affected, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affected)
}
