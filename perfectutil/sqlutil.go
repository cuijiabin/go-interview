package main

import (
	"database/sql"
	"fmt"
)

func GetMysqlConnection() *sql.DB {
	db, _ := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/beego?charset=utf8")
	return db
}
func GetMysqlDB(host string, port int, user string, passwd string, db string) string {
	dataSourceName := fmt.Sprintf("%s:%s@@tcp(%s:%d)/%s?charset=utf8", user, passwd, host, port, db)
	return dataSourceName
}
func QueryAll(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		fmt.Println("fetech data failed:", err.Error())
		return
	}

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))

	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}
		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
		fmt.Println("-----------------------------------")
	}
	if err = rows.Err(); err != nil {
		panic(err.Error())
	}

	defer rows.Close()
}
