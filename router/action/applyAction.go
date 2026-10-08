package action

import (
	"net/http"
	"time"
	"tuko-gg/sentinel/db"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ApplyActionEndpoint(c *gin.Context) {
	var action db.LabelAction
	var label db.Label

	if err := c.BindJSON(&action); err != nil {
		c.IndentedJSON(http.StatusOK, gin.H{
			"error": "Invalid data",
		})
		return
	}

	errLabelExist := db.GetDb().
		Where("label_id = ?", action.LabelId).
		Find(&label).Error

	if errLabelExist != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to find label to apply action",
		})
		return
	}

	if label == (db.Label{}) {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"error": "Label not found",
		})
		return
	}

    action.UserId = label.UserId
	action.CreatedTime = time.Now()

	err := db.GetDb().Create(&action).Error
	if err != nil {
		if err == gorm.ErrDuplicatedKey {
			c.IndentedJSON(http.StatusNotFound, gin.H{
				"error": "Action already active",
			})
			return
		}

		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to insert action",
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, action)
}
