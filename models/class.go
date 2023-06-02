package models

import (
	"time"
    "fmt"
)

// -------------------------------
// ------- Class structure -------
// -------------------------------

type Class struct {
	ID        int        `json:"id"`
	Name      string     `json:"name" binding:"required"`
	StartDate DailyDate  `json:"start_date" binding:"required,ltefield=EndDate" time_format:"2006-01-02" time_utc:"1"`
	EndDate   DailyDate  `json:"end_date" binding:"required" time_format:"2006-01-02" time_utc:"1"`
	Capacity  int        `json:"capacity" binding:"required,gte=1"`
}



// --------------------------------------
// --------- Custom Validators ----------
// --------------------------------------

// ValidateIntersection checks if there is a class within the given timeframe of a new class
func ValidateIntersection(classes []Class, newClass Class) error {
	start := time.Time(newClass.StartDate)
	end   := time.Time(newClass.EndDate)
	for _, class := range classes {
		// Check if there is an intersection between the given timeframe and the existing class:
		if (class.ID != newClass.ID) {
			if (start.Before(time.Time(class.EndDate)) || start.Equal(time.Time(class.EndDate))) && 
			   (end.After(time.Time(class.StartDate)) || end.Equal(time.Time(class.StartDate))) {
				var errorMessage = fmt.Sprintf("Intersection found with %s", class.Name)
				return fmt.Errorf(errorMessage)
			}
		}
	}
	// No intersection found:
	return nil
}
