package main

import (
	f "fmt"
)

func main() {
	if 7%2 == 0 {
		f.Println("7 is even")
	} else {
		f.Println("7 is odd")
	}

	if 8%4 == 0 {
		f.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 {
		f.Println(num, "is negative")
	} else if num < 10 {
		f.Println(num, "has 1 digit")
	} else {
		f.Println(num, "has mutiple digits")
	}

}
