package main

/*
Adapted from https://pliutau.com/go-lang-test
Different types cannot be 'combined' without explicit
casting.
Ref: https://blog.golang.org/constants
*/
import "fmt"

func main() {
	var a int8 = 3
	var b int16 = 4

	sum := a + b

	fmt.Println(sum)
}
