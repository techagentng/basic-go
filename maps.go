package  main

import "fmt"

func main(){
	m := map[string]int {
		"Nnamdi": 34,
		"Ogechi": 23,
		"Chiamaka": 43,
	}
	fmt.Println(m["Chiamaka"])
	v, ok := m["Nnamdi"]
	fmt.Println(v)
	fmt.Println(ok)
}