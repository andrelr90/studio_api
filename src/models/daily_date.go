package models

import (
	"time"
    "strconv"
)

type DailyDate time.Time

func (date DailyDate) MarshalJSON() ([]byte, error) {
	t := time.Time(date)
	formattedDate := t.Format("2006-01-02")
	return []byte(`"` + formattedDate + `"`), nil
}

func (date *DailyDate) UnmarshalJSON(data []byte) error {
	// Strip the surrounding quotes from the JSON string
	unquotedData, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}

	// Parse the string into a time.Time value using the desired format
	t, err := time.Parse("2006-01-02", unquotedData)
	if err != nil {
		return err
	}

	// Assign the parsed time value to the DailyDate pointer
	*date = DailyDate(t)

	return nil
}

func (date DailyDate) Equal(other DailyDate) bool {
	return time.Time(date).Equal(time.Time(other))
}

func (date DailyDate) Hash() uint64 {
	return uint64(time.Time(date).UnixNano())
}
