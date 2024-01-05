package mystruct

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func Run() {
	favFood := []string{"kimchi", "ramen"}
	nico := person{name: "nico", age: 18, favFood: favFood} //아님 키값 다 뺴고 주든가.
	fmt.Println(nico)
}