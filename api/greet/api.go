package greet

type Sayer interface {
	Say() string
	Respond(response string) string
}

type Greeter interface {
	Greet(Sayer) string
}

func Greet(g Greeter, s Sayer, r string) string {
	return g.Greet(s) + s.Respond(r)
}
