package repositories

import (
	"time"
    "sync"
    "strconv"
	"fmt"

	"studio_api_project/main/models"
)

var (
	bookings          map[int]models.Booking = make(map[int]models.Booking) //[]models.Booking
	lastBookingID     int = -1
	idBookingMutex    sync.Mutex
)

func PopulateBookingsWithExamples() {
	// Add sample bookings
	bookings[0] = models.Booking{
		ID:       0,
		Name:     "John Doe",
		Date:     models.DailyDate(time.Now()),
	}
	bookings[1] =models.Booking{
		ID:       1,
		Name:     "Jane Smith",
		Date:     models.DailyDate(time.Now()),
	}
	lastBookingID = 1
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
	// for _, booking := range bookings {
	// 	if strconv.Itoa(booking.ID) == id {
	// 		return &booking
	// 	}
	// }
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
	// bookings = append(bookings, booking)
	bookings[booking.ID] = booking

	// Returns the booking with its id
	return &booking
}

func DeleteBooking(id string) error {
	idInt, _ := strconv.Atoi(id)
	if _, exists := bookings[idInt]; exists {
		delete(bookings, idInt)
		return nil
	}
	// for index, booking := range bookings {
	// 	if strconv.Itoa(booking.ID) == id {
	// 		bookings = append(bookings[:index], bookings[index+1:]...)
	// 		return nil
	// 		break
	// 	}
	// }
	return fmt.Errorf("Booking not found")
}

// UpdateBookingInStorage updates a class booking in the storage
func UpdateBookingInStorage(updatedBooking *models.Booking) (*models.Booking, error) {
	// Verify if there are classes in that day
	if classInDate := classes.Find(time.Time(updatedBooking.Date)); classInDate == nil {
		return nil, fmt.Errorf("There are no classes in this date")
	}
	
	// Updates the booking
	if _, exists := bookings[updatedBooking.ID]; exists {
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
