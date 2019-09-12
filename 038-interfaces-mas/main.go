package main

import "fmt"

type person struct {
	first string
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) speak() {
	fmt.Printf("I am a person named %s\n", p.first)
}

func (s secretAgent) speak() {
	fmt.Printf("I am a secret agent named %s and it is %t that I have a license to kill\n", s.first, s.ltk)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	jm := person{
		first: "jenny",
	}
	saySomething(jm)

	jb := secretAgent{
		person: person{
			first: "James",
		},
		ltk: true,
	}
	saySomething(jb)
}
