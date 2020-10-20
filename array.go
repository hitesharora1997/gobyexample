package main

import "fmt"

func main() {
var a [5]int                //array of sizze 5
fmt.Println("emp : ", a)

a[4] = 100
a[1] = 12
a[3] = 153

fmt.Println("set : ",a)
fmt.Println("get : ",a[1] )
fmt.Println("get : ",a[2] )
fmt.Println("get : ",a[3] )
fmt.Println("get : ",a[4] )

for i := 0 ; i < 5; i++ {
	fmt.Printf("this value is %d \n", a[i])
}
fmt.Println("len :",len(a))

b := [5]int{1,2,3,4,5}
fmt.Println("dcl : ",b)

var twoD [2][3]int
for i := 0 ; i < 2 ; i++ {
	for j := 0 ; j < 3 ; j++ {
		twoD[i][j] = i + j
	}
}
fmt.Println("2d : ", twoD)

}
