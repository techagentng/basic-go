package  main

import "fmt"

func main(){
	m := map[string]int {
		"Nnamdi": 34,
		"Ogechi": 23,
		"Chiamaka": 43,
	}
	fmt.Println(m["Chiamaka"])
	// A pattern to take note of for if statement
	if v, ok := m["Nnamdi"]; ok {
		fmt.Println("Its is ok boom", v)
	}
	//fmt.Println(v)
	//fmt.Println(ok)
}