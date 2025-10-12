package logic

import "fmt"

func Findnumber(h map[string]string, names string) {
	fmt.Print("Please enter the name: ")
	fmt.Scanln(&names)
	switch names {
	case "Aziz", "aziz":
		fmt.Printf("\nThe phone-number is %v \n", h["Aziz"])
	case "Jasur", "jasur":
		fmt.Printf("\nThe phone-number is %v \n", h["Jasur"])
	case "Mukhammad", "mukhammad":
		fmt.Printf("\nThe phone-number is %v \n", h["Mukhammad"])
	case "Rinat", "rinat":
		fmt.Printf("\nThe phone-number is %v \n", h["Rinat"])
	default:
		fmt.Println("\nName doesn't exist, try other name")
	}

}
