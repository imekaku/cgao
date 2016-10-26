package handle

import (
	"database/sql"
	"fmt"
	"strings"
)

var tx *sql.Tx
var select_num int64 = 0
var commit_num int64 = 0

func Select_update() {
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("db.Begin() err=", err)
	}

	var tweets_index int64 = 0
	var created_at string

	for {
		select_num = select_num + 1
		if select_num > 100 {
			select_num = 0
			commit_num = commit_num + 1
			tx.Commit()

			fmt.Println("Change recoder num =", commit_num*100)
			tx, err = db.Begin()
			if err != nil {
				fmt.Println("Select_update db.Begin() err=", err)
			}
		}

		stmt, err := tx.Prepare("select tweets_index, created_at from tweets limit ?, ?")
		if err != nil {
			fmt.Println("tx.Prepare() err=", err)
		}

		rows, err := stmt.Query(tweets_index, tweets_index+100)
		if err != nil {
			if err == sql.ErrNoRows {
				break
			} else {
				fmt.Println("stmt.Query() err=", err)
			}
		}

		for rows.Next() {
			err := rows.Scan(&tweets_index, &created_at)
			if err != nil {
				fmt.Println("change.go Select_update rows.Scan err=", err)
			}
			new_created_at := Change_time_format(created_at)
			tx.Exec("update tweets set created_at = ? where tweets_index = ?", new_created_at, tweets_index)
		}
	}
}

func Change_time_format(created_at string) string {
	var month string
	time_slice := strings.Split(created_at, " ")
	switch time_slice[1] {
	case "Jan":
		month = "01"
	case "Feb":
		month = "02"
	case "Mar":
		month = "03"
	case "Apr":
		month = "04"
	case "May":
		month = "05"
	case "Jun":
		month = "06"
	case "Jul":
		month = "07"
	case "Aug":
		month = "08"
	case "Sept":
		month = "09"
	case "Oct":
		month = "10"
	case "Nov":
		month = "11"
	case "Dec":
		month = "12"
	}
	new_created_at := time_slice[5] + "-" + month + "-" + time_slice[2] + "T" + time_slice[3]
	return new_created_at
}
