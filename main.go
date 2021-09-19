package main

import (
	"github.com/gin-gonic/gin"
	"web-service-go/manager"
)

func main(){
	router := gin.Default()
	router.GET("albums/", manager.GetAlbums)
	router.POST("albums/", manager.AddAlbum)
	router.GET("albums/:id/", manager.GetSpecificAlbum)
	router.Run(":8000")
}


