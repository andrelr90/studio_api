package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// The intention of this test is exclusively to check if the APIs are being started.
// It does not test extensively the endpoints. For that, see the unit test of the StartAPI methods.
func TestSetupServerShouldStartAPIs(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	// Run the setup
	SetupServer(router)

	// Validate if /classes is accessible
	reqGetClasses, _ := http.NewRequest(http.MethodGet, "/classes", nil)
	resGetClasses := httptest.NewRecorder()
	router.ServeHTTP(resGetClasses, reqGetClasses)
	assert.Equal(t, http.StatusOK, resGetClasses.Code)

	// Validate if /bookings is accessible
	reqGetBookings, _ := http.NewRequest(http.MethodGet, "/bookings", nil)
	resGetBookings := httptest.NewRecorder()
	router.ServeHTTP(resGetBookings, reqGetBookings)
	assert.Equal(t, http.StatusOK, resGetBookings.Code)
}