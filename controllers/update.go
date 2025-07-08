package controllers

import (
	"context"
	"dashboard-backend/models"
	"dashboard-backend/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Update contact godoc
// @Summary Update contact
// @Description Update contact
// @Tags contact
// @Accept json
// @Produce json
// @Param request body models.ContactInput true "Update contact"
// @Param id path integer true "Update contact"
// @Success 201 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /contacts/{id} [patch]
func UpdateContact(c *gin.Context) {
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

	var req models.ContactInput
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `
		UPDATE contacts 
		SET name = $1, email = $2, phone = $3, updated_at = CURRENT_TIMESTAMP 
		WHERE id = $4
	`

	result, err := conn.Exec(context.Background(), query, req.Name, req.Email, req.Phone, id)
	if err != nil {
		// Check for duplicate email constraint
		if err.Error() == "ERROR: duplicate key value violates unique constraint \"contacts_email_key\" (SQLSTATE 23505)" {
			c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update contact"})
		return
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contact not found"})
		return
	}

	// Return updated contact
	contact := models.Contact{
		ID:    int(id),
		Name:  req.Name,
		Email: req.Email,
		Phone: req.Phone,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Contact updated successfully",
		"contact": contact,
	})

	utils.RedisClient().Del(context.Background(), "/contacts")
	utils.RedisClient().Del(context.Background(), fmt.Sprintf("/%d", id))
}
