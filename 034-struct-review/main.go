package main

import "fmt"

type person struct {
	first string
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) String() string {
	return fmt.Sprintf("I am a person named %s", p.first)
}

func (s secretAgent) String() string {
	return fmt.Sprintf("I am a secret agent named %s and it is %t that I have a license to kill", s.first, s.ltk)
}

func main() {
	jm := person{
		first: "jenny",
	}
	fmt.Println(jm)
	a := jm.String()
	fmt.Println(a)

	jb := secretAgent{
		person: person{
			first: "James",
		},
		ltk: true,
	}
	fmt.Println(jb)
	b := jb.person.String()
	c := jb.String()
	fmt.Println(b)
	fmt.Println(c)
}
