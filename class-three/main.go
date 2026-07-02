package main

import "fmt"

func main() {
	// estrutura de dados array

	array := [5]string{"um", "dois", "tres", "quatro", "cinco"}

	// estrutura de dados slice

	slice := []int{1, 2, 3, 4, 5}

	// xx1 = 1, xx2 = 2, xx3 = 3, xx4 = 4, xx5 = 5 = limpar

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	slice = append(slice, 6, 7, 8, 9, 10)
	// xx55 = 1, xx56 = 2, xx57 = 3, xx58 = 4, xx59 = 5, xx60 = 6, xx61 = 7, xx62 = 8, xx63 = 9, xx64 = 10

	for i := 0; i < len(array); i++ {

		fmt.Println(&array[i])

		if array[i] == "tres" {
			fmt.Println("encontrei o tres")
		}

		fmt.Println(array[i])
	}
	fmt.Println("------------")

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// estrutura de dados map

	maps := make(map[string]int)
	// dictionary
	// dict
	// hash map

	maps["um"] = 1
	maps["dois"] = 2
	maps["tres"] = 5000
	maps["quatro"] = 4
	maps["cinco"] = 5

	fmt.Println()

	value, ok := maps["tres"]

	if ok {
		fmt.Println("encontrei o tres", &value)
	}

}
