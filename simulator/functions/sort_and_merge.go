package simulator

// Performs merge on the array of minicons while removing duplicates
// First array is the parent array (so no duplicates in it, but is being checked to be sure)
func mergeTwoArraysWithoutDuplicates(Array1 []int, Array2 []int) (MergedArray []int) {
	duplicatesSet := make(map[int]bool, 0)
	for _, Object := range Array1 {
		if duplicatesSet[Object] {
			continue
		} else {
			MergedArray = append(MergedArray, Object)
			duplicatesSet[Object] = true
		}
	}
	for _, Object := range Array2 {
		if duplicatesSet[Object] {
			continue
		} else {
			MergedArray = append(MergedArray, Object)
			duplicatesSet[Object] = true
		}
	}
	return MergedArray
}

func MergeAllDKHMiniconObjects(
	ParentDKHMiniconObject DKHMinicons,
	OtherObjects ...DKHMinicons) (NewDKHMiniconObject DKHMinicons) {

	for _, DKHMiniconObject := range OtherObjects {

		NewDKHMiniconObject.DeadMinicons = mergeTwoArraysWithoutDuplicates(
			ParentDKHMiniconObject.DeadMinicons,
			DKHMiniconObject.DeadMinicons)

		NewDKHMiniconObject.KillMinicons = mergeTwoArraysWithoutDuplicates(
			ParentDKHMiniconObject.KillMinicons,
			DKHMiniconObject.KillMinicons)

		NewDKHMiniconObject.HurtMinicons = mergeTwoArraysWithoutDuplicates(ParentDKHMiniconObject.HurtMinicons,
			DKHMiniconObject.HurtMinicons)
	}

	return NewDKHMiniconObject
}
