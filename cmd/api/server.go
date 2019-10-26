package main

import (
	"fmt"
	"os"
	"github.com/iris-garcia/workday/api"
)

func main() {
	cfg, err := api.LoadDBConfig("./db_config.toml")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	db, err := api.ConnectDB(cfg)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	ginRouter := api.GinRouter(db)
	ginRouter.Run(":8080")

	os.Exit(0)
}
