//package main
//
//import (
//	"todoapp/config"
//	"todoapp/db"
//	"todoapp/routes"
//
//	"github.com/gin-gonic/gin"
//)
//
//func main() {
//	r := gin.Default()
//
//	config.ConnectDatabase()
//	db.Migrate()
//
//	r.Static("/css", "./public/css")
//	r.Static("/js", "./public/js")
//	r.LoadHTMLGlob("views/*.html")
//
//	r.GET("/", func(c *gin.Context) {
//		c.HTML(200, "index.html", nil)
//	})
//
//	routes.SetupRoutes(r)
//
//	r.Run(":8080")
//}

package main

import (
	"github.com/gin-gonic/gin"
	"todoapp/config"
	"todoapp/db"
	"todoapp/routes"
)

func main() {
	r := gin.Default()

	config.ConnectDatabase()
	db.Migrate()

	r.Static("/css", "./public/css")
	r.Static("/js", "./public/js")
	r.LoadHTMLFiles("views/index.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	routes.SetupRoutes(r)

	r.Run(":8080")
}
