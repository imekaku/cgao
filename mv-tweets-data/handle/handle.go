package handle

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/cgao/mv-tweets-data/conf"
	"github.com/cgao/mv-tweets-data/model"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var db *sql.DB
var tx *sql.Tx
var g_file_num int
var commit_num int

func Handle() {
	InitHandle()
	filepath.Walk("./data/tweets/", Travel)
	if g_file_num != 0 {
		tx.Commit()
	}
}

func InitHandle() {
	g_file_num = 0
	commit_num = 0

	var err error
	db, err = GetDbConn()

	if err != nil {
		fmt.Println("GetDbConn err=", err)
	}

	tx, err = db.Begin()
	if err != nil {
		fmt.Println("db.Begin() err=", err)
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

func Travel(path string, f os.FileInfo, err error) error {
	if f == nil {
		return err
	} else {
		if path == "./data/tweets/" {
			return nil
		}
		if err := ReadLine(path); err != nil {
			fmt.Println("Handle ReadLine ERR =", err)
		}
	}
	g_file_num = g_file_num + 1
	if g_file_num > 10 {
		g_file_num = 0
		commit_num = commit_num + 1
		tx.Commit()

		fmt.Println("Commit times =", commit_num)
		tx, err = db.Begin()
		if err != nil {
			fmt.Println("db.Begin() err=", err)
		}
	}
	return nil
}

func ReadLine(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		line = strings.TrimSpace(line)
		HandleJson(line)
	}
	return nil
}

func HandleJson(line string) {
	if len(strings.TrimSpace(line)) == 0 {
		return
	}
	byte_line := []byte(line)
	var twi model.TWEETS
	if err := json.Unmarshal(byte_line, &twi); err != nil {
		fmt.Println("handlejson error =", err)
	}
	analyze_tweets(&twi)
}
