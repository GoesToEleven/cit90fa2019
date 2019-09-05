package main

import "fmt"

func main() {
	m := map[string]int{
		"james": 32,
		"jenny": 24,
	}
	fmt.Println(m)
	fmt.Println(m["james"])
	fmt.Println(m["jenny"])

	v, ok := m["james"]
	fmt.Println(v)
	fmt.Println(ok)

	v, ok = m["jack"]
	fmt.Println(v)
	fmt.Println(ok)

	xi := []int{42, 43, 44, 45}
	for i, v := range xi {
		fmt.Println("slice", i, v)
	}

	for k, v := range m {
		fmt.Println("map", k, v)
	}

	type flavor string
	xf := []flavor{"chocolate", "vanilla", "double chocolate"}
	for i, v := range xf {
		fmt.Printf("flavors %d %s - ", i, v)
		fmt.Printf("type %T - converted %T\n ", v, string(v))
	}

	type orders map[int][]flavor

	o := orders{
		1: []flavor{"chocolate"},
		2: []flavor{"vanilla"},
		3: []flavor{"strawberry"},
		4: []flavor{"peanut butter chocolate"},
		5: []flavor{"mint chocolate"},
	}

	for k, v := range o {
		fmt.Printf("order number %d had ", k)
		for _, vv := range v {
			fmt.Printf("%s ", vv)
		}
		fmt.Printf("\n")
	}

	// init; cond; post
	for i := 0; i < 10; i++ {
		fmt.Printf("map loop number %d - ", i)
		for _, v := range o {
			fmt.Printf("%s ", v)
		}
		fmt.Println()
	}

	xf2 := []flavor{"chocolate", "vanilla", "strawberry", "peanut butter chocolate", "mint chocolate"}
	for i := 0; i < 10; i++ {
		fmt.Printf("slice loop number %d - ", i)
		for _, v := range xf2 {
			fmt.Printf("%s ", v)
		}
		fmt.Println()
	}

	fmt.Println(xf2[4])
}
