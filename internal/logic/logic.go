package logic

import "fmt"

func Findnumber(h map[string]string, names string) {
	fmt.Print("Please enter the name: ")
	fmt.Scanln(&names)

	for key, value := range h {
		if key == names {
			fmt.Printf("\nThe name is %s and the number is %s\n", key, value)
		}
	}

}
