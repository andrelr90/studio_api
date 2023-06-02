package models

import (
	// "time"
)

// -------------------------------
// ------ Booking structure ------
// -------------------------------

type Booking struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"  binding:"required"`
	Date      DailyDate  `json:"date"  binding:"required" time_format:"2006-01-02" time_utc:"1"`
}
