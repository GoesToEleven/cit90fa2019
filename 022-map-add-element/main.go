package main

import "fmt"

func main() {
	//composite literal
	xs := []string{"james", "jenny"}
	fmt.Println(xs)

	//composite literal
	m := map[string]string{
		"james": "martini",
		"jenny": "manhattan",
	}
	fmt.Println(m)

	m["jack"] = "jack daniels"
	fmt.Println(m)
	
	delete(m, "jack")
	fmt.Println(m)
}
