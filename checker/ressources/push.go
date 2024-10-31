package checker

// Push top element from B to A
func PushA(data *Data) {
	if len(data.StackB) == 0 {
        return
    }
    
    topElement := data.StackB[0]
    data.StackB = data.StackB[1:]
    
    data.StackA = append([]int{topElement}, data.StackA...)
}

// Push top element from A to B
func PushB(data *Data) {
	if len(data.StackA) == 0 {
        return
    }
    
    topElement := data.StackA[0]
    data.StackA = data.StackA[1:]
    
    data.StackB = append([]int{topElement}, data.StackB...)
}