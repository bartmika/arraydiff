package arraydiff

import (
	"reflect"
	"testing"
)

func TestStringsCase1(t *testing.T) {
	// Sample data.
	oldArr := []string{"1", "2", "3", "4", "5"}
	newArr := []string{"1", "2", "6", "7", "8", "9", "10"}

	// Run our test.
	actualAddedArr, actualSameArr, actualRemovedArr := Strings(oldArr, newArr)

	// Correct results.
	expectedAddedArr := []string{"6", "7", "8", "9", "10"}
	expectedSameArr := []string{"1", "2"}
	expectedRemovedArr := []string{"3", "4", "5"}

	// Verify.
	if reflect.DeepEqual(expectedAddedArr, actualAddedArr) == false {
		t.Errorf("incorrect added array contents, got %v but was expecting %v", actualAddedArr, expectedAddedArr)
	}
	if reflect.DeepEqual(expectedSameArr, actualSameArr) == false {
		t.Errorf("incorrect same array contents, got %v but was expecting %v", actualSameArr, expectedSameArr)
	}
	if reflect.DeepEqual(expectedRemovedArr, actualRemovedArr) == false {
		t.Errorf("incorrect removed array contents, got %v but was expecting %v", actualRemovedArr, expectedRemovedArr)
	}
}

func TestStringsCase2(t *testing.T) {
	// Sample data.
	oldArr := []string{"1", "2", "3", "4", "5"}
	newArr := []string{}

	// Run our test.
	actualAddedArr, actualSameArr, actualRemovedArr := Strings(oldArr, newArr)

	// Correct results.
	expectedAddedArr := []string{}
	expectedSameArr := []string{}
	expectedRemovedArr := []string{"1", "2", "3", "4", "5"}

	// Verify.
	if len(actualAddedArr) != 0 {
		t.Errorf("incorrect added array contents, got %v but was expecting %v", actualAddedArr, expectedAddedArr)
	}
	if len(actualSameArr) != 0 {
		t.Errorf("incorrect same array contents, got %v but was expecting %v", actualSameArr, expectedSameArr)
	}
	if reflect.DeepEqual(expectedRemovedArr, actualRemovedArr) == false {
		t.Errorf("incorrect removed array contents, got %v but was expecting %v", actualRemovedArr, expectedRemovedArr)
	}
}

func TestStringsCase3(t *testing.T) {
	// Sample data.
	oldArr := []string{}
	newArr := []string{"1", "2", "3", "4", "5"}

	// Run our test.
	actualAddedArr, actualSameArr, actualRemovedArr := Strings(oldArr, newArr)

	// Correct results.
	expectedAddedArr := []string{"1", "2", "3", "4", "5"}
	expectedSameArr := []string{}
	expectedRemovedArr := []string{}

	// Verify.
	if reflect.DeepEqual(expectedAddedArr, actualAddedArr) == false {
		t.Errorf("incorrect added array contents, got %v but was expecting %v", actualAddedArr, expectedAddedArr)
	}
	if len(actualSameArr) != 0 {
		t.Errorf("incorrect same array contents, got %v but was expecting %v", actualSameArr, expectedSameArr)
	}
	if len(actualRemovedArr) != 0 {
		t.Errorf("incorrect removed array contents, got %v but was expecting %v", actualRemovedArr, expectedRemovedArr)
	}
}

func TestStringsCase4(t *testing.T) {
	// Sample data.
	oldArr := []string{}
	newArr := []string{}

	// Run our test.
	actualAddedArr, actualSameArr, actualRemovedArr := Strings(oldArr, newArr)

	// Correct results.
	expectedAddedArr := []string{}
	expectedSameArr := []string{}
	expectedRemovedArr := []string{}

	// Verify.
	if len(actualAddedArr) != 0 {
		t.Errorf("incorrect added array contents, got %v but was expecting %v", actualAddedArr, expectedAddedArr)
	}
	if len(actualSameArr) != 0 {
		t.Errorf("incorrect same array contents, got %v but was expecting %v", actualSameArr, expectedSameArr)
	}
	if len(actualRemovedArr) != 0 {
		t.Errorf("incorrect removed array contents, got %v but was expecting %v", actualRemovedArr, expectedRemovedArr)
	}
}
