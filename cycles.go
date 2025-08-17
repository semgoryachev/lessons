package main

import f "fmt"

func main() {
	//simple cycle
	for i := 1; i <= 5; i++ {
		f.Println(i)
	}

	//cycle with range
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		f.Printf("Index: %d, Value: %d\n", index, value)
	}

}
