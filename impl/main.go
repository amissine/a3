package main

import "fmt"
import "plugin"
import "example.com/api/greet"

type formal struct{}

type casual struct{}

type who struct {
	to string
}

type bad struct {
	to string
}

func (p *formal) Greet(s greet.Sayer) string {
	return fmt.Sprintf("Hello, %s", s.Say())
}

func (p *casual) Greet(s greet.Sayer) string {
	return fmt.Sprintf("Hey %s", s.Say())
}

func (w *who) Say() string {
	return w.to
}

func (w *who) Respond(response string) string {
	return response
}

func (b *bad) Say() string {
	return b.to + "!!!"
}

func (b *bad) Respond(response string) string {
	return response
}

func main() {
	f := &formal{}
	c := &casual{}
	w := &who{to: "World."}
	b := &bad{to: "schmucks"}
	p, e := plugin.Open("../plugin/p.so")
	if e != nil {
		panic(e)
	}
	out, e := p.Lookup("PrettyPrint")
	if e != nil {
		panic(e)
	}
	out.(func(greet.Greeter, greet.Sayer, string))(f, w, " - Hello. How have you been?")
	out.(func(greet.Greeter, greet.Sayer, string))(f, b, " - Now need to be rude.")
	out.(func(greet.Greeter, greet.Sayer, string))(c, w, " - Hi there.")
	out.(func(greet.Greeter, greet.Sayer, string))(c, b, " - Bug off.")
}
