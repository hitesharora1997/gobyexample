package main

import "fmt"

func main() {
	s := make([]string,3)
	fmt.Println("emp : ",s)


	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set : ",s)
	fmt.Println("get : ",s[0])
	fmt.Println("get : ",s[1])
	fmt.Println("get : ",s[2])

	
	fmt.Println("get : ",len(s))

	s = append(s,"d")
	s = append(s,"e","f")
	s = append(s,"z","m","p")
	fmt.Println("apd :",s)

	fmt.Println("get : ",len(s))

	c := make([]string,len(s))
	copy(c,s)
	fmt.Println("cpy : ",c)

	l := s[5:9]
	fmt.Println("sl1 : ", l)

	l = s[:7]
	fmt.Println("sl2 : " ,l)

	l = s[1:]
	fmt.Println("sl3 : ", l)

	t := []string{"g","h","i","y"}
	fmt.Println("dcl : ",t)

	twoD := make([][]int ,3)
	for i := 0 ; i < 3 ; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0 ; j < innerLen ; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d : ",twoD)
}
