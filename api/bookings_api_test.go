package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"
	"time"
	"bytes"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"studio_api_project/main/models"
	"studio_api_project/main/repositories"
)

func tearDown() {
	// Perform teardown tasks
	repositories.ResetBookings()
	repositories.ResetClasses()
}

func TestStartBookingsAPI(t *testing.T) {
	// The intention of this test is just to validate the availability of the bookings URLs

	// Create a new Gin router
	router := gin.Default()

	// Start the bookings API by calling StartBookingsAPI
	StartBookingsAPI(router)

	repositories.PopulateClassesWithExamples()
	// Define the test data
	testBooking := models.Booking{
		ID:	  0,
		Name: "John Doe",
		Date: models.DailyDate(time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)),
	}
	testBookingJSON, _ := json.Marshal(testBooking)

	// Perform POST request on /bookings
	reqCreateBooking, _ := http.NewRequest(http.MethodPost, "/bookings", bytes.NewBuffer(testBookingJSON))
	resCreateBooking := httptest.NewRecorder()
	router.ServeHTTP(resCreateBooking, reqCreateBooking)
	assert.Equal(t, http.StatusCreated, resCreateBooking.Code)

	// Perform GET request on /bookings
	reqGetBookings, _ := http.NewRequest(http.MethodGet, "/bookings", nil)
	resGetBookings := httptest.NewRecorder()
	router.ServeHTTP(resGetBookings, reqGetBookings)
	assert.Equal(t, http.StatusOK, resGetBookings.Code)

	// Perform GET request on /bookings/0
	reqGetBooking, _ := http.NewRequest(http.MethodGet, "/bookings/0", nil)
	resGetBooking := httptest.NewRecorder()
	router.ServeHTTP(resGetBooking, reqGetBooking)
	assert.Equal(t, http.StatusOK, resGetBooking.Code)

	// Perform PUT request on /bookings/0
	reqUpdateBooking, _ := http.NewRequest(http.MethodPut, "/bookings/0", bytes.NewBuffer(testBookingJSON))
	resUpdateBooking := httptest.NewRecorder()
	router.ServeHTTP(resUpdateBooking, reqUpdateBooking)
	assert.Equal(t, http.StatusOK, resUpdateBooking.Code)

	// Perform DELETE request on /bookings/0
	reqDeleteBooking, _ := http.NewRequest(http.MethodDelete, "/bookings/0", nil)
	resDeleteBooking := httptest.NewRecorder()
	router.ServeHTTP(resDeleteBooking, reqDeleteBooking)
	assert.Equal(t, http.StatusOK, resDeleteBooking.Code)

	tearDown()
}

func TestGetBookings(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	repositories.PopulateClassesWithExamples()
	repositories.PopulateBookingsWithExamples()

	// Define a route handler for GetBookings
	router.GET("/bookings", GetBookings)

	// Perform the request
	req, _ := http.NewRequest("GET", "/bookings", strings.NewReader(``))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, rec.Code)

	tearDown()
}

func TestGetBookingNotFound(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()
	router.GET("/bookings/:id", GetBooking)

	// Perform the request
	req, _ := http.NewRequest("GET", "/bookings/123", strings.NewReader(``))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Check the response status
	assert.Equal(t, http.StatusNotFound, rec.Code)

	tearDown()
}

func TestCreateBooking(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	repositories.PopulateClassesWithExamples()

	// Define a route handler for CreateBooking
	router.POST("/bookings", CreateBooking)

	// Create a request with valid JSON payload
	jsonData := `{"name": "John Doe", "date": "2023-02-01"}`

	// Perform the request
	req, _ := http.NewRequest("POST", "/bookings", strings.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Check the response status code
	assert.Equal(t, http.StatusCreated, rec.Code)

	// Verify the response body contains the created booking
	expectedResponse := `{"id":0,"name":"John Doe","date":"2023-02-01"}`
	assert.Equal(t, expectedResponse, rec.Body.String())

	tearDown()
}

func TestCreateBookingRequiredFields(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	repositories.PopulateClassesWithExamples()

	// Define a route handler for CreateBooking
	router.POST("/bookings", CreateBooking)

	// Create missing a required field (date)
	jsonData := `{"name": "John Doe"}`

	// Perform the request
	req, _ := http.NewRequest("POST", "/bookings", strings.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Check the response status code
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	// Create missing a required field (name)
	jsonData = `{"date":"2023-02-01"}`

	// Perform the request
	req, _ = http.NewRequest("POST", "/bookings", strings.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	rec = httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Check the response status code
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	tearDown()
}

func TestCreateBookingInADateWithoutClasses(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	// Define a route handler for CreateBooking
	router.POST("/bookings", CreateBooking)

	// Create valid Booking
	jsonData := `{"name": "John Doe", "date": "2023-02-01"}`

	// Perform the request
	req, _ := http.NewRequest("POST", "/bookings", strings.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Check the response status code
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	
	tearDown()
}

func TestDeleteBookingNotFound(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()
	router.DELETE("/bookings/:id", DeleteBooking)

	// Perform the request
	req, _ := http.NewRequest("DELETE", "/bookings/123", strings.NewReader(``))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Check the response status
	assert.Equal(t, http.StatusNotFound, rec.Code)

	tearDown()
}

func TestUpdateBookingNotFound(t *testing.T) {
	// Populate with examples
	repositories.PopulateClassesWithExamples()
	repositories.PopulateBookingsWithExamples()
	
	// Create a new Gin router
	router := gin.Default()
	router.PUT("/bookings/:id", UpdateBooking)

	// Perform the request
	req, _ := http.NewRequest("PUT", "/bookings/123", strings.NewReader(`{"id": "123", "name": "John Doe", "date": "2023-02-01"}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Check the response status
	assert.Equal(t, http.StatusNotFound, rec.Code)

	tearDown()
}

func TestUpdateBookingChangingID(t *testing.T) {
	// Populate with examples
	repositories.PopulateClassesWithExamples()
	repositories.PopulateBookingsWithExamples()
	
	// Create a new Gin router
	router := gin.Default()
	router.PUT("/bookings/:id", UpdateBooking)

	// Perform the request
	req, _ := http.NewRequest("PUT", "/bookings/0", strings.NewReader(`{"id": 3, "name": "John Doe", "date": "2023-02-01"}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Check the response status
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	tearDown()
}

func TestUpdateBookingBiding(t *testing.T) {
	// Populate with examples
	repositories.PopulateClassesWithExamples()
	repositories.PopulateBookingsWithExamples()
	
	// Create a new Gin router
	router := gin.Default()
	router.PUT("/bookings/:id", UpdateBooking)

	// Perform the request
	req, _ := http.NewRequest("PUT", "/bookings/0", strings.NewReader(`{"id": "0", "name": "John Doe", "date": "2023-02-01"}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Check the response status
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	tearDown()
}

func TestUpdateBookingDateWithoutClass(t *testing.T) {
	// Populate with examples
	repositories.PopulateClassesWithExamples()
	repositories.PopulateBookingsWithExamples()
	
	// Create a new Gin router
	router := gin.Default()
	router.PUT("/bookings/:id", UpdateBooking)

	// Perform the request
	req, _ := http.NewRequest("PUT", "/bookings/0", strings.NewReader(`{"id": 0, "name": "John Doe", "date": "1999-02-01"}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Check the response status
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	tearDown()
}

func TestUpdateBookingRequiredFields(t *testing.T) {
	// Populate with examples
	repositories.PopulateClassesWithExamples()
	repositories.PopulateBookingsWithExamples()

	// Create a new Gin router
	router := gin.Default()

	// Define a route handler for CreateBooking
	router.PUT("/bookings/:id", CreateBooking)

	// Update missing a required field (date)
	jsonData := `{"id": 0, "name": "John Doe"}`

	// Perform the request
	req, _ := http.NewRequest("PUT", "/bookings/0", strings.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Check the response status code
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	// Update missing a required field (name)
	jsonData = `{"id": 0, "date":"2023-02-01"}`

	// Perform the request
	req, _ = http.NewRequest("PUT", "/bookings/0", strings.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	rec = httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Check the response status code
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	tearDown()
}

