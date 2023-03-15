package controllers

import (
	"net/http"
	models "personal-backend/api/models/notion"

	"github.com/gin-gonic/gin"
)

func GetAllJournal(c *gin.Context) {
	data, err := models.NewJournalModelNotion().GetAllJournal()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusNotFound,
			"data":    nil,
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data":   data,
		})
	}

}
