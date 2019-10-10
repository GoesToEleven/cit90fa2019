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
	fmt.Println("hello from person", p.first)
}

func (sa secretAgent) speak() {
	fmt.Println("hello from secretAgent", sa.first)
}

type human interface {
	speak()
}

func main() {
	m := map[int]string{1: "jenny", 2: "james"}
	fmt.Println(m)

	v, ok := m[1]
	fmt.Println(v, ok)

	v, ok = m[3]
	fmt.Println(v, ok)

	if v, ok := m[1]; ok {
		fmt.Println("printing again", v)
	}

	type hotdog int
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	// conversion
	y := int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	var sa human
	fmt.Printf("TYPE OF SA %T\n", sa)
	sa = secretAgent{
		person: person{
			first: "Jenny",
		},
		ltk: true,
	}

	fmt.Println(sa)
	fmt.Printf("%T\n", sa)

	z := sa.(person)
	fmt.Println(z)
	fmt.Printf("%T\n", z)

	c := make(chan string)
	g
}

/*

v, ok = m[key]
v, ok = x.(T)
v, ok = <-ch

*/
