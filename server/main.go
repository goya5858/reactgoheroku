package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/goya5858/reactgoheroku/controllers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	StartBackendServer()
}

func StartBackendServer() {
	// buildフォルダを公開
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("../client/build"))))
	// rootディレクトリにアクセスしたときに相対パスで指定した静的サイトをホストする
	// 相対ディレクトリを指定するために client/package.jsonに　"homepage": "./"　を追記する

	// 上記の"homepage": "./"を記述してないと下の記述でしか動かない。他のパスだとバグる
	// http.Handle("/", http.FileServer(http.Dir("./build")))

	fmt.Println("Server Started Port 8080")

	http.HandleFunc("/api/ping", pingResponse)
	http.HandleFunc("/api", samplePage)
	http.HandleFunc("/api/items", controllers.GET_all_items)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Port:" + port)
	http.ListenAndServe(":"+port, nil) // Herokuで実装する場合は環境変数からPORTを取得する
	//http.ListenAndServe(":8080", nil)
}

func pingResponse(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong2")
}

func samplePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("おはよう世界")
	fmt.Fprintf(w, "おはよう世界")
	//migrateDB()
}

// 先に定義しておく必要がある
//var db *sql.DB
//var driver database.Driver
//var m *migrate.Migrate
//var err error

func migrateDB() {
	// DataBase接続
	fmt.Println("Connect MySQL")
	driverName := os.Getenv("DATABASE_URL")
	if driverName == "" {
		fmt.Println("DATABASE_URL is empty")
		driverName = os.Getenv("FOR_LOCAL_DBURL")
		fmt.Println(driverName)
	}
	db, err := sql.Open("mysql", driverName)
	//         sql.Open("mysql", "userName: pass@tcp(hostName:3306)/DBname")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success Connect")

	fmt.Println("Get MySQL Diver")
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Success Get Driver")

	m, err := migrate.NewWithDatabaseInstance(
		"file:///root/migrations/example1",
		"mysql", driver,
	)

	if err := m.Up(); err != nil {
		panic(err.Error())
	}
}
