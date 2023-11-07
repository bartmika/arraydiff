package arraydiff

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestObjectIDsCase1(t *testing.T) {
	one, _ := primitive.ObjectIDFromHex("6541c91defc6cd8f2cb66e2e")
	two, _ := primitive.ObjectIDFromHex("6541c91defc6cd8f2cb66e30")
	three, _ := primitive.ObjectIDFromHex("6541c91defc6cd8f2cb66e32")
	four, _ := primitive.ObjectIDFromHex("6541c91defc6cd8f2cb66e34")
	five, _ := primitive.ObjectIDFromHex("6541c91defc6cd8f2cb66e36")
	six, _ := primitive.ObjectIDFromHex("6541c91defc6cd8f2cb66e38")
	seven, _ := primitive.ObjectIDFromHex("6541c91defc6cd8f2cb66e3a")
	eight, _ := primitive.ObjectIDFromHex("6541c91defc6cd8f2cb66e3c")
	nine, _ := primitive.ObjectIDFromHex("6541c91defc6cd8f2cb66e3e")
	ten, _ := primitive.ObjectIDFromHex("6541c91defc6cd8f2cb66e40")

	// Sample data.
	oldArr := []primitive.ObjectID{one, two, three, four, five}
	newArr := []primitive.ObjectID{one, two, six, seven, eight, nine, ten}

	// Run our test.
	actualAddedArr, actualSameArr, actualRemovedArr := ObjectIDs(oldArr, newArr)

	// Correct results.
	expectedAddedArr := []primitive.ObjectID{six, seven, eight, nine, ten}
	expectedSameArr := []primitive.ObjectID{one, two}
	expectedRemovedArr := []primitive.ObjectID{three, four, five}

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

func TestObjectIDsCase2(t *testing.T) {
	one, _ := primitive.ObjectIDFromHex("6541c91defc6cd8f2cb66e2e")
	two, _ := primitive.ObjectIDFromHex("6541c91defc6cd8f2cb66e30")
	three, _ := primitive.ObjectIDFromHex("6541c91defc6cd8f2cb66e32")
	four, _ := primitive.ObjectIDFromHex("6541c91defc6cd8f2cb66e34")
	five, _ := primitive.ObjectIDFromHex("6541c91defc6cd8f2cb66e36")

	// Sample data.
	oldArr := []primitive.ObjectID{one, two, three, four, five}
	newArr := []primitive.ObjectID{}

	// Run our test.
	actualAddedArr, actualSameArr, actualRemovedArr := ObjectIDs(oldArr, newArr)

	// Correct results.
	expectedAddedArr := []primitive.ObjectID{}
	expectedSameArr := []primitive.ObjectID{}
	expectedRemovedArr := []primitive.ObjectID{one, two, three, four, five}

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

func TestObjectIDsCase3(t *testing.T) {
	one, _ := primitive.ObjectIDFromHex("6541c91defc6cd8f2cb66e2e")
	two, _ := primitive.ObjectIDFromHex("6541c91defc6cd8f2cb66e30")
	three, _ := primitive.ObjectIDFromHex("6541c91defc6cd8f2cb66e32")
	four, _ := primitive.ObjectIDFromHex("6541c91defc6cd8f2cb66e34")
	five, _ := primitive.ObjectIDFromHex("6541c91defc6cd8f2cb66e36")

	// Sample data.
	oldArr := []primitive.ObjectID{}
	newArr := []primitive.ObjectID{one, two, three, four, five}

	// Run our test.
	actualAddedArr, actualSameArr, actualRemovedArr := ObjectIDs(oldArr, newArr)

	// Correct results.
	expectedAddedArr := []primitive.ObjectID{one, two, three, four, five}
	expectedSameArr := []primitive.ObjectID{}
	expectedRemovedArr := []primitive.ObjectID{}

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

func TestObjectIDsCase4(t *testing.T) {
	// Sample data.
	oldArr := []primitive.ObjectID{}
	newArr := []primitive.ObjectID{}

	// Run our test.
	actualAddedArr, actualSameArr, actualRemovedArr := ObjectIDs(oldArr, newArr)

	// Correct results.
	expectedAddedArr := []primitive.ObjectID{}
	expectedSameArr := []primitive.ObjectID{}
	expectedRemovedArr := []primitive.ObjectID{}

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
