package checker

// Swap top two elements of stack A
func SwapA(data *Data) {
    if len(data.StackA) < 2 {
        return
    }
    data.StackA[0], data.StackA[1] = data.StackA[1], data.StackA[0]
}

// Swap top two elements of stack B
func SwapB(data *Data) {
    if len(data.StackB) < 2 {
        return
    }
    data.StackB[0], data.StackB[1] = data.StackB[1], data.StackB[0]
}

// Swap top elements of both stacks
func SwapAB(data *Data) {
    SwapA(data)
    SwapB(data)
}
