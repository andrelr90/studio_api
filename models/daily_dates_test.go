package models

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDailyDate_MarshalJSON(t *testing.T) {
	date := DailyDate(time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC))

	data, err := json.Marshal(date)

	assert.NoError(t, err)
	assert.Equal(t, `"2023-01-15"`, string(data))
}

func TestDailyDate_UnmarshalJSON(t *testing.T) {
	data := []byte(`"2023-01-15"`)

	var date DailyDate

	err := json.Unmarshal(data, &date)

	expectedDate := DailyDate(time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC))

	assert.NoError(t, err)
	assert.True(t, date.Equal(expectedDate))
}

func TestDailyDate_UnmarshalInvalidDateFormat(t *testing.T) {
	data := []byte(`"2023-01-15T"`)

	var date DailyDate

	err := json.Unmarshal(data, &date)

	assert.Error(t, err)
}

func TestDailyDate_UnmarshalInvalidQuotes(t *testing.T) {
	data := []byte(`0`)

	var date DailyDate

	err := json.Unmarshal(data, &date)

	assert.Error(t, err)
}

func TestDailyDate_Equal(t *testing.T) {
	date1 := DailyDate(time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC))
	date2 := DailyDate(time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC))
	date3 := DailyDate(time.Date(2023, 1, 16, 0, 0, 0, 0, time.UTC))

	assert.True(t, date1.Equal(date2))
	assert.False(t, date1.Equal(date3))
}

func TestDailyDate_Hash(t *testing.T) {
	date := DailyDate(time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC))

	expectedHash := uint64(time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC).UnixNano())

	assert.Equal(t, expectedHash, date.Hash())
}