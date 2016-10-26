package handle

import (
	"database/sql"
	"fmt"
	"strings"
)

var tx *sql.Tx

func Select_update() {
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("db.Begin() err=", err)
	}

	var tweets_index int64 = 0
	var created_at string
	var tweets_index_slice []int64
	var new_created_at_slice []string

	for {
		rows, err := tx.Query("select tweets_index, created_at from tweets limit ?, ?", tweets_index, tweets_index+1000)
		if err != nil {
			// 为什么没有等于sql.ErrNoRows
			if err == sql.ErrNoRows {
				tx.Commit()
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
			new_created_at_slice = append(new_created_at_slice, new_created_at)
			tweets_index_slice = append(tweets_index_slice, tweets_index)
		}
		tx.Commit()
		tx, err = db.Begin()
		if err != nil {
			fmt.Println("Select_update db.Begin() err=", err)
		}
		tweets_index_slice_len := len(tweets_index_slice)
		for i := 0; i < tweets_index_slice_len; i++ {
			fmt.Println("update =", i)
			tx.Exec("update tweets set created_at = ? where tweets_index = ?", new_created_at_slice[i], tweets_index_slice[i])
		}

		// 相等于 sql.ErrNoRows
		if tweets_index_slice_len == 0 {
			tx.Commit()
			break
		}
		new_created_at_slice = new_created_at_slice[:0]
		tweets_index_slice = tweets_index_slice[:0]
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
