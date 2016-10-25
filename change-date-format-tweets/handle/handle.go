package handle

import (
	"database/sql"
	"fmt"
	"github.com/cgao/change-date-format-tweets/conf"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Handle() {
	InitHandle()
}

func InitHandle() {
	g_file_num = 0
	commit_num = 0

	var err error
	db, err = GetDbConn()

	if err != nil {
		fmt.Println("GetDbConn err=", err)
	}
}

func GetDbConn() (conn *sql.DB, err error) {
	conn, err = sql.Open("mysql", "root:"+conf.Config().Dbpasswd+"@tcp(127.0.0.1:3306)/"+conf.Config().Dbname+"?charset=utf8")
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		conn.Close()
	}

	return conn, err
}
