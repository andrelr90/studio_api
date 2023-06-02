package repositories

import (
	"sort"
	"time"
	"strconv"

	"studio_api_project/main/models"
)

type ClassesStructure struct {
	classes []models.Class
}

func NewClassesStructure() *ClassesStructure {
	return &ClassesStructure{}
}

func (cs *ClassesStructure) Insert(class models.Class) {
	// Find the index to insert the class using binary search
	index := sort.Search(len(cs.classes), func(i int) bool {
		return time.Time(class.StartDate).Before(time.Time(cs.classes[i].StartDate))
	})

	// Insert the class at the found index
	cs.classes = append(cs.classes[:index], append([]models.Class{class}, cs.classes[index:]...)...)
}

func (cs *ClassesStructure) Remove(classID int, cascadeAllBookings bool) bool {
	// Find the index of the class with the given ID
	for i, class := range cs.classes {
		if class.ID == classID {
			if (cascadeAllBookings == true) {
				// Cascade bookings removal
				for bookingId, _ := range class.Bookings {
					DeleteBooking(strconv.Itoa(bookingId))
				}
			}
			// Remove the class from the slice
			cs.classes = append(cs.classes[:i], cs.classes[i+1:]...)
			return true
		}
	}

	// Class with the given ID not found
	return false
}

func (cs *ClassesStructure) UpdateClass(class  *models.Class) {
	// Remove only bookings in dates that do not exist anymore
	for key, _ := range class.Bookings {
		timeDate := time.Time(bookings[key].Date)
		if time.Time(class.StartDate).After(timeDate) || time.Time(class.EndDate).Before(timeDate) {
			DeleteBooking(strconv.Itoa(key))
		}
	}

	// As the slice is sorted, updates are done by removing and reinserting the class in the slice. Cascade is not activated in this case.
	cs.Remove(class.ID, false)
	cs.Insert(*class)
}

func (cs *ClassesStructure) Find(date time.Time) *models.Class {
	// Perform binary search to find the class on the given date
	index := sort.Search(len(cs.classes), func(i int) bool {
		return !date.After(time.Time(cs.classes[i].EndDate))
	})

	// Check if the found class contains the given date
	if (index < len(cs.classes) && !date.Before(time.Time(cs.classes[index].StartDate))) {
		return &cs.classes[index]
	}

	// No class found on the given date
	return nil
}
