package label

import (
    "net/http"
    "tuko-gg/sentinel/db"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func DeleteLabelEndpoint(c *gin.Context) {
    labelId := c.Query("label_id")

    if labelId == "" {
        c.IndentedJSON(http.StatusBadRequest, gin.H{
            "error": "Missing label_id query parameter",
        })
        return
    }

    err := db.GetDb().Transaction(func(tx *gorm.DB) error {

        // Check that the label exists
        var label db.Label

        if err := tx.
            Where("label_id = ?", labelId).
            First(&label).Error; err != nil {

            if err == gorm.ErrRecordNotFound {
                return err
            }

            return err
        }

        // Delete all actions belonging to this label
        if err := tx.
            Where("label_id = ?", labelId).
            Delete(&db.LabelAction{}).Error; err != nil {
            return err
        }

        // Delete the label
        if err := tx.
            Where("label_id = ?", labelId).
            Delete(&db.Label{}).Error; err != nil {
            return err
        }

        return nil
    })

    if err != nil {
        if err == gorm.ErrRecordNotFound {
            c.IndentedJSON(http.StatusNotFound, gin.H{
                "error": "Label not found",
            })
            return
        }

        c.IndentedJSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to delete label and associated actions",
        })
        return
    }

    c.IndentedJSON(http.StatusOK, gin.H{
        "message": "Label and associated actions deleted",
    })
}
