package arraydiff

import (
	"math/big"
)

// BigInt Function takes the original array and compare it with the change array to return three arrays: (1) Array with new values added to original array. (2) Array with values not changed between original and change array. (3) Array with values removed from original array.
func BigInts(originalArr []*big.Int, changeArr []*big.Int) (addedArr, sameArr, removedArr []*big.Int) {
	var found bool = false
	// Run the following double-loops to find the items that were removed or
	// stayed the same.
	for _, original := range originalArr {
		for _, change := range changeArr {
			if original.Cmp(change) == 0 {
				found = true
				break
			}
		}
		if found {
			sameArr = append(sameArr, new(big.Int).Set(original))
		} else {
			removedArr = append(removedArr, new(big.Int).Set(original))
		}
		found = false
	}

	// Run the following double-loops to find the items that were added.
	for _, change := range changeArr {
		for _, original := range originalArr {
			if change.Cmp(original) == 0 {
				found = true
				break
			}
		}
		if !found {
			addedArr = append(addedArr, new(big.Int).Set(change))
		}
		found = false
	}
	return addedArr, sameArr, removedArr
}
