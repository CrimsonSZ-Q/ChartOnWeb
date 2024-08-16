package main

import (
	"startlab/ChartOnWeb/controllers"
	"startlab/ChartOnWeb/models"

	"github.com/gin-gonic/gin"
)

func main() {
	models.InitDB()

	r := gin.Default()
	r.LoadHTMLGlob("file/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	r.GET("/chart-data/:id", controllers.GetLatenessGraphData)
	r.GET("/chart-data-bar/:type", controllers.GetScoreData)

	r.Run(":8080") // Menjalankan server di port 8080
}
