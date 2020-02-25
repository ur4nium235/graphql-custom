package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

/**
 *
 * @author: hoangtq
 * @timeCreate: 25/02/2020 13:55
 * To change this template use File | Settings | Editor | File and Code Template | Includes
 * */

//func init()  {
//	// Register giúp db driver available với "mysql"
//	// nếu hàm này được gọi 2 lần cùng tên db hoặc driver nil sẽ gây ra panic
//	sql.Register("mysql", &mysql.MySQLDriver{})
//}

func NewConnectionSQL(user, pass, host, name string) (*sql.DB, error) {
	dbname := fmt.Sprint(user, ":",
						 pass, "@",
						 host, "/",
						 name, "?charset=utf8")
	
	db, err := sql.Open("mysql", dbname)

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(30)
	db.SetMaxIdleConns(16)
	db.SetConnMaxLifetime(10 * time.Second)
	return db, nil
}
