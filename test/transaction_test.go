package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	sqlite "github.com/glebarez/sqlite"
	"github.com/pratomoadhi/golden-trail/config"
	"github.com/pratomoadhi/golden-trail/controller"
	"github.com/pratomoadhi/golden-trail/model"
	"github.com/pratomoadhi/golden-trail/utils"
	"gorm.io/gorm"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// Dummy JWT middleware
	r.Use(func(c *gin.Context) {
		// Set userID = 1
		c.Set("userID", uint(1))
		c.Next()
	})

	r.GET("/transactions", controller.ListTransactions)
	return r
}

func setupTestDB() {
	// Open SQLite in-memory DB for isolated tests
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to in-memory database: " + err.Error())
	}

	// ✅ Assign it to config.DB
	config.DB = db

	// ✅ AutoMigrate
	err = db.AutoMigrate(&model.User{}, &model.Transaction{})
	if err != nil {
		panic("Failed to migrate tables: " + err.Error())
	}

	// ✅ Seed user
	config.DB.Create(&model.User{
		Username: "testuser",
		Password: "testpass",
	})

	// ✅ Seed transactions
	for i := 1; i <= 25; i++ {
		config.DB.Create(&model.Transaction{
			UserID: 1,
			Amount: float64(i),
			Type:   "income",
			Note:   fmt.Sprintf("Test #%d", i),
			Date:   utils.Today(),
		})
	}
}

func TestListTransactionsPagination(t *testing.T) {
	setupTestDB()
	router := setupTestRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/transactions?page=2&limit=10", nil)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", w.Code)
	}

	var resp map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	if err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	data := resp["data"].([]interface{})
	if len(data) != 10 {
		t.Errorf("Expected 10 transactions on page 2, got %d", len(data))
	}

	if int(resp["page"].(float64)) != 2 {
		t.Errorf("Expected page 2, got %v", resp["page"])
	}
}
