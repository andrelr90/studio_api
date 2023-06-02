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

func (s *ClassesStructure) Insert(class models.Class) {
	// Find the index to insert the class using binary search
	index := sort.Search(len(s.classes), func(i int) bool {
		return time.Time(class.StartDate).Before(time.Time(s.classes[i].StartDate))
	})

	// Insert the class at the found index
	s.classes = append(s.classes[:index], append([]models.Class{class}, s.classes[index:]...)...)
}

func (s *ClassesStructure) Remove(classID int) bool {
	// Find the index of the class with the given ID
	for i, class := range s.classes {
		if class.ID == classID {
			// Remove the class from the slice
			for bookingId, _ := range class.Bookings {
				DeleteBooking(strconv.Itoa(bookingId))
			}
			s.classes = append(s.classes[:i], s.classes[i+1:]...)
			return true
		}
	}

	// Class with the given ID not found
	return false
}

func (s *ClassesStructure) Find(date time.Time) *models.Class {
	// Perform binary search to find the class on the given date
	index := sort.Search(len(s.classes), func(i int) bool {
		return !date.After(time.Time(s.classes[i].EndDate))
	})

	// Check if the found class contains the given date
	if (index < len(s.classes) && !date.Before(time.Time(s.classes[index].StartDate))) {
		return &s.classes[index]
	}

	// No class found on the given date
	return nil
}
