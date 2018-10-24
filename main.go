package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB
var dbErr error

func main() {
	os.Remove("./foo.db")
	db, dbErr = sql.Open("sqlite3", "./foo.db")
	//db, dbErr = sql.Open("sqlite3", ":memory:")
	checkErr(dbErr)
	defer db.Close()

	//create table
	sqlStmt := `create table foo(id integer not null primary key, name text,language text,gender text);
	delete from foo;
	`
	_, err := db.Exec(sqlStmt)
	checkErr(err)

	dbWriterService()
	simpleSelect()
	//selectWithWhere(2)
}

func dbWriterService() {
	log.Println("db writer service started")

	tx, err := db.Begin()
	checkErr(err)

	for index := 0; index < 1; index++ {
		sqliteInsert(tx, index)
	}
	tx.Commit()

	log.Println("all rows inserted")
}

func sqliteInsert(tx *sql.Tx, index int) {
	/* stmt, err := tx.Prepare("insert into foo(id, name,language,gender) values(?, ?, ?, json('{\"cell\":\"+491765\", \"home\":\"+498973\"}'))")
	checkErr(err)
	defer stmt.Close() */

	/* _, err = stmt.Exec(index, fake.FirstName(), fake.Language(), fake.Gender())
	checkErr(err) */

	/* sqlStmt := "insert into foo(id, name,language,gender) values(3,'foo','en',json('{\"cell\":\"+491765\", \"home\":\"+498973\"}'))" */
	sqlStmt := "insert into foo(id, name,language,gender) values(3,'foo','en',json('{\"cell\":\"+789\", \"home\":\"+101112\"}'))"
	_, err := db.Exec(sqlStmt)
	sqlStmt = "insert into foo(id, name,language,gender) values(4,'bar','en',json('{\"cell\":\"+456\", \"home\":\"+123\"}'))"
	_, err = db.Exec(sqlStmt)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func simpleSelect() {
	rows, err := db.Query("select id, name,language,gender from foo where json_extract(gender,'$.cell')='+456'")
	/* rows, err := db.Query("select id, name,language,gender from foo") */
	checkErr(err)

	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var language string
		var gender string
		err = rows.Scan(&id, &name, &language, &gender)
		checkErr(err)
		fmt.Println(id, name, language, gender)
	}

	err = rows.Err()
	checkErr(err)
}

func selectWithWhere(id int) {
	stmt, err := db.Prepare("select name from foo where id = ?")
	checkErr(err)

	defer stmt.Close()

	var name string
	err = stmt.QueryRow(id).Scan(&name)
	checkErr(err)
	fmt.Println(name)
}
