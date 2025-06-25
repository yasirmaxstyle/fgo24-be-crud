package controllers

import (
	"context"
	"dashboard-backend/models"
	"dashboard-backend/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func GetContactByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid contact ID",
		})
		return
	}

	conn := utils.DBConnect()
	defer func() {
		conn.Conn().Close(context.Background())
	}()

	var contact models.Contact
	query := "SELECT id, name, email, phone FROM contacts WHERE id = $1"

	err = conn.QueryRow(context.Background(), query, id).Scan(
		&contact.ID, &contact.Name, &contact.Email, &contact.Phone,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Contact not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch contact"})
		return
	}

	c.JSON(http.StatusOK, contact)
}
