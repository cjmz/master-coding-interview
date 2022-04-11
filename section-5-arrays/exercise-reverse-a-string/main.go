package main

import "fmt"

func main() {
	const str = "Hi my name is Claudio"

	fmt.Println(reverse(str))
}

func reverse(str string) string {
	reverse := make([]byte, len(str))
	for i := len(str) - 1; i >= 0; i-- {
		reverse = append(reverse, str[i])
	}
	return string(reverse)

}
