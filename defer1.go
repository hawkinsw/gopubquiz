package main

/*
Taken from https://pliutau.com/go-lang-test/
The order of execution of defer'd executions.
Ref:  https://golang.org/ref/spec#Defer_statements
*/
import "fmt"

func main() {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}

/*
Answer:
3210
*/
