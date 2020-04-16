package main

import "fmt"
import "example.com/api/greet"

func PrettyPrint(g greet.Greeter, s greet.Sayer) {
	fmt.Println("===", greet.Greet(g, s), "===")
}
