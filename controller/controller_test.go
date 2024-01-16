package controller

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"part2/repository"
	"part2/service"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupController(numberService service.NumberService) *gin.Engine {
	router := gin.Default()

	router.GET("/", GetAll)
	router.GET("/:number", GetType)
	router.POST("/", Insert)

	return router
}

func TestController(t *testing.T) {
	// Set up the service and controller
	repo := &repository.NumberCollection{}
	numberService := &service.NumberServiceImpl{NumberCollection: repo}

	t.Run("GetAll", func(t *testing.T) {
		// Create a request to get all values
		req, _ := http.NewRequest("GET", "/", nil)
		rec := httptest.NewRecorder()

		// Perform the request
		setupController(numberService).ServeHTTP(rec, req)

		// Check the response
		assert.Equal(t, http.StatusOK, rec.Code)

		// Check the response body
		expectedBody := `[]`
		assert.Equal(t, expectedBody, rec.Body.String())
	})

	t.Run("GetAll must return values", func(t *testing.T) {

		// Create a request to insert a value
		jsonStr := []byte(`{"number": 15}`)
		req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		// Perform the request
		setupController(numberService).ServeHTTP(rec, req)

		// Check the response
		assert.Equal(t, http.StatusCreated, rec.Code)

		// Create a request to get all values
		req, _ = http.NewRequest("GET", "/", nil)
		rec = httptest.NewRecorder()

		// Perform the request
		setupController(numberService).ServeHTTP(rec, req)

		// Check the response
		assert.Equal(t, http.StatusOK, rec.Code)

		// Check the response body
		expectedBody := `["Type 3"]`
		assert.Equal(t, expectedBody, rec.Body.String())
	})

	t.Run("GetType", func(t *testing.T) {
		// Create a request to get type by ID
		req, _ := http.NewRequest("GET", "/1", nil)
		rec := httptest.NewRecorder()

		// Perform the request
		setupController(numberService).ServeHTTP(rec, req)

		// Check the response
		assert.Equal(t, http.StatusOK, rec.Code)

		// Check the response body
		expectedBody := `"1"`
		assert.Equal(t, expectedBody, rec.Body.String())
	})

	t.Run("Insert", func(t *testing.T) {
		// Create a request to insert a value
		jsonStr := []byte(`{"number": 42}`)
		req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		// Perform the request
		setupController(numberService).ServeHTTP(rec, req)

		// Check the response
		assert.Equal(t, http.StatusCreated, rec.Code)
	})
}
