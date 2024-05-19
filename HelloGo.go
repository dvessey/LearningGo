package main

import (
	"fmt"
	"log"
	"reflect" // required to use tags
	"runtime"
	"strconv"
)

// capital first letter means it can be exported
type Vehicle struct {
	Make  string `required max:"100"` //`tag`
	Cost  int
	Model []string
}

// Package Scope
var k = 12

// upper case first character in package scope globally accessible (exported)
var I int = 30

func main() {
	// All uninitalized values in go are a 0 value or false for booleans

	// Signed Integers
	// int8 -128 to 127
	// int16 -32,768 to 32,767
	// int 32 -2,147,483,648 to 2,147,483,647
	// int 64 -9,223,372,036,854,775,808 to int 64 -9,223,372,036,854,775,807

	// Unsigned Integers
	// uint8 0 to 255
	// uint16 0 to 65,535
	// uint32 0 to 4,294,967,295

	var j int = 42
	i := 100

	//best practice uppercase acronyms
	var theURL string = "https://google.com"

	// string are aliaes for bytes utf8 character
	s := "this is a string"
	fmt.Printf("%v, %T\n", s, s)

	//convert string to byte slice
	b := []byte(s)
	fmt.Printf("%v, %T\n", b, b)

	// rune utf32 character type alias for int32
	r := 'a'
	fmt.Printf("%v, %T\n", r, r)

	//convert an int to a string
	var m string = strconv.Itoa(i)

	//constants best practice - if you wrote MyConst, the first letter capitalized means it will be exported
	const myConst int = 90

	fmt.Printf("%v, %T\n", i, i) //print value and type
	fmt.Println(j)
	fmt.Println(myConst)
	fmt.Println(k)
	fmt.Printf("%v, %T\n", m, m) //print value and type
	fmt.Println(theURL)

	// ARRAYS
	grades := [...]int{97, 85, 93, 96}
	var students [4]string
	students[0] = "Damon"
	students[1] = "Matthew"
	students[2] = "Jim"
	students[3] = "Bob"
	fmt.Printf("Grades: %v\n", grades)
	fmt.Printf("Students: %v\n", students)
	//derefernce and get the value from the array
	fmt.Printf("Students: %v\n", students[0])
	fmt.Printf("Number of students: %v\n", len(students))

	//arrays in Go are actually considered values, so arrays are copied when assigning an array to another array
	a := [...]int{1, 2, 3}
	c := a
	c[1] = 5
	fmt.Println(a)
	fmt.Println(c)

	//to point to the array with the same data
	d := &a
	fmt.Println(d)

	//SLICES
	e := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(e)
	fmt.Printf("Length of slice: %v\n", len(e))
	fmt.Printf("Capacity of slice: %v\n", cap(e))

	f := e[:]   //slice of all elements
	g := e[3:]  //slice from 4th element to end
	h := e[:6]  // slice of first 6 elements
	n := e[3:6] //slice of 4th, 5th and 6th elements

	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(n)

	o := make([]int, 3, 100) //make an int array with length of 3 with capacity of 100
	o = append(o, 1)         // resizing arrays copies all the data to a new array
	fmt.Println(o)
	fmt.Printf("Length of slice: %v\n", len(o))
	fmt.Printf("Capacity of slice: %v\n", cap(o))

	// MAPS
	//statePopulations := make(map[string]int)
	statePopulations := map[string]int{
		"California": 1000,
		"Texas":      2000,
		"Florida":    3000,
	}
	//return order is not guaranteed
	fmt.Println(statePopulations)
	//slices are not valid keys to a map, but arrays are
	delete(statePopulations, "California")
	fmt.Println(statePopulations)
	fmt.Println(statePopulations["California"]) // Returns 0 because the key no longer exists, can cause confusion
	// is the population 0 or does it not exist? ok syntax to check:
	pop, ok := statePopulations["Colorado"]
	fmt.Println(pop, ok) // ok returns false if key was not found in map

	//STRUCTS
	// Go does not have inheritance ("is a" relationship), but does support composition ("has a" relationship)
	//gets copied when assigning to a new struct
	aVehicle := Vehicle{
		Make: "Lambo",
		Cost: 100000000,
		Model: []string{
			"Aventador",
			"Diablo",
		},
	}
	fmt.Println(aVehicle)

	// to get the tag to be used in a validation library or something you create
	v := reflect.TypeOf(Vehicle{})
	field, _ := v.FieldByName("Make")
	fmt.Println(field.Tag)

	panicker()
	fmt.Println("back in main from panicker")

	//POINTERS
	//Go does not allow pointer arthimatic (unless using the unsafe package)
	// 0 intialization value for a pointer in Go is nil
	x := 35
	y := x // y is a new copy
	fmt.Println(x, y)
	var z *int = &x //z is pointing to the data of x, this is not a copy
	fmt.Println(z)  //holding the memory address of x
	fmt.Println(*z) //dereference to get the value of x

	greeting := "Hello"
	name := "Stacey"
	sayGreeting(&greeting, &name)
	fmt.Println(name) //prints ted

	//GO ROUTINES
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1)) // max number of operating system threads, can set higher
	// go run -race // will tell you about data race conditions at compile time

}

func sayGreeting(greeting, name *string) {
	fmt.Println(*greeting, *name) //prints stacey
	*name = "Ted"
	fmt.Println(*name) //prints ted
}

// defer occurs LIFO order and occurs once Go recognizes it's at the end of the function, before it returns
// panics happen AFTER defer statements
func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error: ", err)
		}
	}()
	panic("Something bad happened")
	fmt.Println("done panicking") // this will not run, but the defer function will return control to main program

}
