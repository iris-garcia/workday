package main

import (
	"fmt"
	"os"

	"github.com/iris-garcia/workday"
	"github.com/kataras/iris"
)

func main() {
	app, err := workday.IrisHTTPServer()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	app.Run(iris.Addr(":8080"))

	os.Exit(0)
}
