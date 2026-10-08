package label

import (
	"log"
	"net/http"
	"tuko-gg/sentinel/db"
	"github.com/gin-gonic/gin"
	"tuko-gg/sentinel/util"
)

func GetLabelsFromUserEndpoint(c *gin.Context) {
	userId := c.Query("user_id")

	if userId == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Missing user_id query parameter"})
		return
	}

	offset, offsetErr := util.ParseQueryOffset(c)
	limit, limitErr := util.ParseQueryLimit(c, 100)

	if offsetErr != nil || limitErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid offset or limit query parameter"})
		return
	}

	var labels []db.Label

	err := db.GetDb().
		Where("user_id = ?", userId).
		Offset(offset).
		Limit(limit).
		Order("created_time DESC").
		Find(&labels).Error

	if err != nil {
		log.Default().Println("Error retrieving labels from database:", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve labels"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"labels": labels,
	})
}
