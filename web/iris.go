package web

import (
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

func IRIS() {

	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())
	app.Get("/iris", func(ctx iris.Context) {
		_ = ctx.JSON(&Unit{
			Name:      "iris",
			Timestamp: time.Now().Unix(),
		})
	})
	_ = app.Listen(":8080")
}
