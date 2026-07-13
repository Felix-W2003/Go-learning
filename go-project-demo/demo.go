package main

import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	//创建数据库对象
	db,err := sql.Open("mysql","root:620WFwf0111!@tcp(127.0.0.1:3306)/school")
	db.Ping() 	//与数据库建立连接
	defer db.Close()	//延迟关闭数据库


	if err != nil{
		fmt.Println("数据库连接失败！")
		log.Fatalln(err)
	}


	var version string //声明Mysql数据库的版本
	err2 := db.QueryRow("SELECT VERSION()").Scan(&version) //单行查询

	if err2 != nil{
		log.Fatal(err2)
	}

	fmt.Println(version)
}