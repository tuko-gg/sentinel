package router

import (
	"github.com/gin-gonic/gin"
	"tuko-gg/sentinel/router/label"
	"tuko-gg/sentinel/router/action"
)

func InitRouter(router *gin.Engine) {
    // labels
	router.POST("/labels", label.CreateLabelEndpoint)
	router.POST("/labels/edit", label.EditLabelEndpoint)
	router.GET("/labels/from-user", label.GetLabelsFromUserEndpoint)
	router.DELETE("/labels", label.DeleteLabelEndpoint)
	// actions
	router.POST("/actions/apply-label", action.ApplyActionEndpoint)
	router.POST("/actions/delete", action.DeleteActionEndpoint)
	router.POST("/actions/for-items", action.GetActionsForItems)
}
