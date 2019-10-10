package main

import "fmt"

type person struct {
	first string
}

// concrete types
// a concrete type specifies the exact representation of its values
// and exposes the intrinsic operations of that representation
// such as arithmetic for numbers, or indexing / append / range for slices
// a concrete type may also provide additional behaviors through methods
// when you have a value of a concrete type, you know exactly
// what it is, and what it does
type mdb map[int]person
type pgdb map[int]person
type mysql map[int]person

func (my mysql) save(n int, p person) {
	fmt.Println("mysql save")
	my[n] = p
}

func (my mysql) retrieve(n int) person {
	fmt.Println("mysql retrieve")
	return my[n]
}

func (m mdb) save(n int, p person) {
	fmt.Println("mongo save")
	m[n] = p
}

func (m mdb) retrieve(n int) person {
	fmt.Println("mongo retrieve")
	return m[n]
}

func (pg pgdb) save(n int, p person) {
	fmt.Println("pg save")
	pg[n] = p
}

func (pg pgdb) retrieve(n int) person {
	fmt.Println("pg retrieve")
	return pg[n]
}

// abstract type
type dbAccessor interface {
	save(n int, p person)
	retrieve(n int) person
}

func set(db dbAccessor, n int, p person) {
	db.save(n, p)
}

func pull(db dbAccessor, n int) person {
	return db.retrieve(n)
}

func main() {
	mongo := mdb{}
	postgres := pgdb{}
	mysql := mysql{}

	p1 := person{
		first: "james",
	}

	set(mongo, 1, p1)
	set(postgres, 1, p1)
	set(mysql, 1, p1)
	fmt.Println(pull(mongo, 1))
	fmt.Println(pull(postgres, 1))
	fmt.Println(pull(mysql, 1))
}
