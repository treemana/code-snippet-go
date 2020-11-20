package web

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GIN() {
	app := gin.Default()
	app.GET("/gin", func(c *gin.Context) {
		c.JSON(http.StatusOK, &Unit{
			Name:      "iris",
			Timestamp: time.Now().Unix(),
		})
	})
	_ = app.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
