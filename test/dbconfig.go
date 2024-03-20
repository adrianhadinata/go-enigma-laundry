package test

import (
	"submission-project-enigma-laundry/config"

	"github.com/gin-gonic/gin"
)

func TestConnectDB(c *gin.Context) {
	config.TestConnectDB()
	c.JSON(200, gin.H{"message": "Success Connect Database"})
}
