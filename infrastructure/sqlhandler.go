package infrastructure

import (
	"fmt"
	"time"

	"github.com/H-J-Ainashi/CATech-RESTwithGolang/interfaces/database"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type SqlHandler struct {
	DB *gorm.DB
}

// データベースハンドラの取得
func NewMySqlDb() database.SqlHandler {
	const (
		USERNAME   = "user"
		PASSWORD   = "password"
		SERVERNAME = "mysql_container"
		PORTNUM    = "3306"
		DBNAME     = "mysql_database"
		ARGUMENT   = "charset=utf8mb4&parseTime=true&loc=Local"
	)
	connect := USERNAME + ":" + PASSWORD + "@tcp(" + SERVERNAME + ":" + PORTNUM + ")/" + DBNAME + "?" + ARGUMENT

	db, err := open(connect, 30)
	if err != nil {
		panic(err)
	}

	// 接続の確認
	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}

	// ログの詳細を出力
	db.LogMode(true)
	// データベースのエンジンを設定
	db.Set("gorm:table_options", "ENGINE=InnoDB")

	sqlHandler := new(SqlHandler)
	sqlHandler.DB = db

	return sqlHandler
}

// MySQLの起動確認し、APIコンテナを起動させる
func open(path string, count uint) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", path)
	if err != nil {
		if count == 0 {
			return nil, fmt.Errorf("Failed to Open SQL.")
		}
		time.Sleep(time.Second)
		count--
		return open(path, count)
	}

	return db, nil
}
