package diffkit

import (
	"reflect"
	"testing"
)

func TestUints64Case1(t *testing.T) {
	// Sample data.
	oldArr := []uint64{1, 2, 3, 4, 5}
	newArr := []uint64{1, 2, 6, 7, 8, 9, 10}

	// Run our test.
	actualAddedArr, actualSameArr, actualRemovedArr := Uints64(oldArr, newArr)

	// Correct results.
	expectedAddedArr := []uint64{6, 7, 8, 9, 10}
	expectedSameArr := []uint64{1, 2}
	expectedRemovedArr := []uint64{3, 4, 5}

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

func TestUints64Case2(t *testing.T) {
	// Sample data.
	oldArr := []uint64{1, 2, 3, 4, 5}
	newArr := []uint64{}

	// Run our test.
	actualAddedArr, actualSameArr, actualRemovedArr := Uints64(oldArr, newArr)

	// Correct results.
	expectedAddedArr := []uint64{}
	expectedSameArr := []uint64{}
	expectedRemovedArr := []uint64{1, 2, 3, 4, 5}

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

func TestUints64Case3(t *testing.T) {
	// Sample data.
	oldArr := []uint64{}
	newArr := []uint64{1, 2, 3, 4, 5}

	// Run our test.
	actualAddedArr, actualSameArr, actualRemovedArr := Uints64(oldArr, newArr)

	// Correct results.
	expectedAddedArr := []uint64{1, 2, 3, 4, 5}
	expectedSameArr := []uint64{}
	expectedRemovedArr := []uint64{}

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
