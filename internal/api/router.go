package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/usysrc/support/internal/models"
	"github.com/usysrc/support/internal/services"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// Register endpoints
	router.POST("/tickets", createTicketHandler(db))
	router.PATCH("/tickets/:id", updateTicketHandler(db))
	router.GET("/tickets", getTicketsHandler(db))
	router.GET("/healthcheck", healthcheckHandler)

	return router
}

func createTicketHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
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
	}
}

func updateTicketHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		ticketID := c.Param("id")

		existingTicket := models.Ticket{}
		if err := db.First(&existingTicket, ticketID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
			return
		}

		if err := c.BindJSON(&existingTicket); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := services.UpdateTicket(&existingTicket, db); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"ticket": existingTicket})
	}
}

func getTicketsHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tickets, err := services.GetTickets(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"tickets": tickets})
	}
}

func healthcheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
