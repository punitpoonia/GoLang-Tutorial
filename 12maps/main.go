package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]string)
	m["name"] = "Punit"
	m["surname"] = "Poonia"
	fmt.Println(m)
	fmt.Println(m["name"] + m["surname"])
	delete(m, "surname")
	fmt.Println(m)
	clear(m)
	fmt.Println(m)

	val, ok := m["name"]
	fmt.Println(val)

	if ok {
		fmt.Println("All Good")
	} else {
		fmt.Println("nOT oK")
	}

	// maps.Equals

	m1 := map[string]int{"age": 21, "phone": 12}
	m2 := map[string]int{"age": 21, "phone": 8}

	fmt.Println(maps.Equal(m1, m2))

}
