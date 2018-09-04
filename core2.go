package main

/*
Adapted from https://pliutau.com/go-lang-test
Variables can be of different types in the same
declaration.
*/
import "fmt"

func main() {
	var a, b, c = 3, 4, "c"
	fmt.Println(a, b, c)
}

/*
Answer:
3 4 c
*/
