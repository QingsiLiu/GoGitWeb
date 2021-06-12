package utils

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	//驱动包
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func InitMysql() {
	fmt.Println("InitMysql......")
	driverName := beego.AppConfig.String("driverName")

	//数据库连接
	user := beego.AppConfig.String("mysqluser")
	pwd := beego.AppConfig.String("mysqlpwd")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	dbname := beego.AppConfig.String("dbname")

	//数据库链接
	dbconn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"

	db1, err := sql.Open(driverName, dbconn)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("数据库链接成功")
		db = db1
		CreateTableWithUser()
	}
}

//操作数据库
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}

//创建用户表
func CreateTableWithUser() {
	sql := `CREATE TABLE IF NOT EXISTS users(
			id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
 			username VARCHAR(100),
			password VARCHAR(400),
			status INT(4),
			createtime INT(10)
			);`
	ModifyDB(sql)
	fmt.Println("已创建数据库表格：users")
}

//查询
func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}

//密码加密
func MD5(str string) string {
	md5str := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return md5str
}
