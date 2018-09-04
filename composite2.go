package main

/*
Taken from https://twitter.com/davecheney/status/1031389514890006528
Calling method
Ref: https://golang.org/ref/spec#Method_expressions
*/
import "fmt"

type Q struct{}

func (Q) Hola() string { return "Bueno" }

func main() {
	var q Q
	fmt.Println(Q.Hola(q))
}

/*
Answer:
Bueno
*/
