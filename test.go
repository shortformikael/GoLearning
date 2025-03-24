package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := randNum()
	y := randNum()

	numCheck()
	printConcat()
	forLoop()
	fmt.Println("x + y = ", add(x, y))
}

func add(x, y int) int {
	return x + y
}

func forLoop() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func printConcat() {
	var name string = "John"
	age := 25
	const pi = 3.14

	fmt.Printf("Name: %s, Age %d, Pi %.2f\n", name, age, pi)
}

func numCheck() {
	num := 10
	numCheck := 5
	if num > numCheck {
		fmt.Printf("Greater than %d\n", numCheck)
	} else {
		fmt.Printf("Lesser than %d\n", numCheck)
	}

}

func printArr(pArr interface{}) {
	switch a := pArr.(type) {
	case [2]int:
		for i := 0; i < 2; i++ {
			fmt.Printf("%d: %d \n", i, a[i])
		}
	default:
		fmt.Println("Uknown Data type at printArr")
	}
}

func randNum() int {
	randomInt := rand.Intn(1000)
	return randomInt
}
