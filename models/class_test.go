package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestValidateIntersection(t *testing.T) {
	// Create a list of existing classes
	classes := []Class{
		{
			ID:        1,
			Name:      "Yoga",
			StartDate: DailyDate(time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)),
			EndDate:   DailyDate(time.Date(2023, time.January, 31, 0, 0, 0, 0, time.UTC)),
			Capacity:  20,
		},
		{
			ID:        2,
			Name:      "Pilates",
			StartDate: DailyDate(time.Date(2023, time.February, 1, 0, 0, 0, 0, time.UTC)),
			EndDate:   DailyDate(time.Date(2023, time.February, 28, 0, 0, 0, 0, time.UTC)),
			Capacity:  30,
		},
		{
			ID:        3,
			Name:      "Zumba",
			StartDate: DailyDate(time.Date(2023, time.March, 1, 0, 0, 0, 0, time.UTC)),
			EndDate:   DailyDate(time.Date(2023, time.March, 31, 0, 0, 0, 0, time.UTC)),
			Capacity:  25,
		},
	}

	// Create a new class with a conflicting timeframe
	newClass := Class{
		ID:        4,
		Name:      "Pilates2",
		StartDate: DailyDate(time.Date(2023, time.February, 15, 0, 0, 0, 0, time.UTC)),
		EndDate:   DailyDate(time.Date(2023, time.March, 15, 0, 0, 0, 0, time.UTC)),
		Capacity:  35,
	}

	// Validate the intersection
	err := ValidateIntersection(classes, newClass)

	// Check that an error is returned with the expected error message
	expectedError := "Intersection found with Pilates"
	assert.EqualError(t, err, expectedError)

	// Create a new class with a conflicting timeframe in limits (start equals to another end)
	newClass = Class{
		ID:        4,
		Name:      "Pilates2",
		StartDate: DailyDate(time.Date(2023, time.February, 28, 0, 0, 0, 0, time.UTC)),
		EndDate:   DailyDate(time.Date(2023, time.March, 15, 0, 0, 0, 0, time.UTC)),
		Capacity:  35,
	}

	// Validate the intersection
	err = ValidateIntersection(classes, newClass)

	// Check that an error is returned with the expected error message
	expectedError = "Intersection found with Pilates"
	assert.EqualError(t, err, expectedError)

	// Create a new class with a non-conflicting timeframe
	nonConflictingClass := Class{
		ID:        5,
		Name:      "Dance",
		StartDate: DailyDate(time.Date(2023, time.April, 1, 0, 0, 0, 0, time.UTC)),
		EndDate:   DailyDate(time.Date(2023, time.April, 30, 0, 0, 0, 0, time.UTC)),
		Capacity:  15,
	}

	// Validate the intersection for the non-conflicting class
	err = ValidateIntersection(classes, nonConflictingClass)

	// Check that no error is returned
	assert.NoError(t, err)
}
