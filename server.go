package workday

import (
	"fmt"

	"github.com/kataras/iris"
)

// Runs a HTTP server using iris web framework
func RunHTTPServer() error {
	cfg, err := LoadDBConfig("./db_config.toml")
	if err != nil {
		return err
	}

	fmt.Println(cfg)

	app := iris.Default()

	app.Handle("GET", "/employees", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "GET to employess endpoint"})
	})

	app.Run(iris.Addr(":8080"))

	return nil
}
