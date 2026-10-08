package action

import (
	"net/http"
	"tuko-gg/sentinel/db"

	"github.com/gin-gonic/gin"
)

type GetActionsForItemsRequestPayload struct {
	ItemIds           []string
	SubscribedUserIds []string
}

func GetActionsForItems(c *gin.Context) {
	var payload GetActionsForItemsRequestPayload
	var actions []db.LabelAction

	if err := c.BindJSON(&payload); err != nil {
		c.IndentedJSON(http.StatusOK, gin.H{
			"error": "Invalid data",
		})
		return
	}

	err := db.GetDb().
		Where("item_id IN ? AND user_id IN ?", payload.ItemIds, payload.SubscribedUserIds).
		Preload("Label").
		Find(&actions).Error

	if err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to retrieve actions",
        })
        return
    }

    actionsByItemId := make(map[string][]db.LabelAction)
    for _, action := range actions {
        actionsByItemId[action.ItemId] = append(actionsByItemId[action.ItemId], action)
    }
	c.IndentedJSON(http.StatusOK, actionsByItemId)
}
