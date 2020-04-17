package main

import "fmt"
import "example.com/api/greet"

func PrettyPrint(g greet.Greeter, s greet.Sayer, r string) {
	fmt.Println("===", greet.Greet(g, s, r), "===")
}
