package main

import "fmt"

func main() {
	i := 0
	for i < 3 {
		fmt.Println("hello", i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println("bose", j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println("abazz", n)

	}
}
