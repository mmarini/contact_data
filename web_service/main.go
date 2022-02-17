package main

import (
	"net/http"

	"contact_data/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/contacts", postContact)

	router.Run("localhost:8080")
}

func postContact(c *gin.Context) {

	c.Header("Access-Control-Allow-Origin", "*")

	var newContact models.Contact

	if err := c.BindJSON(&newContact); err != nil {
		return
	}

	models.SaveContact(newContact)

	c.IndentedJSON(http.StatusCreated, newContact)
}
