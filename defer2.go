package main

/*
Taken from https://pliutau.com/go-lang-test/
Capturing values of parameters to a defer'd function
Ref:  https://golang.org/ref/spec#Defer_statements
*/
import "fmt"

func main() {
	j := 4
	for i := 0; i < 4; i++ {
		defer func(k int) {
			fmt.Print(k)
			j--
		}(j)
	}
}

/*
Answer:
4444
*/
