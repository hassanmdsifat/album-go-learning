package manager

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web-service-go/fixtures"
	"web-service-go/model"
)

var allAlbum = fixtures.AllAlbum

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, allAlbum)
}

func AddAlbum(c *gin.Context){
	var newAlbum model.Album
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, newAlbum)
	}
	allAlbum = append(allAlbum, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func GetSpecificAlbum(c *gin.Context){
	id := c.Param("id")
	for _, album := range allAlbum {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Not found"})
}

