package main

import "fmt"

func main() {
	fmt.Printf("The expression %s has matching brackets: %v\n", "([])(){}(())()()", matchingBrackets("([])(){}(())()()"))
}

func matchingBrackets(s string) bool {
	stack := []rune{}
	brackets := map[rune]rune{rune('{'): rune('}'), rune('['): rune(']'), rune('('): rune(')')}
	closing := map[rune]bool{rune('}'): true, rune(']'): true, rune(')'): true}
	idx := 0
	for _, r := range s {
		fmt.Printf("Matching %s %d\n", string(r), r)
		if val, ok := brackets[r]; ok {
			stack = append(stack, val)
			idx++
		} else if closing[r] {
			fmt.Println("Closing bracket")
			if idx >= 0 && r == stack[idx-1] {
				stack = stack[:idx-1]
				idx--
				fmt.Printf("Matching bracket found and idx is %d\n", idx)
			} else {
				return false
			}
		}
		continue
	}
	return idx == 0
}

// 1. k most frequent elements in array
// 2. array with sorted intervals, output array with overlapping intervals merged
