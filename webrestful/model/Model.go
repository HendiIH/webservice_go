package model

import(
  "github.com/gin-gonic/gin"
  "net/http"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Xeena", Price: 56.99},
	{ID: "2", Title: "Red Doorz", Artist: "Za Roku", Price: 55.99},
	{ID: "3", Title: "Azuz Dura", Artist: "Bambam", Price: 55.99},
	{ID: "4", Title: "Tree Joe", Artist: "Ned and Needle Fish", Price: 56.99},
}

func CallAlbums()[]album{
  return albums
}

func GetAlbum(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func PostAlbum(c *gin.Context) {
	var newAlbum album
	if err := c.ShouldBind(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusOK, albums)
}
