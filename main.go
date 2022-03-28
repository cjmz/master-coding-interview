package main

import (
	"fmt"
	"time"
)

func main() {
	// a := [...]string{"tubarao", "nemo"}
	// a := []string{"tubarao", "nemo", "jeremias", "pereira", "artur", "joaquim", "otavio", "oie"}
	a := make([]string, 100000)
	for i := 7; i < 100000; i++ {
		// append(a, "oi")
		a[i] = "nemo"
	}

	s := time.Now()
	for _, v := range a {
		if v == "nemo" {
			fmt.Println("found nemo")
		}
	}
	e := time.Since(s)
	fmt.Printf("Time took %s\n", e)
}
