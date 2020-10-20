package main

import (
	f "fmt"
	z "time"
)

func main() {
	i := 2
	f.Print(" write " , i , " as " )
	switch i {
	case 1 :
		f.Println("one")
	case 2 :
		f.Println("two")
	case 3 :
		f.Println("three")
	}

	switch z.Now().Weekday() {
	case z.Saturday, z.Sunday:
		f.Println("its the weekend")
	default:
		f.Println("its a weekday")
	}

	t := z.Now()
	switch {
	case t.Hour() < 12:
		f.Println("its the weekend")
	default:
		f.Println("its after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool :
			f.Println("")

}
