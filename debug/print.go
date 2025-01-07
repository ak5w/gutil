package debug

import "fmt"

// Print any type value.
func Print(s ...any) {
	for _, v := range s {
		fmt.Printf("[type: %T, value: %v]\n", v, v)
	}
}
