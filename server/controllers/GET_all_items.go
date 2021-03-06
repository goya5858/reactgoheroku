package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func GET_all_items(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/items endpoint is hooked!")
	w.Header().Set("Content-Type", "application/json")
	var items []*ItemParams = GET_all_items_from_SQL()
	json.NewEncoder(w).Encode(items)
}

func GET_all_items_from_SQL() []*ItemParams {
	// DataBase接続
	fmt.Println("Connect MySQL")
	driverName := os.Getenv("DATABASE_URL")
	if driverName == "" {
		driverName = os.Getenv("FOR_LOCAL_DBURL")
	}
	db, err := sql.Open("mysql", driverName)
	//         sql.Open("mysql", "userName: pass@tcp(hostName:3306)/DBname")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success Connect")

	// Query
	fmt.Println("Query Start")
	rows, err := db.Query("SELECT * FROM test_table")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	fmt.Println("Query Done")

	// 返値用のデータ作成
	var items []*ItemParams
	for rows.Next() {
		var one_item ItemParams
		rows.Scan(&one_item.Id, &one_item.ItemName, &one_item.Price, &one_item.Stock)
		items = append(items, &one_item)
		fmt.Println("ID:", one_item.Id, ", ItemName:", one_item.ItemName, ", Price:", one_item.Price, ", Stock:", one_item.Stock)
	}

	return items
}

type ItemParams struct {
	Id       string `json:"id"`
	ItemName string `json:"item_name,omitempty"`
	Price    int    `json:"price,omitempty"`
	Stock    int    `json:"stock,omitempty"`
}
