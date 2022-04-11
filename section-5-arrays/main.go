package main

import "fmt"

func main() {
	// 4x4 = 16 bytes of storage
	strings := []string{"a", "b", "c", "d"}
	fmt.Println(strings[2])

	// append item in slice (push equivalent in JS)
	strings = append(strings, "e") // O(1) = constant
	fmt.Println(strings)

	// remove the last item from slice (pop equivalent in JS)
	strings = append(strings[:len(strings)-1], strings[len(strings):]...) // O(1) = constant
	fmt.Println(strings)

	// append item as first element in array (unshift equivalent in JS)
	strings = append([]string{"X"}, strings...) // O(n) = constant
	fmt.Println(strings)

	// append item in specific position of an array (splice equivalent in JS)
	strings = append(strings[:3], strings[2:]...)
	strings[2] = "O"
	fmt.Println(strings)
}
