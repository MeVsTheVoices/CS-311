// CS-311-Homework-2 project main.go
//Author: Joshua Dunne
//Class: 20225_CS_311_01_1
package main

import (
	"fmt"
	"regexp"
)

func doTestRegexForUserIDs() {
	exampleUserIDs := [...]string{"ab7044", "Gz9911", "hO3062", "08hghg", ".uI53km"}
	for i := 0; i < len(exampleUserIDs); i++ {
		matched, _ := regexp.MatchString(`^[a-zA-Z]{2}[0-9]{4}$`, exampleUserIDs[i])
		if matched {
			fmt.Println(exampleUserIDs[i], " is a correctly formated userID")
		} else {
			fmt.Println(exampleUserIDs[i], " is not a correctly formated userID")
		}
	}
}

func doTestRegexForHorizonEmails() {
	exampleEmails := [...]string{
		"jdunne2@horizon.csueastbay.edu",
		"billybobthronton@gmail.com",
		"IMJUSTARANDOMSTRING",
		"thatGuy@horizon.csueastbay.edu",
		"someDudeWithsomeNumbers99@horizon.csueastbay.edu",
	}
	for i := 0; i < len(exampleEmails); i++ {
		matched, _ := regexp.MatchString(`^[^@]+@horizon.csueastbay.edu$`, exampleEmails[i])
		if matched {
			fmt.Println(exampleEmails[i], " is a correctly formated horizon email")
		} else {
			fmt.Println(exampleEmails[i], " is not a correctly formated horizon email")
		}
	}
}

func main() {
	doTestRegexForUserIDs()
	doTestRegexForHorizonEmails()
}
