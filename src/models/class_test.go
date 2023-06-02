package models

import (
	"testing"
	"time"
)

func TestNewClass(t *testing.T) {
	// Create test data
	ID := 1
	Name := "Pilates"
	StartDate := DailyDate(time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC))
	EndDate := DailyDate(time.Date(2023, time.February, 1, 0, 0, 0, 0, time.UTC))
	Capacity := 30

	// Call the NewClass function
	class := NewClass(ID, Name, StartDate, EndDate, Capacity)

	// Check if the class fields are set correctly
	if class.ID != ID {
		t.Errorf("Expected ID: %d, got: %d", ID, class.ID)
	}

	if class.Name != Name {
		t.Errorf("Expected Name: %s, got: %s", Name, class.Name)
	}

	if class.StartDate != StartDate {
		t.Errorf("Unexpected StartDate")
	}

	if class.EndDate != EndDate {
		t.Errorf("Unexpected EndDate")
	}

	if class.Capacity != Capacity {
		t.Errorf("Expected Capacity: %d, got: %d", Capacity, class.Capacity)
	}

	// Check if the Bookings field is initialized as an empty map
	if class.Bookings == nil {
		t.Error("Expected Bookings to be initialized as a map, got nil")
	}

	if len(class.Bookings) != 0 {
		t.Errorf("Expected Bookings length: 0, got: %d", len(class.Bookings))
	}
}