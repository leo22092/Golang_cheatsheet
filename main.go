package main

import (
	"errors"
	"fmt"
)

func main() {
	//BASICS
	// fmt.Println("Hello world")
	// var name string = "Hello tim"
	// var number uint8 = 255
	// fmt.Println(number)

	// fmt.Println(name)
	// printMe(45)
	// var result, reminder, err = intdivision(10, 0)

	// switch {
	// case err != nil:
	// 	fmt.Printf(err.Error())
	// case reminder == 0:
	// 	fmt.Printf("The result is %v", result)
	// default:

	// 	// fmt.Print(result, reminder)
	// 	fmt.Printf("The result is %v with reminder %v ", result, reminder)
	// }

	array_learn()
}
func array_learn() {
	// ARRAys
	// 1. Fixed length -->Length cannot be changed after initializasion
	// 2. same type -->All elements of array are of same type like int 32,default value(like 0 for int32) are assigned
	// 3. Indexable -->slicing etc.
	// 4. Contaguos in memory -->int32 is 4 bytes of memmory ie 2^4
	// for a 3 member array contagious memmory of 4x3=12 bytes are alloted
	// (int 8 -->-128to127 ie 2^8 and uint8 --> 0 to 255)

	var intArr [3]int32
	fmt.Println(intArr[2]) //Prints 0
	intArr[2] = 544
	fmt.Println(intArr[2])
	// to print out memmory location
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])
	// Different way to initialise array
	var intArr2 [3]int32 = [3]int32{1, 2}
	fmt.Println(intArr2)

	// 3rd value will be zer
	//  or the initialisation can be done like,
	intArr3 := [3]int32{1, 2, 3}
	fmt.Println(intArr3)
	// Can also omit the [3] by letting the compiler assume the length usin ...
	intArr4 := [...]int32{1, 2, 3, 4, 5}
	fmt.Println(intArr4)
	// SLICE
	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Println(intSlice)
	fmt.Printf("The length is %v with capacity %v ", len(intSlice), cap(intSlice))
	// Appending here 3 more spaces are allocated that is intSlice is of size 3 and to append 7 its capacity is full
	// so another slice of size 3 is also allocated ie a new slixe variable is formed with capacity  of 6
	// 1.check if underlying array have the capacity to hold 4 values.
	// 2. if not a new array is made with enough capacity and values are copied there
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)
	fmt.Printf("The length is %v with capacity %v ", len(intSlice), cap(intSlice))
	fmt.Println(intSlice)
	// Reallocation of arrays can slow the programme so decalre the size in advance if you have rough idea
	// Eg. using make(size,capacity)  -->make an array to store int 32 values of size 3 and capacity 8 values at
	//  first 3 index will be filled with 0 and if tried to access the 4 th it will be index error
	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Printf("The length is %v with capacity %v ", len(intSlice3), cap(intSlice3))
	// MAPPING
	var myMap map[string]uint8 = make(map[string]uint8, 10)
	fmt.Printf("The length is %v ", len(myMap))
	fmt.Println(myMap)
	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2["Adam"])
	// Map returns the default value of data type even if key doesnt exist
	fmt.Println((myMap2["Jason"]))
	//But map also returns an optional boolean value if the key exists -->true else-->false
	var age, ok = myMap2["jason"]
	if ok {
		fmt.Printf("The age is %v \n", age)
	} else {
		fmt.Println("Invalid number")
	}
	fmt.Println(ok, age)
	//1 to iterate over
	// While iterating over a map no order is preserved.
	for name, age := range myMap2 {
		fmt.Printf("Name : %v, Value : %v \n", name, age)
	}
	// for loop can be formed in other syntax like,
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// For loop can iterate over array and slices or use of break inside for

}

func printMe(printval uint32) {
	fmt.Print("HElloooi..")
	fmt.Print(printval)
}

func intdivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot Divide by zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var reminder int = numerator / denominator
	// 	// Error is a data type default value is nill

	return result, reminder, err
}
