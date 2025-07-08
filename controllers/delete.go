package controllers

import (
	"context"
	"dashboard-backend/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Delete contacts godoc
// @Summary Delete a contact
// @Description Delete a contact
// @Tags contact
// @Accept json
// @Produce json
// @Param id path integer true "Delete contact"
// @Success 201 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /contacts/{id} [delete]
func DeleteContact(c *gin.Context) {
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

	query := "DELETE FROM contacts WHERE id = $1"
	result, err := conn.Exec(context.Background(), query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete contact"})
		return
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contact not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Contact deleted successfully"})

	utils.RedisClient().Del(context.Background(), "/contacts")
	utils.RedisClient().Del(context.Background(), fmt.Sprintf("/%d", id))
}
