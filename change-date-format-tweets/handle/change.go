package handle

import (
	"database/sql"
	"fmt"
)

var tx *sql.Tx

func Update_begin() {

}

func Update() {
	tx, err = db.Begin()
	if err != nil {
		fmt.Println("db.Begin() err=", err)
	}

	var tweets_index int64
	var created_at string

	stmt, err := tx.Prepare("select tweets_index, created_at from tweets limit ?, ?")
	if err != nil {
		fmt.Println("tx.Prepare() err=", err)
	}

	rows, err := stmt.Query(tweets_index, tweets_index+10)
	if err != nil {
		fmt.Println("stmt.Query() err=", err)
	}

	for rows.Next() {
		err := rows.Scan(&tweets_index, &created_at)
		if err != nil {

		}
		arr = append(arr, uid)
	}
}
