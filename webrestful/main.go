package main

import (
	"html/template"
  "webrestful/model"

	"github.com/gin-gonic/gin"
)


func main() {
	router := gin.Default()
	router.POST("/albums", model.PostAlbum)
	router.GET("/albums", model.GetAlbum)

	tmpl := template.Must(template.ParseFiles("layout.html"))

	router.GET("/", func(c *gin.Context) {
		tmpl.Execute(c.Writer, model.CallAlbums())
	})

	router.POST("/", func(c *gin.Context) {
		tmpl.Execute(c.Writer, model.CallAlbums())
	})

	router.Run("localhost:8088")
}
