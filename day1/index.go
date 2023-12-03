package main

import "fmt"

// learning go on the way
/**
 * @Description: main function
 * @param
 * @return
**/
func main() {

	type farehnheit int
	type celsius int

	var f farehnheit = 32
	var c celsius = 0

	c = celsius((f - 32) * 5 / 9)
	fmt.Println(f, c)
}
