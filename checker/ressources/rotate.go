package checker

// Rotate stack A (move top element to bottom)
func RotateA(data *Data) {
	if len(data.StackA) < 2 {
		return
	}
	first := data.StackA[0]
	data.StackA = append(data.StackA[1:], first)
}

// Rotate stack B (move top element to bottom)
func RotateB(data *Data) {
	if len(data.StackB) < 2 {
		return
	}
	first := data.StackB[0]
	data.StackB = append(data.StackB[1:], first)
}

// Rotate both stacks
func RotateAB(data *Data) {
	RotateA(data)
	RotateB(data)
}

// Reverse rotate stack A (move bottom element to top)
func RevRotateA(data *Data) {
	if len(data.StackA) < 2 {
		return
	}
	last := data.StackA[len(data.StackA)-1]
	data.StackA = append([]int{last}, data.StackA[:len(data.StackA)-1]...)
}

// Reverse rotate stack B (move bottom element to top)
func RevRotateB(data *Data) {
	if len(data.StackB) < 2 {
		return
	}
	last := data.StackB[len(data.StackB)-1]
	data.StackB = append([]int{last}, data.StackB[:len(data.StackB)-1]...)
}

// Reverse rotate both stacks
func RevRotateAB(data *Data) {
	RevRotateA(data)
	RevRotateB(data)
}
