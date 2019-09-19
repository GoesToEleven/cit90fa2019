package main

import "fmt"

type person struct {
	first string
}

func (p person) foo() {
	fmt.Printf("person foo here %s\n", p)
}

type talker interface {
	foo()
}

func bar(t talker) {
	fmt.Println(t)
}

func (p person) fooToo() {
	fmt.Printf("person foo here %s\n", p)
}

type talkerToo interface {
	fooToo()
}

func barToo(t talkerToo) {
	fmt.Println(t)
}

func main() {
	p1 := person{
		first: "James",
	}
	bar(p1)
	
	p1.fooToo()
	
	barToo(p1)
}
