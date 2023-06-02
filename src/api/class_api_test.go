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

	"studio_api/src/models"
	"studio_api/src/repositories"
)

func TestStartClassesAPI(t *testing.T) {
	// The intention of this test is just to validate the availability of the classes URLs

	// Create a new Gin router
	router := gin.Default()

	// Start the classes API by calling StartClassesAPI
	StartClassesAPI(router)

	// Define the test data
	testClass := *models.NewClass(
		0,
		"Pilates2",
		models.DailyDate(time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)),
		models.DailyDate(time.Date(2023, time.January, 31, 0, 0, 0, 0, time.UTC)),
		20,
	)
	testClassJSON, _ := json.Marshal(testClass)

	// Perform POST request on /classes
	reqCreateClass, _ := http.NewRequest(http.MethodPost, "/classes", bytes.NewBuffer(testClassJSON))
	resCreateClass := httptest.NewRecorder()
	router.ServeHTTP(resCreateClass, reqCreateClass)
	assert.Equal(t, http.StatusCreated, resCreateClass.Code)

	// Perform GET request on /classes
	reqGetClasses, _ := http.NewRequest(http.MethodGet, "/classes", nil)
	resGetClasses := httptest.NewRecorder()
	router.ServeHTTP(resGetClasses, reqGetClasses)
	assert.Equal(t, http.StatusOK, resGetClasses.Code)

	// Perform GET request on /classes/0
	reqGetClass, _ := http.NewRequest(http.MethodGet, "/classes/0", nil)
	resGetClass := httptest.NewRecorder()
	router.ServeHTTP(resGetClass, reqGetClass)
	assert.Equal(t, http.StatusOK, resGetClass.Code)

	// Perform PUT request on /classes/0
	reqUpdateClass, _ := http.NewRequest(http.MethodPut, "/classes/0", bytes.NewBuffer(testClassJSON))
	resUpdateClass := httptest.NewRecorder()
	router.ServeHTTP(resUpdateClass, reqUpdateClass)
	assert.Equal(t, http.StatusOK, resUpdateClass.Code)

	// Perform DELETE request on /classes/0
	reqDeleteClass, _ := http.NewRequest(http.MethodDelete, "/classes/0", nil)
	resDeleteClass := httptest.NewRecorder()
	router.ServeHTTP(resDeleteClass, reqDeleteClass)
	assert.Equal(t, http.StatusOK, resDeleteClass.Code)

	tearDown()
}

func TestGetClasses(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	repositories.PopulateClassesWithExamples()
	repositories.PopulateBookingsWithExamples()

	// Define a route handler for GetClasses
	router.GET("/classes", GetClasses)

	// Perform the request
	req, _ := http.NewRequest("GET", "/classes", strings.NewReader(``))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, rec.Code)

	tearDown()
}

func TestGetClassNotFound(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	// Create a new HTTP request with a non-existent class ID
	router.GET("/classes/:id", GetClass)

	// Perform the request
	req, _ := http.NewRequest("GET", "/classes/123", strings.NewReader(``))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Check the response status
	assert.Equal(t, http.StatusNotFound, rec.Code)

	tearDown()
}

func TestCreateClass(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	// Define a route handler for CreateClass
	router.POST("/classes", CreateClass)

	// Create a request with valid JSON payload
	jsonData := `{"name": "Pilates2", "start_date": "2023-02-01", "end_date": "2023-02-28", "capacity": 30}`

	// Perform the request
	req, _ := http.NewRequest("POST", "/classes", strings.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Check the response status code
	assert.Equal(t, http.StatusCreated, rec.Code)

	// Verify the response body
	expectedResponse := `{"id":0,"name":"Pilates2","start_date":"2023-02-01","end_date":"2023-02-28","capacity":30}`
	assert.Equal(t, expectedResponse, rec.Body.String())

	tearDown()
}

func TestCreateClassRequiredFields(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	repositories.PopulateClassesWithExamples()

	// Define a route handler for CreateClass
	router.POST("/classes", CreateClass)

	// Create a missing required field (name)
	jsonData := `{"start_date": "2023-02-01", "end_date": "2023-02-10", "capacity": 10}`

	// Perform the request
	req, _ := http.NewRequest("POST", "/classes", strings.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Check the response status code
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	// Create a missing required field (start_date)
	jsonData = `{"name": "Class A", "end_date": "2023-02-10", "capacity": 10}`

	// Perform the request
	req, _ = http.NewRequest("POST", "/classes", strings.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	rec = httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Check the response status code
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	// Create a missing required field (end_date)
	jsonData = `{"name": "Class A", "start_date": "2023-02-01", "capacity": 10}`

	// Perform the request
	req, _ = http.NewRequest("POST", "/classes", strings.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	rec = httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Check the response status code
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	// Create a missing required field (capacity)
	jsonData = `{"name": "Class A", "start_date": "2023-02-01", "end_date": "2023-02-10"}`

	// Perform the request
	req, _ = http.NewRequest("POST", "/classes", strings.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	rec = httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Check the response status code
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	tearDown()
}

func TestCreateClassDates(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	repositories.PopulateClassesWithExamples()

	// Define a route handler for CreateClass
	router.POST("/classes", CreateClass)

	// Create a class with start_date greater than end_date
	jsonData := `{"name": "Class A", "start_date": "2023-02-10", "end_date": "2023-02-01", "capacity": 10}`

	// Perform the request
	req, _ := http.NewRequest("POST", "/classes", strings.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Check the response status code
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestCreateClassWithOverlap(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	repositories.PopulateClassesWithExamples()

	// Define a route handler for CreateClass
	router.POST("/classes", CreateClass)

	// Create a request with valid JSON payload
	jsonData := `{"name": "Pilates2", "start_date": "2023-02-01", "end_date": "2023-02-02", "capacity": 30}`

	// Perform the request
	req, _ := http.NewRequest("POST", "/classes", strings.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Check the response status code
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	// Verify the response body
	expectedResponse := `{"error":"Intersection found with Pilates"}`
	assert.Equal(t, expectedResponse, rec.Body.String())

	tearDown()
}

func TestDeleteClassNotFound(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	// Create a new HTTP request with a non-existent class ID
	router.DELETE("/classes/:id", DeleteClass)

	// Perform the request
	req, _ := http.NewRequest("DELETE", "/classes/123", strings.NewReader(``))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Check the response status
	assert.Equal(t, http.StatusNotFound, rec.Code)

	tearDown()
}

func TestUpdateClassNotFound(t *testing.T) {
	// Populate with examples
	repositories.PopulateClassesWithExamples()
	repositories.PopulateBookingsWithExamples()
	
	// Create a new Gin router
	router := gin.Default()

	// Create a new HTTP request with a non-existent class ID
	router.PUT("/classes/:id", UpdateClass)

	// Perform the request
	req, _ := http.NewRequest("PUT", "/classes/123", strings.NewReader(`{"id": "123", "name": "John Doe", "date": "2023-02-01"}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Check the response status
	assert.Equal(t, http.StatusNotFound, rec.Code)

	tearDown()
}

func TestUpdateClassChangingID(t *testing.T) {
	// Populate with examples
	repositories.PopulateClassesWithExamples()
	repositories.PopulateBookingsWithExamples()
	
	// Create a new Gin router
	router := gin.Default()

	// Create a new HTTP request with a non-existent class ID
	router.PUT("/classes/:id", UpdateClass)

	// Perform the request
	req, _ := http.NewRequest("PUT", "/classes/0", strings.NewReader(`{"id": 3, "name": "John Doe", "date": "2023-02-01"}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Check the response status
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	tearDown()
}

func TestUpdateClassBinding(t *testing.T) {
	// Populate with examples
	repositories.PopulateClassesWithExamples()
	repositories.PopulateBookingsWithExamples()
	
	// Create a new Gin router
	router := gin.Default()

	// Create a new HTTP request with a non-existent class ID
	router.PUT("/classes/:id", UpdateClass)

	// Perform the request
	req, _ := http.NewRequest("PUT", "/classes/0", strings.NewReader(`{"id": "3", "name": "John Doe", "date": "2023-02-01"}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Check the response status
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	tearDown()
}

func TestUpdateClassWithOverlap(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	repositories.PopulateClassesWithExamples()

	// Define a route handler for UpdateClass
	router.PUT("/classes/:id", UpdateClass)

	// Create a request with valid JSON payload
	jsonData := `{"name": "Pilates2", "start_date": "2023-01-01", "end_date": "2023-03-01", "capacity": 30}`

	// Perform the request
	req, _ := http.NewRequest("PUT", "/classes/0", strings.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Check the response status code
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	// Verify the response body
	expectedResponse := `{"error":"Intersection found with Yoga"}`
	assert.Equal(t, expectedResponse, rec.Body.String())

	tearDown()
}

