package workday

import (
	"database/sql"

	"github.com/kataras/iris"
)

// Creates and returns an Iris HTTP server
func IrisHTTPServer(db *sql.DB) *iris.Application {
	app := iris.Default()

	app.Handle("GET", "/employees", func(ctx iris.Context) {
		emps, err := GetAllEmployees(db)
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
		}
		ctx.JSON(emps)
	})

	return app
}
