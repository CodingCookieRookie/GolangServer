package Server

import (
	"Webservice/Album"
	"Webservice/Database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RetrieveAlbumByName(c *gin.Context) {
	name := c.Param("Name")
	for i := range Database.Albums {
		if Database.Albums[i].Name == name {
			c.IndentedJSON(http.StatusOK, Database.Albums[i])
			return
		}
	}
	//panic("Name does not exist in database")
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func RetrieveAllAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Database.Albums)
}

func PostAlbum(c *gin.Context) {
	var newAlbum Album.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	Database.Albums = append(Database.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func SetupHTTP() {
	router := gin.Default()
	router.GET("/albums/:Name", RetrieveAlbumByName)
	router.GET("/albums", RetrieveAllAlbums)
	router.POST("/albums", PostAlbum)
	router.Run("localhost:8080")
}

// cmds to use
/*
curl http://localhost:8080/albums/Talking%20to%20the%20moon
curl http://localhost:8080/
curl http://localhost:8080/albums     --include     --header "Content-Type: application/json"     --request "POST"     --data '{"Id": 3,"Name": "The Modern Sound of Betty Carter","Singer": "Betty Carter","Price": 49.99}'
*/
