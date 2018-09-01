package main

/*
init is a function that is defined but not declared.
Ref: https://golang.org/ref/spec#Package_initialization
*/
import "fmt"

func init() {
	fmt.Println("I am init()")
}

func i() {
	fmt.Println("I am i()")
}

func runner(r func()) {
	r()
}

func main() {
	runner(i)
	runner(init)
}
