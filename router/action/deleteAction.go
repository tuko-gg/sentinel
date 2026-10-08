package action

import (
	"net/http"
	"tuko-gg/sentinel/db"

	"github.com/gin-gonic/gin"
)

type DeleteActionEndpointRequestPayload struct {
	LabelId string
	ItemId  string
}

func DeleteActionEndpoint(c *gin.Context) {
	var payload DeleteActionEndpointRequestPayload

	if err := c.BindJSON(&payload); err != nil {
		c.IndentedJSON(http.StatusOK, gin.H{
			"error": "Invalid data",
		})
		return
	}

	if err := db.GetDb().
		Where("label_id = ? AND item_id = ?", payload.LabelId, payload.ItemId).
		Delete(&db.LabelAction{}).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete action",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Action deleted",
	})
}
