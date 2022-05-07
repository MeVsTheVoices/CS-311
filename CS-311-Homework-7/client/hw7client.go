package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

//convenience function, get strings from user
func getMapOfStrings(strs ...string) map[string]string {
	var aMap = make(map[string]string)
	for _, j := range strs {
		var aValue string
		fmt.Print("Enter value for " + j + ": ")
		fmt.Scan(&aValue)
		aMap[j] = aValue
	}
	return aMap
}

func doAddEntry() {
	//get name and email
	aMap := getMapOfStrings("name", "email")
	aURL, _ := url.Parse("http://localhost:8080/add")
	values := aURL.Query()
	//enter values to query
	values.Add("name", aMap["name"])
	values.Add("email", aMap["email"])
	aURL.RawQuery = values.Encode()
	resp, err := http.PostForm(aURL.String(), values)
	//attempt to add
	if err != nil {
		fmt.Println("wasn't able to add " + aMap["name"] + " to the contact list")
		fmt.Println(err.Error())
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	}
}

func doLookupEntry() {
	//get name
	aMap := getMapOfStrings("name")
	aURL, _ := url.Parse("http://localhost:8080/lookup")
	values := aURL.Query()
	values.Add("name", aMap["name"])
	aURL.RawQuery = values.Encode()
	resp, err := http.Get(aURL.String())
	//report body, should contain email
	if err != nil {
		fmt.Println(aMap["name"] + " is not in the contact list")
		fmt.Println(err.Error())
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	}
}

func main() {
	fmt.Println("Managing email contacts")

	//go through until we get a 'yes' to quit
	for quitEntry := "no"; !strings.EqualFold(quitEntry, "yes"); {
		fmt.Print("Add of lookup (a/l): ")
		var addOrLookupChar string
		fmt.Scan(&addOrLookupChar)
		switch addOrLookupChar {
		case "a":
			doAddEntry()
		case "l":
			doLookupEntry()
		default:
			fmt.Println(addOrLookupChar + " is not a valid choice")
			continue
		}
		//prompt to continue
		fmt.Print("Do you want to quit? (yes/no): ")
		fmt.Scan(&quitEntry)
	}
}
