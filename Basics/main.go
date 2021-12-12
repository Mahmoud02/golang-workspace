package main

import "fmt"

// Define Const Values
// iota start from 0 value and increment every time
// iota vale reset every const block
const (
	const1 = iota
	const2 = iota
	const3 // will evaluate last const expression or value
)

var (
	varblock = "ff"
)

func main() {

	fmt.Println("Hello, world")

	//Variables
	var i int = 42
	fmt.Println(i)

	//omit type
	var x = 42
	fmt.Println(x)

	//initialize in two lines
	var a int
	a = 42
	fmt.Println(a)

	//omit 'var' and implicit variable Datatype
	firstName := "Mahmoud"
	fmt.Println(firstName)

	//complex Data type
	complexNumber := complex(1, 6)
	fmt.Println(complexNumber)

	//Multiple  Assignment
	realNumber, imaginaryNumber := real(complexNumber), imag(complexNumber)
	fmt.Println(realNumber, imaginaryNumber)

	/*
		Pointers : The Dark Evil
	*/
	var namePointer *string
	fmt.Println(namePointer) // output <Nil>:Empty pointer, POINT TO NOTHING

	var namePointerInitialize *string = new(string)
	fmt.Println(namePointerInitialize) // MemoryAddress of pointer

	*namePointerInitialize = "Belal"    // de reference operation
	fmt.Println(*namePointerInitialize) // value that pointer , point to

	// point to variable
	name := "Belal"
	ptr := &name
	fmt.Println(ptr, *ptr) // value that pointer , point to

	/*
		Constant
	*/

	const pi = 3.14
	/*
		Collections: built-in
	*/

	//Array
	var intArr [3]int
	intArr[0] = 1
	intArr[1] = 2
	intArr[2] = 3
	fmt.Println(intArr)
	fmt.Println(intArr[0])

	intArr1 := [3]int{4, 5, 6}
	fmt.Println(intArr1)

	//Slicing: the act like pointers ,  point to array value
	//it is not a fixed size
	intArrSlice := intArr[:]
	fmt.Println(intArrSlice)
	intArr[2] = 22
	fmt.Println(intArrSlice)
	intArrSlice[2] = 55
	fmt.Println(intArrSlice)
	fmt.Println(intArr)

	sliceArrayImplicit := []int{1, 2, 3, 4}
	fmt.Println(sliceArrayImplicit)
	sliceArrayImplicit = append(sliceArrayImplicit, 5)
	fmt.Println(sliceArrayImplicit)

	//Maps:
	//map[key]value
	mapCollection := map[string]int{"key1": 1}
	fmt.Println(mapCollection)
	fmt.Println(mapCollection["key1"])
	delete(mapCollection, "key1")
	fmt.Println(mapCollection)

	//Struct
	type User struct {
		ID        int
		FirstName string
	}
	var user User
	fmt.Println(user)
	user.ID, user.FirstName = 1, "Mahmoud"
	fmt.Println(user)
	user2 := User{ID: 2, FirstName: "Belal"}
	fmt.Println(user2)
	/*
		Loops:
	*/
	var counter int // initial value of int is 0
	for counter < 5 {
		print(counter)
		counter++
		if counter == 3 {
			break
		}
	}
	/*
		Panic Function:
	*/
	panic("something very bad happen")
	print("I will not be executed")

	/*
		Notes
		 -can't add int to float
	*/
}
