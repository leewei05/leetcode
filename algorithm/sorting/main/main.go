package main

import "fmt"

func main() {
	s := []int{4, 6, 3, 2, 5, 7, 9, 8, 1}
	bubble := bubbleSort(s)
	fmt.Println("bubble sort:", bubble)

	c := []int{4, 6, 3, 2, 5, 7, 9, 8, 1}
	selection := selectSort(c)
	fmt.Println("selection sort:", selection)
}

func bubbleSort(slice []int) []int {
	fmt.Println("Bubble Sort")
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice)-1; j++ {
			if slice[j] < slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
				fmt.Println(slice)
			}
		}
	}

	return slice
}

func selectSort(slice []int) []int {
	fmt.Println("Selection Sort")
	for i := 0; i < len(slice)-1; i++ {
		min := i
		for j := i + 1; j < len(slice); j++ {
			if slice[min] > slice[j] {
				min = j
			}
		}
		slice[i], slice[min] = slice[min], slice[i]
		fmt.Println(slice)
	}

	return slice
}
