package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/usysrc/support/internal/models"
	"github.com/usysrc/support/internal/services"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	// Create a new Gin router
	router := gin.Default()

	// Register endpoint to create a new ticket
	router.POST("/tickets", func(c *gin.Context) {
		var ticket models.Ticket
		if err := c.BindJSON(&ticket); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := services.CreateTicket(&ticket, db); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"ticket": ticket})
	})

	return router
}
