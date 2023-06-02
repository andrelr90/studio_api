package repositories

import (
	"testing"
	"time"
	"reflect"

	"studio_api_project/main/models"
)

func tearDownClassTests() {
	lastID = -1
	classes = NewClassesStructure()
}

func TestPopulateClassesWithExamples(t *testing.T) {
	PopulateClassesWithExamples()

	// Verify the size of the classes slice
	if len(classes.classes) != 2 {
		t.Errorf("Expected classes map size to be 2, got: %d", len(classes.classes))
	}

	// Verify the value of the lastID
	if lastID != 1 {
		t.Errorf("Expected lastID to be 1, got: %d", lastID)
	}

	tearDownClassTests()
}

func TestGetClasses(t *testing.T) {
	// Setup test data
	classDate := time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC)
	class := models.Class{
		ID:        0,
		Name:      "Pilates",
		StartDate: models.DailyDate(classDate),
		EndDate:   models.DailyDate(classDate.AddDate(0, 0, 3)),
		Capacity:  30,
	}
	classes = &ClassesStructure{}
	classes.Insert(class)
	lastID = 1
	
	// Call the GetClasses function
	result := GetClasses()

	// Check if the returned classes match the expected classes
	expectedClasses := []models.Class{0: class}
	if !reflect.DeepEqual(result, expectedClasses) {
		t.Errorf("GetClasses returned unexpected classes. ")
	}

	tearDownClassTests()
}

func TestGetClass(t *testing.T) {
	// Prepare test data
	classDate := time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC)
	class := models.Class{
		ID:        0,
		Name:      "Pilates",
		StartDate: models.DailyDate(classDate),
		EndDate:   models.DailyDate(classDate.AddDate(0, 0, 3)),
		Capacity:  30,
	}
	classes = &ClassesStructure{}
	classes.Insert(class)
	lastID = 1

	// Test case 1: Class exists
	id := "0"
	expectedClass := &class
	result := GetClass(id)
	if result.ID != expectedClass.ID {
		t.Errorf("GetClass(%s) returned an unexpected class. Got %+v, expected %+v", id, result, expectedClass)
	}

	// Test case 2: Class doesn't exist
	id = "3"
	expectedClass = nil
	result = GetClass(id)
	if result != expectedClass {
		t.Errorf("GetClass(%s) returned an unexpected class. Got %+v, expected %+v", id, result, expectedClass)
	}

	tearDownClassTests()
}

func TestCreateClass(t *testing.T) {
	// Setup test data
	classDate := time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)
	classes = NewClassesStructure()

	// Create a class with the test date
	class := models.Class{
		ID:        0,
		Name:      "Pilates",
		StartDate: models.DailyDate(classDate),
		EndDate:   models.DailyDate(classDate.AddDate(0, 0, 3)),
		Capacity:  30,
	}

	// Call the CreateClass function
	result := CreateClass(class)

	// Check if the class was created successfully
	if result == nil {
		t.Error("CreateClass failed to create the class")
	}

	// Check if the class was assigned a new ID
	if result.ID != 0 {
		t.Errorf("CreateClass assigned an unexpected ID. Got %d, expected 0", result.ID)
	}

	// Check if the class is added to the classes slice
	classes := GetClasses()
	if len(classes) != 1 {
		t.Errorf("CreateClass failed to add the class to the classes slice. Got %d classes, expected 1", len(classes))
	}
	if !reflect.DeepEqual(*result, classes[0]) {
		t.Errorf("CreateClass added the class with incorrect details. Got %+v, expected %+v", classes[0], *result)
	}

	tearDownClassTests()
}

func TestDeleteClass(t *testing.T) {
	// Prepare test data
	classDate := time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC)
	class := models.Class{
		ID:        0,
		Name:      "Pilates",
		StartDate: models.DailyDate(classDate),
		EndDate:   models.DailyDate(classDate.AddDate(0, 0, 3)),
		Capacity:  30,
	}
	classes = &ClassesStructure{}
	classes.Insert(class)
	lastID = 1

	// Test case 1: Class exists
	id := "0"
	err := DeleteClass(id)
	if err != nil {
		t.Errorf("DeleteClass(%s) should delete an existing class", id)
	}

	// Test case 2: Class doesn't exist
	err = DeleteClass(id)
	if err == nil || err.Error() != "Class not found" {
		t.Errorf("DeleteClass(%s) should not delete a class that doesn't exist", id)
	}

	tearDownClassTests()
}

func TestUpdateClass(t *testing.T) {
	// Prepare test data
	classDate := time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC)
	class := models.Class{
		ID:        0,
		Name:      "Pilates",
		StartDate: models.DailyDate(classDate),
		EndDate:   models.DailyDate(classDate.AddDate(0, 0, 3)),
		Capacity:  30,
	}
	classes = &ClassesStructure{}
	classes.Insert(class)
	lastID = 1

	// Test case 1: Class exists
	id := "0"
	_, err := UpdateClassInStorage(&class)
	if err != nil {
		t.Errorf("UpdateClass(%s) should update an existing class", id)
	}

	// Create a class that is not registered in the classes
	classDate2 := time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC)
	class2 := models.Class{
		ID:        1,
		Name:      "Pilates2",
		StartDate: models.DailyDate(classDate2),
		EndDate:   models.DailyDate(classDate2.AddDate(0, 0, 3)),
		Capacity:  30,
	}

	// Test case 2: Class doesn't exist
	_, err = UpdateClassInStorage(&class2)
	if err == nil || err.Error() != "Class not found" {
		t.Errorf("UpdateClass(%s) should not update a class that doesn't exist", id)
	}

	tearDownClassTests()
}

func TestResetClasses(t *testing.T) {
	// Create a class
	classDate := time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)
	class := models.Class{
		ID:        0,
		Name:      "Pilates",
		StartDate: models.DailyDate(classDate),
		EndDate:   models.DailyDate(classDate.AddDate(0, 0, 3)),
		Capacity:  30,
	}
	classes = &ClassesStructure{}
	classes.Insert(class)
	lastID = 0

	// Call the ResetClasses function
	ResetClasses()

	// Verify the state of the classes map and lastID
	if len(classes.classes) != 0 {
		t.Errorf("Expected classes map to be empty, got length: %d", len(classes.classes))
	}

	if lastID != -1 {
		t.Errorf("Expected lastID to be -1, got: %d", lastID)
	}

	tearDownClassTests()
}