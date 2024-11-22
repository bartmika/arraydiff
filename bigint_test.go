package arraydiff

import (
	"math/big"
	"reflect"
	"testing"
)

func TestBigIntsCase1(t *testing.T) {
	// Sample data.
	oldArr := []*big.Int{big.NewInt(1), big.NewInt(2), big.NewInt(3), big.NewInt(4), big.NewInt(5)}
	newArr := []*big.Int{big.NewInt(1), big.NewInt(2), big.NewInt(6), big.NewInt(7), big.NewInt(8), big.NewInt(9), big.NewInt(10)}

	// Run our test.
	actualAddedArr, actualSameArr, actualRemovedArr := BigInts(oldArr, newArr)

	// Correct results.
	expectedAddedArr := []*big.Int{big.NewInt(6), big.NewInt(7), big.NewInt(8), big.NewInt(9), big.NewInt(10)}
	expectedSameArr := []*big.Int{big.NewInt(1), big.NewInt(2)}
	expectedRemovedArr := []*big.Int{big.NewInt(3), big.NewInt(4), big.NewInt(5)}

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

func TestBigIntsCase2(t *testing.T) {
	// Sample data.
	oldArr := []*big.Int{big.NewInt(1), big.NewInt(2), big.NewInt(3), big.NewInt(4), big.NewInt(5)}
	newArr := []*big.Int{}

	// Run our test.
	actualAddedArr, actualSameArr, actualRemovedArr := BigInts(oldArr, newArr)

	// Correct results.
	expectedAddedArr := []*big.Int{}
	expectedSameArr := []*big.Int{}
	expectedRemovedArr := []*big.Int{big.NewInt(1), big.NewInt(2), big.NewInt(3), big.NewInt(4), big.NewInt(5)}

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

func TestBigIntsCase3(t *testing.T) {
	// Sample data.
	oldArr := []*big.Int{}
	newArr := []*big.Int{big.NewInt(1), big.NewInt(2), big.NewInt(3), big.NewInt(4), big.NewInt(5)}

	// Run our test.
	actualAddedArr, actualSameArr, actualRemovedArr := BigInts(oldArr, newArr)

	// Correct results.
	expectedAddedArr := []*big.Int{big.NewInt(1), big.NewInt(2), big.NewInt(3), big.NewInt(4), big.NewInt(5)}
	expectedSameArr := []*big.Int{}
	expectedRemovedArr := []*big.Int{}

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

func TestBigIntsCase4(t *testing.T) {
	// Sample data.
	oldArr := []*big.Int{}
	newArr := []*big.Int{}

	// Run our test.
	actualAddedArr, actualSameArr, actualRemovedArr := BigInts(oldArr, newArr)

	// Correct results.
	expectedAddedArr := []*big.Int{}
	expectedSameArr := []*big.Int{}
	expectedRemovedArr := []*big.Int{}

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
