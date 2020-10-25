package main 

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13
	m["k3"] = 14
	m["k4"] = 16


	fmt.Println("maps : ",m)
	
	v1 := m["k1"]
	fmt.Println("v1 : ",v1)
	v2 := m["k2"]
	fmt.Println("v2 : ",v2)
	v3 := m["k3"]
	fmt.Println("v3 : ",v3)
	v4 := m["k4"]
	fmt.Println("v4 : ",v4)
	v5 := m["k5"]
	fmt.Println("v5 : ",v5)

	fmt.Println("len : ",len(m))

	delete(m,"k2")
	fmt.Println("maps : ",m)

	_, prs := m["k2"]
	fmt.Println("prs : ",prs)
	
	_, prsm := m["k3"]
	fmt.Println("prs : ",prsm)

	n := map[string]int{"foo" : 1,"bar" : 2}
	fmt.Println("map : ", n)

}
