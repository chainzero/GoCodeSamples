package main

import "fmt"

type Person struct {
	First string
	Last  string
}

type SecretAgent struct {
	Person
	ltk bool
}

type Human interface {
	speak()
}

func (p Person) speak() {
	fmt.Println("I am", p.First, p.Last, " - the person speak")
}

func (s SecretAgent) speak() {
	fmt.Println("I am", s.First, s.Last, " - the secretAgent speak")
}

func Speaker(h Human) {
	switch h.(type) {
	case Person:
		fmt.Println("I was called by person and my name is: ", h.(Person).First, h.(Person).Last)
	case SecretAgent:
		fmt.Println("I was called by SecretAgent and my name is: ", h.(SecretAgent).First, h.(SecretAgent).Last)
	}
}

func main() {
	sa := SecretAgent{
		Person: Person{
			First: "jessie",
			Last:  "carr",
		},
		ltk: false,
	}

	p := Person{
		First: "scott",
		Last:  "carr",
	}

	fmt.Println("Examples of calling the polymorphic function:")
	Speaker(sa)
	Speaker(p)
	fmt.Println("Examples of calling the interface's methods:")
	sa.speak()
	p.speak()

}
