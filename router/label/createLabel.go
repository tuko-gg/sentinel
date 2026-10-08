package label

import (
	"net/http"
	"time"
	"tuko-gg/sentinel/db"

	"github.com/gin-gonic/gin"
)

func CreateLabelEndpoint(c *gin.Context) {
	var newLabel db.Label

	if err := c.BindJSON(&newLabel); err != nil {
		c.IndentedJSON(http.StatusOK, gin.H{
			"error": "Invalid data",
		})
		return
	}

	newLabel.CreatedTime = time.Now()
	newLabel.ModifiedTime = newLabel.CreatedTime

	err := db.GetDb().Create(&newLabel).Error
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
		    "error": "Failed to insert label",
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, newLabel)
}
