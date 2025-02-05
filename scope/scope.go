package main

import "fmt"

var a = 20
var b = 30

func add(x int, y int) {
	z := x + y
	fmt.Println(z)
}

/*func main() {
	p := 30
	q := 40

	if p < q {
		l := 2
		fmt.Printf("p is greater than q so, %v\n", l)
	}

	add(p, q)
	add(a, b)
	add(a, p) // you can not use z in this main block because not in scope
	//add(z+l)// you cannot use z because outside of main scope and you ca
	//not use l because it is in a nested local scope
}*/
