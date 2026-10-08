package action

import (
	"net/http"
	"time"
	"tuko-gg/sentinel/db"

	"github.com/gin-gonic/gin"
)

type GetActionsForItemsRequestPayload struct {
	ItemIds           []string
	SubscribedUserIds []string
}

type GetActionsForItemsActionItem struct {
	db.Label
	AddedTime time.Time
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

	actionsByItemId := make(map[string][]GetActionsForItemsActionItem)
	for _, action := range actions {

		actionsByItemId[action.ItemId] = append(
			actionsByItemId[action.ItemId],
			GetActionsForItemsActionItem{
				Label:     action.Label,
				AddedTime: action.CreatedTime,
			},
		)
	}
	c.IndentedJSON(http.StatusOK, actionsByItemId)
}
