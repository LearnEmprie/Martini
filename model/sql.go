package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"
)

type Db struct {
	con *sql.DB
}

func (db *Db) New ()  {
	con, err := sql.Open("mysql", "root:JSKZ-fc4b679d0%@tcp(101.200.143.50:3306)/usop")

	fmt.Println(err)


	err = con.Ping()

	fmt.Println(err)

	if err != nil {
		log.Fatalf("Open database error: %s\n", err)
	}
	//defer con.Close()
	db.con = con
}

func (db *Db) Insert()  {
	conn   , err := db.con.Prepare("INSERT INTO u_interface(name,type,uri,is_mock,mock_id) VALUES(?,?,?,?,?)")
	result , err := conn.Exec("1","2","3",1,2)
	fmt.Println(result)
	fmt.Println(err)
}