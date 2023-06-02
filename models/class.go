package models

// -------------------------------
// ------- Class structure -------
// -------------------------------

type Class struct {
	ID        int        `json:"id"`
	Name      string     `json:"name" binding:"required"`
	StartDate DailyDate  `json:"start_date" binding:"required,ltefield=EndDate" time_format:"2006-01-02" time_utc:"1"`
	EndDate   DailyDate  `json:"end_date" binding:"required" time_format:"2006-01-02" time_utc:"1"`
	Capacity  int        `json:"capacity" binding:"required,gte=1"`

	Bookings map[int]int `json:"-"`
}

func NewClass(ID int, Name string, StartDate DailyDate, EndDate DailyDate, Capacity int) *Class {
	return &Class{
		ID:        ID,
		Name:      Name,
		StartDate: StartDate,
		EndDate:   EndDate,
		Capacity:  Capacity,
		Bookings:  make(map[int]int),
	}
}
