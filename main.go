package main

import "fmt"


func modifyValue (val *int) {
	*val = 50
}

func main() {
	x:=60

	fmt.Println("Before:", x)

	modifyValue(&x)
	fmt.Println("After:", x)
}