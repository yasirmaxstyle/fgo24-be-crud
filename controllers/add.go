package controllers

import (
	"context"
	"dashboard-backend/models"
	"dashboard-backend/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddContact(c *gin.Context) {
	var req models.ContactInput
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	conn := utils.DBConnect()
	defer func() {
		conn.Conn().Close(context.Background())
	}()

	var contactID int64
	query := `
		INSERT INTO contacts (name, email, phone) 
		VALUES ($1, $2, $3) 
		RETURNING id
	`
	err := conn.QueryRow(context.Background(), query, req.Name, req.Email, req.Phone).Scan(&contactID)
	if err != nil {
		// Check for duplicate email constraint
		if err.Error() == "ERROR: duplicate key value violates unique constraint \"contacts_email_key\" (SQLSTATE 23505)" {
			c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create contact"})
		return
	}

	// Return the created contact
	contact := models.Contact{
		ID:    int(contactID),
		Name:  req.Name,
		Email: req.Email,
		Phone: req.Phone,
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Contact created successfully",
		"contact": contact,
	})

	//reset key in redis
	utils.RedisClient().Del(context.Background(), "/contacts")
	utils.RedisClient().Del(context.Background(), fmt.Sprintf("/%d", contact.ID))
}
