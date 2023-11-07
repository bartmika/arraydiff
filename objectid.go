package arraydiff

import "go.mongodb.org/mongo-driver/bson/primitive"

// ObjectIDs function takes the original array and compare it with the change array to return three arrays: (1) Array with new values added to original array. (2) Array with values not changed between original and change array. (3) Array with values removed from original array.
func ObjectIDs(originalArr []primitive.ObjectID, changeArr []primitive.ObjectID) (addedArr, sameArr, removedArr []primitive.ObjectID) {
	var found bool = false
	// Run the following double-loops to find the items that were removed or
	// stayed the same.
	for _, original := range originalArr {
		for _, change := range changeArr {
			if original == change {
				found = true
				break
			}
		}
		if found {
			sameArr = append(sameArr, original)
		} else {
			removedArr = append(removedArr, original)
		}
		found = false
	}

	// Run the following double-loops to find the items that were added.
	for _, change := range changeArr {
		for _, original := range originalArr {
			if original == change {
				found = true
				break
			}
		}
		if !found {
			addedArr = append(addedArr, change)
		}
		found = false
	}
	return addedArr, sameArr, removedArr
}
