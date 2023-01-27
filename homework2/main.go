package main

import "fmt"

func main() {
	valA := false
	valB := true
	valC := false

	firstResult := (valA && valB) == (valB && valA)
	secondResult := ((valA && valB) && valC) == (valA && (valB && valC))
	thirdResult := (valA && (valB || valC)) == ((valA && valB) || (valA && valC))
	forthResult := !!valA == valA
	fifthResult := (valA && valA) == valA

	fmt.Println(firstResult)
	fmt.Println(secondResult)
	fmt.Println(thirdResult)
	fmt.Println(forthResult)
	fmt.Println(fifthResult)
}
