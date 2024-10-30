package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	var instructions []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}
	fmt.Println(instructions)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
		return
	}
	
	for _, instruction := range instructions {
		switch instruction {
		case "sa":
			sa()
		case "sb":
			sb()
		case "ss":
			ss()
		case "ra":
			ra()
		case "rb":
			rb()
		case "rr":
			rr()
		case "rra":
			rra()
		case "rrb":
			rrb()
		case "rrr":
			rrr()
		case "pa":
			pa()
		case "pb":
			pb()
		default:
			fmt.Println(os.Stderr, "Error")
			return
		}
	}
	if isSorted(stackA) {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}
