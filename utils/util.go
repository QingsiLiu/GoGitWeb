package utils

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	"time"

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
		db = db1
		CreateTableWithUser()
		CreateTableWithArticle()
		CreateTableWithAlbum()
	}
}

// ModifyDB 操作数据库
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

// CreateTableWithUser 创建用户表
func CreateTableWithUser() {
	sql := `CREATE TABLE IF NOT EXISTS users(
			id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
 			username VARCHAR(100),
			password VARCHAR(400),
			status INT(4),
			createtime INT(10)
			);`
	ModifyDB(sql)
}

// CreateTableWithArticle 创建文章表
func CreateTableWithArticle() {
	sql := `CREATE TABLE IF NOT EXISTS article(
		id int(4) primary key auto_increment not null,
		title varchar(40),
		author varchar(20),
		tags varchar(40),
		short varchar(250),
		content longtext,
		createtime int(10)
		);`
	ModifyDB(sql)
}

//创建图片表
func CreateTableWithAlbum() {
	sql := `CREATE TABLE IF NOT EXISTS ALBUM(
		id 			int(4)	PRIMARY KEY auto_increment not null,
		filepath	varchar(255),
		filename	varchar(64),
		STATUS		int(4),
		createtime	int(10)
	);`
	ModifyDB(sql)
}

// QueryRowDB 查询行数
func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}

func QueryDB(sql string) (*sql.Rows, error) {
	return db.Query(sql)
}

// MD5 密码加密
func MD5(str string) string {
	md5str := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return md5str
}

// SwitchTimeStampToData 时间--->数字格式的转换
func SwitchTimeStampToData(timeStamp int64) string {
	t := time.Unix(timeStamp, 0)
	return t.Format("2006-01-02  15:04:05")
}
