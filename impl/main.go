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
	return fmt.Sprintf("Hey %s!", s.Say())
}

func (w *who) Say() string {
	return w.to
}

func (b *bad) Say() string {
	return b.to + "!!!"
}

func main() {
	f := &formal{}
	c := &casual{}
	w := &who{to: "World"}
	b := &bad{to: "schmucks"}
	p, e := plugin.Open("../plugin/p.so")
	if e != nil {
		panic(e)
	}
	out, e := p.Lookup("PrettyPrint")
	if e != nil {
		panic(e)
	}
	out.(func(greet.Greeter, greet.Sayer))(f, w)
	out.(func(greet.Greeter, greet.Sayer))(f, b)
	out.(func(greet.Greeter, greet.Sayer))(c, w)
	out.(func(greet.Greeter, greet.Sayer))(c, b)
}
