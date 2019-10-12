package main

import (
	"fmt"
	"os"

	"github.com/iris-garcia/workday"
	"github.com/kataras/iris"
)

func main() {
	cfg, err := workday.LoadDBConfig("./db_config.toml")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	db, err := workday.ConnectDB(cfg)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	app := workday.IrisHTTPServer(db)
	app.Run(iris.Addr(":8080"))

	os.Exit(0)
}
