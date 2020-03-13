package main

import "fmt"

func main() {
	var nums []int
	fmt.Println(nums, len(nums), cap(nums))
	nums = make([]int, 5)
	fmt.Println(nums, len(nums), cap(nums))
	capitais := []string{}
	fmt.Println(capitais, len(capitais), cap(capitais))
	capitais = append(capitais, "Brasilia")
	fmt.Println(capitais, len(capitais), cap(capitais))
	capitais = append(capitais, "New")
	fmt.Println(capitais, len(capitais), cap(capitais))
	capitais = append(capitais, "york")
	fmt.Println(capitais, len(capitais), cap(capitais))
	capitais = append(capitais, "china")
	fmt.Println(capitais, len(capitais), cap(capitais))
	capitais[0] = "sao paulo"
	fmt.Println(capitais, len(capitais), cap(capitais))
	for indice, cidade := range capitais {
		fmt.Printf("Cidade[%d] = %s\n", indice, cidade)
	}

}
