package main

/*
Adapted from https://pliutau.com/go-lang-test
Order of initialization of package variables.
Ref: https://golang.org/ref/spec#Package_initialization
*/
import "fmt"

var (
	a = c + b
	b = B()
	c = C()
	d = D()
	e = 1
)

func B() int {
	fmt.Println("B")
	return c
}

func C() int {
	fmt.Println("C")
	return d
}

func D() int {
	fmt.Println("D")
	return e
}

func main() {
	fmt.Println("Hello, World.")
}

/*
Answer:
D
C
B
Hello, World.
*/
