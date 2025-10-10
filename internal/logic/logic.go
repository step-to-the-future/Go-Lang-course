package logic

import "fmt"

func Cheest() {
	var n int
	fmt.Print("Please enter max numbers of elements in an array: ")
	fmt.Scanln(&n)
	var arr []int
	var r int
	fmt.Println("Please enter elements of an array: ")
	for i := 0; i < n; i++ {
		fmt.Printf("Please enter the %v number: ", i+1)
		fmt.Scanln(&r)
		arr = append(arr, r)
	}

	fmt.Println("\nThe array is:", arr)

	var minVal int = arr[0]
	for y := 1; y < n; y++ {
		if minVal > arr[y] {
			minVal = arr[y]
		}
	}
	fmt.Println("\nThe min val in array is", minVal)
}
