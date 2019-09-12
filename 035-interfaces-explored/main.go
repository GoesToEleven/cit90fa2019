package main

import "fmt"

type person struct {
	first, last string
}

type dbPG map[int]person

type dbMongo map[int]person

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

func main() {
	p1 := person{
		first: "jenny",
		last:  "moneypenny",
	}

	mpg := dbPG{}
	mpg.set(1, p1)
	fmt.Println(mpg)
	fmt.Println(mpg.access(1))

	p2 := person{
		first: "james",
		last:  "bond",
	}

	mm := dbMongo{}
	mm.set(1, p2)
	fmt.Println(mm)
	fmt.Println(mm.access(1))
}
