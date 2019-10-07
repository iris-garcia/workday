package main

import (
	"fmt"
	"github.com/iris-garcia/workday"
	"os"
)

func main() {
	err := workday.RunHTTPServer()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}