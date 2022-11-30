package diffkit

func Uints64(originalArr []uint64, changeArr []uint64) (addedArr, sameArr, removedArr []uint64) {
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
