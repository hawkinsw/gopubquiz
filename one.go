package main

/*
Adapted from https://pliutau.com/go-lang-test
*/
import "fmt"

const (
	a = iota
	b = iota
)

const (
	c, d = iota, iota + iota
)

func main() {
	fmt.Println(a, b, c, d)
}
