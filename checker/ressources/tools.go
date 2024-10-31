package checker

type Data struct {
	StackA []int
	StackB []int
}

// Check for duplicate numbers
func IsDuplicate(slice []int) bool {
	seen := make(map[int]bool)

	for _, part := range slice {
		if seen[part] {
			return true
		}
		seen[part] = true
	}
	return false
}

// Sorting check function
func IsSorted(str []int) bool {
	for i := 0; i < len(str)-1; i++ {
		if str[i] > str[i+1] {
			return false
		}
	}
	return true
}

