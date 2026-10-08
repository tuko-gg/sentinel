package label

import (
	"net/http"
	"time"
	"tuko-gg/sentinel/db"

	"github.com/gin-gonic/gin"
)

func EditLabelEndpoint(c *gin.Context) {
	var newLabel db.Label

	if err := c.BindJSON(&newLabel); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid data",
		})
		return
	}

	var labelExisting db.Label

	errLabelExist := db.GetDb().
		Where("label_id = ?", newLabel.LabelId).Find(&labelExisting).Error

	if errLabelExist != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to find label to update",
		})
		return
	}

	if labelExisting == (db.Label{}) {
        c.IndentedJSON(http.StatusNotFound, gin.H{
            "error": "Label not found",
        })
        return
    }

	newLabel.ModifiedTime = time.Now()

	err := db.GetDb().
		Model(&db.Label{}).
		Where("label_id = ?", newLabel.LabelId).
		Updates(map[string]interface{}{
			"name":          newLabel.Name,
			"action":        newLabel.Action,
			"description":   newLabel.Description,
			"modified_time": newLabel.ModifiedTime,
		}).Error

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update label",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, newLabel)
}
