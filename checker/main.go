package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	checker "checker/ressources"
)

// Program entry point
func main() {
	if len(os.Args) < 2 {
		return
	}

	args := os.Args[1:]
	stackA := []int{}

	for _, arg := range strings.Split(args[0], " ") {
		num, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error")
			os.Exit(1)
		}
		stackA = append(stackA, num)
	}
	if checker.IsDuplicate(stackA) {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	data := &checker.Data{
		StackA: stackA,
		StackB: []int{},
	}
	var instructions []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		instruction := scanner.Text()
		if instruction == "" {
			break
		}
		instructions = append(instructions, instruction)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, instruction := range instructions {
		switch instruction {
		case "sa":
			checker.SwapA(data)
		case "sb":
			checker.SwapB(data)
		case "ss":
			checker.SwapAB(data)
		case "ra":
			checker.RotateA(data)
		case "rb":
			checker.RotateB(data)
		case "rr":
			checker.RotateAB(data)
		case "rra":
			checker.RevRotateA(data)
		case "rrb":
			checker.RevRotateB(data)
		case "rrr":
			checker.RevRotateAB(data)
		case "pa":
			checker.PushA(data)
		case "pb":
			checker.PushB(data)
		default:
			fmt.Fprintln(os.Stderr, "Error")
			return
		}
	}
	if checker.IsSorted(data.StackA) && len(data.StackB) == 0 {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}
