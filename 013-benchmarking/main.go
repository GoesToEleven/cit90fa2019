package main

func main() {
	compLit()
	makeWay()
}

var cowboy int

func compLit() {
	ii := []int{}
	// l := len(ii)
	c := cap(ii)
	// fmt.Println(l)
	// fmt.Println(c)
	for i := 0; i < 1000000; i++ {
		tempC := c
		ii = append(ii, i)
		if tempC != cap(ii) {
			c = cap(ii)
			// fmt.Println("new cap:", c)
		}
	}
}

func makeWay() {
	ii := make([]int, 0, 1136640)
	// l := len(ii)
	c := cap(ii)
	// fmt.Println(l)
	// fmt.Println("MAKEWAY cap:", c)
	for i := 0; i < 1000000; i++ {
		tempC := c
		ii = append(ii, i)
		if tempC != cap(ii) {
			c = cap(ii)
			// fmt.Println("MAKEWAY new cap:", c)
		}
	}
}
