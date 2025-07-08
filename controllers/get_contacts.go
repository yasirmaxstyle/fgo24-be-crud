package controllers

import (
	"context"
	"dashboard-backend/models"
	"dashboard-backend/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

// Get contacts godoc
// @Summary Get all contacts
// @Description Get all contacts
// @Tags contact
// @Accept json
// @Produce json
// @Success 201 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /contacts/ [get]
func GetContacts(c *gin.Context) {
	search := c.DefaultQuery("search", "")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit
	var total int

	conn := utils.DBConnect()
	defer func() {
		conn.Conn().Close(context.Background())
	}()

	result := utils.RedisClient().Exists(context.Background(), c.Request.RequestURI)
	if result.Val() == 0 {
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

		encoded, err := json.Marshal(contacts)
		if err != nil {
			return
		}

		utils.RedisClient().Set(context.Background(), c.Request.RequestURI, (encoded), 0)

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
	} else {
		data := utils.RedisClient().Get(context.Background(), c.Request.RequestURI)
		str := data.Val()
		contacts := []models.Contact{}
		json.Unmarshal([]byte(str), &contacts)
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
}

// Get contact by id godoc
// @Summary Get contact by id
// @Description Get contact by id
// @Tags contact
// @Accept json
// @Produce json
// @Param id path integer true "Get contact by id"
// @Success 201 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /contacts/{id} [get]
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
