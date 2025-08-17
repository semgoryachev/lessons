package main

import f "fmt"

func main() {
	var numbers [5]int
	f.Println(numbers)
	numbers[4] = 22
	numbers[0] = 3
	f.Println(numbers)

	fruits := [3]string{"Apple", "Banana", "Grapes"}
	f.Println(fruits[1])

	originalArray := [3]int{1, 2, 3}
	copiedArray := originalArray
	copiedArray[0] = 100
	f.Println(originalArray)
	f.Println(copiedArray)

	for i := 0; i < len(numbers); i++ {
		f.Printf("Element at index %d : %d\n", i, numbers[i])
	}

	for index, value := range fruits {
		f.Printf("Index - %d, value - %v\n", index, value)
	}

	//"_" - underscore is blank identifier to store unused values
	for _, v := range fruits {
		f.Printf("Fruit: %v\n", v)
	}

	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	f.Println(matrix)

}
