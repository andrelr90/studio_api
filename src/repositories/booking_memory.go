package repositories

import (
	"time"
    "sync"
    "strconv"
	"fmt"

	"studio_api/src/models"
)

var (
	bookings          map[int]models.Booking = make(map[int]models.Booking)
	lastBookingID     int = -1
	idBookingMutex    sync.Mutex
)

func PopulateBookingsWithExamples() {
	// Add sample bookings
	bookings[0] = models.Booking{
		ID:       0,
		Name:     "John Doe",
		Date:     models.DailyDate(time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)),
	}
	bookings[1] =models.Booking{
		ID:       1,
		Name:     "Jane Smith",
		Date:     models.DailyDate(time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)),
	}
	lastBookingID = 1

	// Adds the id of this booking to the booking list of the class of the day
	if class := classes.Find(time.Time(bookings[0].Date)); class != nil {
		class.Bookings[bookings[0].ID] = 1
	}
	if class := classes.Find(time.Time(bookings[1].Date)); class != nil {
		class.Bookings[bookings[1].ID] = 1
	}
}

func GetBookings() []models.Booking {
	var bookingSlice []models.Booking

	// Convert the map to a slice
	for _, value := range bookings {
		bookingSlice = append(bookingSlice, value)
	}

	return bookingSlice
}

func GetBooking(id string) *models.Booking {
	idInt, _ := strconv.Atoi(id)
	if booking, exists := bookings[idInt]; exists {
		return &booking
	}
	
	return nil
}

func CreateBooking(booking models.Booking) *models.Booking {
	// Verify if there are classes in that day
	if classInDate := classes.Find(time.Time(booking.Date)); classInDate == nil {
		return nil
	}

	// Generate a new ID by incrementing the last ID
	idBookingMutex.Lock()
	lastBookingID++
	booking.ID = lastBookingID
	idBookingMutex.Unlock()

	// Add the booking to the slice
	bookings[booking.ID] = booking

	// Adds the id of this booking to the booking list of the class of the day
	if class := classes.Find(time.Time(booking.Date)); class != nil {
		class.Bookings[booking.ID] = 1
	}

	// Returns the booking with its id
	return &booking
}

func DeleteBooking(id string) error {
	idInt, _ := strconv.Atoi(id)
	if _, exists := bookings[idInt]; exists {
		// Removes the id of the old booking of the booking list of the class of the day
		if class := classes.Find(time.Time(bookings[idInt].Date)); class != nil {
			delete(class.Bookings, bookings[idInt].ID)
		}

		delete(bookings, idInt)
		return nil
	}
	
	return fmt.Errorf("Booking not found")
}

func UpdateBookingInStorage(updatedBooking *models.Booking) (*models.Booking, error) {
	// Verify if there are classes in that day
	if classInDate := classes.Find(time.Time(updatedBooking.Date)); classInDate == nil {
		return nil, fmt.Errorf("There are no classes in this date")
	}
	
	// Updates the booking
	if _, exists := bookings[updatedBooking.ID]; exists {
		oldBooking := bookings[updatedBooking.ID]
		// Removes the id of the old booking of the booking list of the class of the day
		if class := classes.Find(time.Time(oldBooking.Date)); class != nil {
			delete(class.Bookings, oldBooking.ID)
		}
		// Insert the id of the new booking to the booking list of the class of the day
		if class := classes.Find(time.Time(updatedBooking.Date)); class != nil {
			class.Bookings[updatedBooking.ID] = 1
		}

		bookings[updatedBooking.ID] = *updatedBooking
		return updatedBooking, nil
	}

	// If no Booking with the matching ID is found, return an error
	return nil, fmt.Errorf("Booking not found")
}

func ResetBookings() {
	// This function is used mostly for tests
	bookings = make(map[int]models.Booking)
	lastBookingID = -1
}
