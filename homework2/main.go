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

	var intA int8 = 127
	var intB int16 = 32767
	var intC int32 = 2147483647
	var intD int64 = 9223372036854775807

	intA++
	intB++
	intC++
	intD++

	var uintA uint8 = 255
	var uintB uint16 = 65535
	var uintC uint32 = 4294967295
	var uintD uint64 = 18446744073709551615

	uintA++
	uintB++
	uintC++
	uintD++

	firstString := "🌎"
	secondString := "A🌎"
	thirdString := "🌎🌎🌎"
	fmt.Println([]byte(firstString))
	fmt.Println([]byte(secondString))
	fmt.Println([]byte(thirdString))

	var firstFloat float32 = 4.0
	var secondFloat float64 = 5.0
	var firstTarget float32 = 1e-8
	secondTarget := 1e-17

	fmt.Println((firstFloat + firstTarget) == 4)
	fmt.Println((secondFloat + secondTarget) == 5)

	i := 1
	i = i << 1
	fmt.Printf("%b\n", i)
	fmt.Println(i)

	i = i << 1
	fmt.Printf("%b\n", i)
	fmt.Println(i)

	i1 := 14
	i1 = i1 >> 1
	fmt.Printf("%b\n", i1)
	fmt.Println(i1)
}
