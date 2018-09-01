package main

/*
Adapted from https://twitter.com/davecheney/status/508076305003196416
Initialization of arrays with composite literals.
Ref: https://golang.org/ref/spec#Composite_literals
*/
import "fmt"

func main() {
	output := []string{"a", 2: "d", "c", 1: "b", 6: "e"}
	fmt.Println(len(output))
}
