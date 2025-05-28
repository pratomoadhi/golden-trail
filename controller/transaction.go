package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pratomoadhi/golden-trail/config"
	"github.com/pratomoadhi/golden-trail/model"
)

// ListTransactions godoc
// @Summary List all transactions with pagination
// @Tags Transactions
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Success 200 {object} map[string]interface{}
// @Router /transactions/ [get]
// @Security BearerAuth
func ListTransactions(c *gin.Context) {
	// Extract userID from context
	userIDRaw, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userID := userIDRaw.(uint)

	// Parse pagination parameters
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit

	var txs []model.Transaction
	var total int64

	// Count total transactions
	if err := config.DB.Model(&model.Transaction{}).
		Where("user_id = ?", userID).
		Count(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Fetch paginated transactions
	if err := config.DB.Where("user_id = ?", userID).
		Order("date DESC").
		Offset(offset).
		Limit(limit).
		Find(&txs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":        txs,
		"total":       total,
		"page":        page,
		"limit":       limit,
		"total_pages": int((total + int64(limit) - 1) / int64(limit)),
	})
}

// CreateTransaction godoc
// @Summary Create new transaction
// @Tags Transactions
// @Accept json
// @Produce json
// @Param transaction body model.TransactionInput true "Transaction data"
// @Success 201 {object} model.Transaction
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /transactions/ [post]
// @Security BearerAuth
func CreateTransaction(c *gin.Context) {
	var input model.TransactionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userIDRaw, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userID := userIDRaw.(uint) // cast accordingly

	tx := model.Transaction{
		UserID: userID,
		Amount: input.Amount,
		Type:   input.Type,
		Note:   input.Note,
		Date:   parseDate(input.Date),
	}

	if err := config.DB.Create(&tx).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, tx)
}

func parseDate(dateStr string) time.Time {
	t, _ := time.Parse("2006-01-02", dateStr)
	return t
}
