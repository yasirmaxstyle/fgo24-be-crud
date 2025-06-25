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

func GetContacts(c *gin.Context) {
	search := c.DefaultQuery("search", "")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))
	offset := (page - 1) * limit

	conn := utils.DBConnect()
	defer func() {
		conn.Conn().Close(context.Background())
	}()

	query := `
		SELECT id, name, email, phone 
		FROM contacts
		WHERE name ILIKE $1
		LIMIT $2 OFFSET $3
	`

	rows, err := conn.Query(
		context.Background(),
		query,
		fmt.Sprintf("%%%s%%", search),
		limit,
		offset,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Failed to fetch contacts",
		})
		return
	}

	var contacts []models.Contact
	for rows.Next() {
		var contact models.Contact
		if err := rows.Scan(&contact.ID, &contact.Name, &contact.Email, &contact.Phone); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan contact"})
			return
		}
		contacts = append(contacts, contact)
	}
	var total int
	countQuery := "SELECT COUNT(*) FROM contacts"
	if err := conn.QueryRow(context.Background(), countQuery).Scan(&total); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get total count"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"contacts": contacts,
		"pagination": gin.H{
			"page":        page,
			"limit":       limit,
			"total":       total,
			"total_pages": (total + limit - 1) / limit,
		},
	})
}
