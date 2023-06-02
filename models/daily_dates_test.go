package models

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDailyDate_MarshalJSON(t *testing.T) {
	// Create a valid date
	date := DailyDate(time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC))

	data, err := json.Marshal(date)

	assert.NoError(t, err)
	assert.Equal(t, `"2023-01-15"`, string(data))
}

func TestDailyDate_UnmarshalJSON(t *testing.T) {
	// Create a valid date
	data := []byte(`"2023-01-15"`)

	var date DailyDate

	err := json.Unmarshal(data, &date)

	// Create the date with the traditional time.Date
	expectedDate := DailyDate(time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC))

	assert.NoError(t, err)
	assert.True(t, date.Equal(expectedDate))
}

func TestDailyDate_UnmarshalInvalidDateFormat(t *testing.T) {
	// Create an invalid date (it contains a T in the end)
	data := []byte(`"2023-01-15T"`)

	var date DailyDate

	err := json.Unmarshal(data, &date)

	assert.Error(t, err)
}

func TestDailyDate_UnmarshalInvalidQuotes(t *testing.T) {
	// Create an invalid date that can't be unquoted
	data := []byte(`0`)

	var date DailyDate

	err := json.Unmarshal(data, &date)

	assert.Error(t, err)
}

func TestDailyDate_Equal(t *testing.T) {
	// Create two equals and one different date
	date1 := DailyDate(time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC))
	date2 := DailyDate(time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC))
	date3 := DailyDate(time.Date(2023, 1, 16, 0, 0, 0, 0, time.UTC))

	// Check equality
	assert.True(t, date1.Equal(date2))
	assert.False(t, date1.Equal(date3))
}

func TestDailyDate_Hash(t *testing.T) {
	// Create a DailyDate
	date := DailyDate(time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC))

	// Get the hash of the same date date used before
	expectedHash := uint64(time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC).UnixNano())

	// Compare the hash of the DailyDate with the one of the time.Date based on uint64
	assert.Equal(t, expectedHash, date.Hash())
}