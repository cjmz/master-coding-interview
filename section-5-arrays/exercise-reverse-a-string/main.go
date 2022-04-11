package main

import "fmt"

func main() {
	const str = "Hi my name is Claudio"
	reverse := make([]byte, len(str))
	for i := len(str) - 1; i >= 0; i-- {
		reverse = append(reverse, str[i])
	}
	fmt.Println(string(reverse))
}
