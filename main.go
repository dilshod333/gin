package main

import (
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []*album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}
var idCounter = len(albums)

func main() {
	r := gin.Default()
	r.GET("/get", getAlbum)
	r.POST("/create", func(c *gin.Context) {

		idCounter++
		num := strconv.Itoa(idCounter)

		var newAlbum = &album{ID: num, Title: "Can't hurt me ", Artist: "David Gogins", Price: 99}

		c.JSON(http.StatusCreated, newAlbum)
		albums = append(albums, newAlbum)

	})

	r.PUT("/update/:id", func(c *gin.Context) {
		res := c.Param("id")

		var updatedAlbum album
		updatedAlbum.Artist = "salomat"
		updatedAlbum.Title = "kimdir"
		updatedAlbum.Price = 200
		updatedAlbum.ID = res

		for i, n := range albums {
			if n.ID == res {
				albums[i] = &updatedAlbum
				c.JSON(http.StatusOK, gin.H{"message": "Album updated successfully bruhh!!!!"})
				return
			}
		}
	})
	r.DELETE("/delete/:id", DeleteBook)
	r.Run(":8080")
}

func getAlbum(c *gin.Context) {
	
	c.JSON(http.StatusOK, albums)
}

func DeleteBook(c *gin.Context) {

	userId := c.Param("id")

	index := -1
	for i, alb := range albums {
		if alb.ID == userId {
			index = i
			break
		}
	}

	if index != -1 {
		albums = append(albums[:index], albums[index+1:]...)
		c.JSON(http.StatusOK, gin.H{"album": "successfully deleted broo!!!!"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"album": "not found bro!!!"})
	}

}
