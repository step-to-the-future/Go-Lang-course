package main

import (
	"github.com/step-to-the-future/Go-Lang-course/internal/logic"
)

func main() {
	var names string
	phonebook := map[string]string{
		"Aziz":      "+(998) 95 007 - 05 - 14",
		"Mukhammad": "+(998) 97 005 - 45 - 94",
		"Jasur":     "+(998) 93 507 - 08 - 88",
		"Rinat":     "+(998) 93 105 - 67 - 63",
	}

	logic.Findnumber(phonebook, names)

}
