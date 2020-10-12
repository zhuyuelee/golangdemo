package helper

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

var (
	userName  string = "root"
	password  string = "123123"
	ipAddrees string = "127.0.0.1"
	port      int    = 3306
	dbName    string = "test"
	charset   string = "utf8"
)

//ConnectMysql 链接数据
func ConnectMysql() *sqlx.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", userName, password, ipAddrees, port, dbName, charset)
	Db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("mysql connect failed, detail is [%v]", err.Error())
	}
	return Db
}
