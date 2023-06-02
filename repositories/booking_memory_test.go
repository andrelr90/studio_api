package repositories

import (
	"testing"
	"time"
	"reflect"

	"studio_api_project/main/models"
)


func TestGetBookings(t *testing.T) {
	// Setup test data
	booking0 := models.Booking{ID: 0, Name: "John Doe", Date: models.DailyDate(time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC))}
	booking1 := models.Booking{ID: 1, Name: "Jane Smith", Date: models.DailyDate(time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC))}
	bookings = map[int]models.Booking{0: booking0, 1: booking1}

	// Call the GetBookings function
	result := GetBookings()

	// Check if the returned bookings match the expected bookings
	expectedBookings := []models.Booking{0: booking0, 1: booking1}
	if result[0].ID != expectedBookings[0].ID || result[1].ID != expectedBookings[1].ID {
		t.Errorf("GetBookings returned unexpected bookings.")
	}
}


func TestGetBooking(t *testing.T) {
	// Prepare test data
	booking0 := models.Booking{ID: 0, Name: "John Doe", Date: models.DailyDate(time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC))}
	booking1 := models.Booking{ID: 1, Name: "Jane Smith", Date: models.DailyDate(time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC))}
	bookings = map[int]models.Booking{0: booking0, 1: booking1}

	// Test case 1: Booking exists
	id := "0"
	expectedBooking := &booking0
	result := GetBooking(id)
	if result.ID != expectedBooking.ID {
		t.Errorf("GetBooking(%s) returned unexpected booking. Got %+v, expected %+v", id, result, expectedBooking)
	}

	// Test case 2: Booking doesn't exist
	id = "3"
	expectedBooking = nil
	result = GetBooking(id)
	if result != expectedBooking {
		t.Errorf("GetBooking(%s) returned unexpected booking. Got %+v, expected %+v", id, result, expectedBooking)
	}
}


func TestCreateBooking(t *testing.T) {
	// Setup test data
	classDate := time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)
	bookings = map[int]models.Booking{}

	// Create a class on the test date
	class := models.Class{
		ID:        1,
		Name:      "Pilates",
		StartDate: models.DailyDate(classDate),
		EndDate:   models.DailyDate(classDate.AddDate(0, 0, 3)),
		Capacity:  30,
	}
	classes = &ClassesStructure{}
	classes.Insert(class)

	// Create a booking with the test class date
	booking := models.Booking{
		ID: 0,
		Name: "John Doe",
		Date: models.DailyDate(classDate),
	}

	// Call the CreateBooking function
	result := CreateBooking(booking)

	// Check if the booking was created successfully
	if result == nil {
		t.Error("CreateBooking failed to create the booking")
	}
	
	// Check if the booking was assigned a new ID
	if result.ID != 0 {
		t.Errorf("CreateBooking assigned an unexpected ID. Got %d, expected 0", result.ID)
	}

	// Check if the booking is added to the bookings slice
	bookings := GetBookings()
	if len(bookings) != 1 {
		t.Errorf("CreateBooking failed to add the booking to the bookings slice. Got %d bookings, expected 1", len(bookings))
	}
	if !reflect.DeepEqual(*result, bookings[0]) {
		t.Errorf("CreateBooking added the booking with incorrect details. Got %+v, expected %+v", bookings[0], *result)
	}
}

func TestCreateBookingInInvalidDate(t *testing.T) {
	// Setup test data
	classDate := time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)
	bookings = map[int]models.Booking{}

	// Create a class on the test date
	class := models.Class{
		ID:        1,
		Name:      "Pilates",
		StartDate: models.DailyDate(classDate),
		EndDate:   models.DailyDate(classDate.AddDate(0, 0, 3)),
		Capacity:  30,
	}
	classes = &ClassesStructure{}
	classes.Insert(class)

	// Create a booking with the test class date
	booking := models.Booking{
		ID: 0,
		Name: "John Doe",
		Date: models.DailyDate(time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC)),
	}

	// Call the CreateBooking function
	result := CreateBooking(booking)

	// Check if the booking was created successfully
	if result != nil {
		t.Error("CreateBooking created a booking in a date without class")
	}
}

func TestDeleteBooking(t *testing.T) {
	// Prepare test data
	booking0 := models.Booking{ID: 0, Name: "John Doe", Date: models.DailyDate(time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC))}
	booking1 := models.Booking{ID: 1, Name: "Jane Smith", Date: models.DailyDate(time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC))}
	bookings = map[int]models.Booking{0: booking0, 1: booking1}

	// Test case 1: Booking exists
	id := "0"
	result := DeleteBooking(id)
	if result != nil {
		t.Errorf("DeleteBooking should delete a booking that exists and return nil")
	}

	// Test case 2: Booking doesn't exist
	id = "3"
	result = DeleteBooking(id)
	if result != nil && result.Error() != "Booking not found" {
		t.Errorf("DeleteBooking should return error when deleting a booking that doesn't exist")
	}
}

func TestUpdateBooking(t *testing.T) {
	// Setup test data
	classDate := time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)
	booking := models.Booking{
		ID: 0,
		Name: "John Doe",
		Date: models.DailyDate(classDate),
	}
	bookings = map[int]models.Booking{0: booking}

	// Create a class on the test date
	class := models.Class{
		ID:        1,
		Name:      "Pilates",
		StartDate: models.DailyDate(classDate),
		EndDate:   models.DailyDate(classDate.AddDate(0, 0, 3)),
		Capacity:  30,
	}
	classes = &ClassesStructure{}
	classes.Insert(class)

	// Test case 1: Booking exists
	_, err := UpdateBookingInStorage(&booking)
	if err != nil {
		t.Errorf("Unable to update a valid booking")
	}

	// Test case 2: Booking doesn't exist
	bookingNotRegistered := models.Booking{
		ID: 2,
		Name: "Peter, the hedgehog",
		Date: models.DailyDate(time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC)),
	}
	_, err = UpdateBookingInStorage(&bookingNotRegistered)
	if err == nil {
		t.Errorf("Updated a booking that doesn't exist")
	}
}

func TestUpdateBookingChangingToInvalidDate(t *testing.T) {
	// Setup test data
	classDate := time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)
	bookings = map[int]models.Booking{0: models.Booking{
		ID: 0,
		Name: "John Doe",
		Date: models.DailyDate(time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC)),
	}}

	// Create a class on the test date
	class := models.Class{
		ID:        1,
		Name:      "Pilates",
		StartDate: models.DailyDate(classDate),
		EndDate:   models.DailyDate(classDate.AddDate(0, 0, 3)),
		Capacity:  30,
	}
	classes = &ClassesStructure{}
	classes.Insert(class)

	// Create a booking with the test class date
	booking := models.Booking{
		ID: 0,
		Name: "John Doe",
		Date: models.DailyDate(time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC)),
	}

	_, error := UpdateBookingInStorage(&booking)

	// Check if the booking was updated successfully
	if error != nil && error.Error() != "There are no classes in this date" {
		t.Error("UpdateBooking updated a booking in a date without class")
	}
}
