package repositories

import (
	"time"
    "sync"
    "strconv"
	"fmt"

	"studio_api_project/main/models"
)

var (
	classes    *ClassesStructure
	lastID     int = -1
	idMutex    sync.Mutex
)

func PopulateClassesWithExamples() {
	// Add sample classes
	classes = NewClassesStructure()

	classes.Insert(models.Class{
		ID:        0,
		Name:      "Pilates",
		StartDate: models.DailyDate(time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)),
		EndDate:   models.DailyDate(time.Date(2023, time.February, 1, 0, 0, 0, 0, time.UTC)),
		Capacity:  30,
	})
	classes.Insert(models.Class{
		ID:        1,
		Name:      "Yoga",
		StartDate: models.DailyDate(time.Date(2023, time.February, 2, 0, 0, 0, 0, time.UTC)),
		EndDate:   models.DailyDate(time.Date(2023, time.March, 1, 0, 0, 0, 0, time.UTC)),
		Capacity:  25,
	})
    lastID = 1
}

func GetClasses() []models.Class {
	return classes.classes
}

func GetClass(id string) *models.Class {
	for _, class := range classes.classes {
		if strconv.Itoa(class.ID) == id {
			return &class
		}
	}
	return nil
}

func CreateClass(class models.Class) *models.Class {
	// Generate a new ID by incrementing the last ID
	idMutex.Lock()
	lastID++
	class.ID = lastID
	idMutex.Unlock()

	// Add the class to the slice
	classes.Insert(class)

	// Returns the class with its id
	return &class
}

func DeleteClass(id string) error {
	// for index, class := range classes {
	// 	if strconv.Itoa(class.ID) == id {
	// 		classes = append(classes[:index], classes[index+1:]...)
	// 		return nil
	// 		break
	// 	}
	// }
	idInt, _ := strconv.Atoi(id)
	result := classes.Remove(idInt)
	if (result == true) {
		return nil
	} else {
		return fmt.Errorf("Class not found")
	}
}

// UpdateClassInStorage updates a class in the storage
func UpdateClassInStorage(updatedClass *models.Class) (*models.Class, error) {
	// As the list is sorted, updates are done by removing and reinserting the class in the list
	removeResult := classes.Remove(updatedClass.ID)
	if (removeResult != true) {
		return nil, fmt.Errorf("Class not found")
	}
	classes.Insert(*updatedClass)
	return updatedClass, nil
	
	// for i, class := range classes.classes {
	// 	if class.ID == updatedClass.ID {
	// 		// Update the properties of the found class with the new values
	// 		classes.classes[i] = *updatedClass
	// 		return &classes[i], nil
	// 	}
	// }

	// If no class with the matching ID is found, return an error
	// return nil, fmt.Errorf("Class not found")
}

func ResetClasses() {
	// This function is used mostly for tests
	classes = NewClassesStructure()
	lastID = -1
}



// --------------------------------------
// --------- Custom Validators ----------
// --------------------------------------

// ValidateIntersection checks if there is a class within the given timeframe of a new class
func ValidateIntersection(newClass models.Class) error {
	start := time.Time(newClass.StartDate)
	end   := time.Time(newClass.EndDate)
	for _, class := range classes.classes {
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