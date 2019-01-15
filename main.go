package main

import (
	"fmt"
	"time"
)

func main() {

	// Values
	fmt.Println("go" + "lang")
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	fmt.Println(!true)

	// Variables
	var a = "intial"
	fmt.Println("a =", a)

	b := "short"
	fmt.Println("b is", b)

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	// If Else
	if num := 7; num == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// Switch Case
	switch time.Now().Weekday() {
	case time.Sunday, time.Saturday:
		fmt.Println("It's weekend")
	default:
		fmt.Println("It's weekday")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)

	// Arrays
	var names [2]string
	names[0] = "jide"
	names[1] = "taiwo"
	fmt.Println(names)
	fmt.Println("My name is", names[0])

	salaries := [5]int{5000, 4000, 50000, 100000, 10000000}
	fmt.Println("Salaries", salaries)

	salariesCount := len(salaries)
	for i := 0; i < salariesCount; i++ {
		fmt.Println(salaries[i])
	}

	// Slices
	s := make([]string, 5)
	s[0] = "a"
	s[1] = "b"
	s[2] = "x"
	s[3] = "y"
	s[4] = "z"
	s = append(s, "c")
	s = append(s, "d")
	s = append(s, "e")
	fmt.Println(s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)
	fmt.Println(s[2:])
	fmt.Println(s[:5])
	fmt.Println(s[2:5])

	// Maps
	
}
