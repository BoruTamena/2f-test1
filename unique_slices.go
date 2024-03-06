package main

import "fmt"

func remove_dup(input_slice []int) []int {

	check := make(map[int]bool)
	unique_slice := []int{}

	for _, ent := range input_slice {

		if _, ok := check[ent]; !ok {

			check[ent] = true
			unique_slice = append(unique_slice, ent)

		}
	}

	return unique_slice

}

func main() {

	// creating slice with duplicate value
	dup_slice := []int{1, 1, 2, 2, 4, 4, 5, 6}
	fmt.Printf("slice with duplicate value\n %v\n", dup_slice)

	// removing duplicate values from te slice
	uniq_slice := remove_dup(dup_slice)
	fmt.Printf(" After duplicate value removed \n %v\n", uniq_slice)

}
