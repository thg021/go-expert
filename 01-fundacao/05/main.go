package main

import "fmt"

const a = "Hello world 😎"

var (
	b bool // default é false
	c int  // 0
	d string
	e float64
)

func main() {
	var myArray [3]int

	myArray[0] = 1
	myArray[1] = 3
	myArray[2] = 9

	fmt.Println(len(myArray))
	fmt.Println(myArray[len(myArray)-1]) // ultimo item

	for i := 0; i < len(myArray); i++ {
		fmt.Printf("O valor do elemento %d é %d\n", i, myArray[i])
	}
}
