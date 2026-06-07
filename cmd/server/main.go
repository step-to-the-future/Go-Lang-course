package main

import (
	"fmt"

	"github.com/step-to-the-future/Go-Lang-course/internal/logic"
)

func main() {

	name := "something new!"
	fmt.Println(name)

	var names string

	phonebook := map[string]string{
		"Abdulaziz": "+(998) 95 007 - 05 - 15",
		"Mukhammad": "+(998) 97 005 - 45 - 18",
		"Jasur":     "+(998) 93 507 - 08 - 88",
		"Rinat":     "+(998) 93 105 - 67 - 63",
	}

	logic.Findnumber(phonebook, names)

}
