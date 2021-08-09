package  main

import "fmt"

func strings(){
	//anotherSlice := []int {1,2,3,4,5}
	//x := append(anotherSlice, 20,30,40)
	//fmt.Println(x)
	var makex = make([] int, 10, 100)
	    makex[9] = 9
	fmt.Println(cap(makex))
	fmt.Println(len(makex))
	fmt.Println(makex)
}
var shift = 10

var	new = shift << 3

const (
	a = iota * 100
	b = iota * 100
	c = iota
)
func main(){
	//var x uint8 = 255
	//fmt.Println(runtime.GOOS) Operating system used
	//fmt.Println(runtime.GOARCH) system architecture
	strings()
	//fmt.Println(shift)
	//fmt.Printf("%b", new)
}