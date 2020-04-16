package greet

type Sayer interface {
	Say() string
}

type Greeter interface {
	Greet(Sayer) string
}

func Greet(g Greeter, s Sayer) string {
	return g.Greet(s)
}
