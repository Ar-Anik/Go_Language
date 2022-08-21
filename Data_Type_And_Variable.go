/*
int8	8-bit signed integer
int16	16-bit signed integer
int32	32-bit signed integer
int64	64-bit signed integer
uint8	8-bit unsigned integer
uint16	16-bit unsigned integer
uint32	32-bit unsigned integer
uint64	64-bit unsigned integer
int		Both int and uint contain same size, either 32 or 64 bit.
uint	Both int and uint contain same size, either 32 or 64 bit.
rune	It is a synonym of int32 and also represent Unicode code points.
byte	It is a synonym of uint8.
uintptr	It is an unsigned integer type. Its width is not defined, but its can hold all the bits of a pointer value.
*/

/*
  "Bangladesh Small Country" -> string type
  21 -> int type
  21.5 -> floating type
  true/false -> bool
*/

// variation of float-> float32, float64

/* Variable Naming Rules
1. Letters, Digit, _
2. Keyword are not allowed as Variable
3. Variable can not have space
4. Variable name can not start with digit
5. Variable case sensitive
*/

package main

import "fmt"

func main() {
	var name, country, university string
	var age int
	var cgpa float32

	//Example - 1
	name = "Aubdur Rob Anik"
	age = 14
	cgpa = 3.57
	country = "Bangladesh"
	university = "University of Asia Pacific"

	fmt.Println("Name :", name)
	fmt.Println("Age :", age)
	fmt.Println("Cgpa :", cgpa)
	fmt.Println("Country :", country)
	fmt.Println("University :", university)

	//Example - 2
	var name1 string = "Anika Ani"
	var age1 int = 18
	var cgpa1 float32 = 3.87
	var country1 string = "Bangladesh"
	var university1 string = "Brac University"

	fmt.Println("My Name is", name1)
	fmt.Println(name1, "Age is", age1)
	fmt.Println(name, "cgpa is", cgpa1)
	fmt.Println(name1, "Country is", country1)
	fmt.Println(name1, "University is", university1)

	//Type Inference
	var name2 = "Jonny Dept"
	var age2 = 32
	var cgpa2 = 3.19
	var country2 = "America"
	var university2 = "Brazzers University"

	fmt.Println(name2, "is My Name.")
	fmt.Println(name2, "age is", age2)
	fmt.Println(name2, "cgpa is", cgpa2)
	fmt.Println(name2, "country is", country2)
	fmt.Println(name2, "university is", university2)

	//Example-4
	name3 := "Anisul Islam Oni"
	age3 := 21
	cgpa3 := 3.52
	country3 := "Noakhali"
	university3 := "Noakhali University"

	fmt.Println(name3, "is My Name")
	fmt.Println(name3, "Age is", age3)
	fmt.Println(name3, "cgpa is", cgpa3)
	fmt.Println(name3, "country is", country3)
	fmt.Println(name3, "university is", university3)
}
