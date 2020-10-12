package main

import (
	"fmt"
	"mysql/helper"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	db := helper.ConnectMysql()
	//关闭数据库
	defer func() {
		defer db.Close()
	}()
	//添加
	//addRecord(db)
	//修改
	//updateRecord(db, 1)
	//updateRecord(db, 2)

	//删除
	//deleteRecord(db, 1)
	//查询数据
	queryData(db)
	//目标字符串
	fmt.Println("hellw go")
}

//queryData 查询数据
func queryData(Db *sqlx.DB) {
	rows, err := Db.Query("select * from `b_test` order by id desc")
	//关闭结果集（释放连接）
	defer rows.Close()
	if err != nil {
		fmt.Printf("query faied, error:[%v]", err.Error())
		return
	}
	for rows.Next() {
		//定义变量接收查询数据
		var id int
		var username, password, create_at string

		err := rows.Scan(&id, &username, &password, &create_at)
		if err != nil {
			fmt.Printf("get data failed, error:[%v]\n", err.Error())
		}
		fmt.Println(id, username, password, create_at)
	}
}

// addRecord 添加
func addRecord(Db *sqlx.DB) {
	for i := 0; i < 2; i++ {

		result, err := Db.Exec("INSERT INTO `b_test`(`username`, `password`, `create_at`) values (?,?,?)", getGUID(), "test", time.Now().Format("2006-01-02 15:04:05"))
		if err != nil {
			fmt.Printf("data insert faied, error:[%v]", err.Error())
			return
		}
		id, _ := result.LastInsertId()
		fmt.Printf("insert success, last id:[%d]\n", id)
	}
}

//updateRecord 修改
func updateRecord(Db *sqlx.DB, id int) {
	result, err := Db.Exec("update `b_test`set `password`=? where `id`=?;", getGUID(), id)
	if err != nil {
		fmt.Printf("data update faied, error:[%v]", err.Error())
		return
	}
	num, _ := result.RowsAffected()
	fmt.Printf("update success, affected rows:[%d]\n", num)
}

//deleteRecord 删除数据
func deleteRecord(Db *sqlx.DB, id int) {
	//删除数据
	result, err := Db.Exec("delete from `b_test` where id=?;", id)
	if err != nil {
		fmt.Printf("delete faied, error:[%v]", err.Error())
		return
	}
	num, _ := result.RowsAffected()
	fmt.Printf("delete success, affected rows:[%d]\n", num)
}

func getGUID() string {
	now := time.Now()
	var guid = fmt.Sprintf("%d%d%d%d%d%d", now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond())
	return guid

}
