package workday

import (
	"fmt"

	"github.com/kataras/iris"
)

// Creates and returns an Iris HTTP server
func IrisHTTPServer() (*iris.Application, error) {
	cfg, err := LoadDBConfig("./db_config.toml")
	if err != nil {
		return nil, err
	}
	fmt.Println(cfg)

	app := iris.Default()

	app.Handle("GET", "/employees", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "GET to employess endpoint"})
	})

	return app, nil
}
