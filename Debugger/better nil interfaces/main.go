package main

import "fmt"

func main() {
	var i interface{}
	fmt.Println(i == nil) // output: true
	var p *int
	i = p
	fmt.Println(i == nil) // output: false
}
