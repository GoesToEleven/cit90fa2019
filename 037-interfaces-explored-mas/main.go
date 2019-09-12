package main

import "fmt"

type person struct {
	first, last string
}

type dbPG map[int]person

type dbMongo map[int]person

type storer interface {
	access(int) person
	set(int, person)
}

func getter(s storer, x int) person {
	return s.access(x)
}

func setter(s storer, x int, p person) {
	s.set(x, p)
}

func main() {
	p1 := person{
		first: "jenny",
		last:  "moneypenny",
	}

	p2 := person{
		first: "james",
		last:  "bond",
	}

	// USING THE SAME METHODS
	// WORKING WITH POSTGRES
	mpg := dbPG{}
	// mpg.set(1, p1)
	setter(mpg, 1, p1)
	setter(mpg, 2, p2)
	fmt.Println(mpg)
	// fmt.Println(mpg.access(1))
	fmt.Println(getter(mpg, 1))
	fmt.Println(getter(mpg, 2))

	// USING THE SAME METHODS
	// WORKING WITH MONGO
	mm := dbMongo{}
	// mm.set(1, p2)
	setter(mm, 1, p1)
	setter(mm, 2, p2)
	fmt.Println(mm)
	// fmt.Println(mm.access(1))
	fmt.Println(getter(mm, 1))
	fmt.Println(getter(mm, 2))

}

func (pg dbPG) access(x int) person {
	return pg[x]
}

func (pg dbPG) set(x int, p person) {
	pg[x] = p
}

func (m dbMongo) access(x int) person {
	fmt.Println("accessing MONGO - note: this method is a different access method")
	return m[x]
}

func (m dbMongo) set(x int, p person) {
	fmt.Println("setting MONGO - note: this method is a different access method")
	m[x] = p
}
