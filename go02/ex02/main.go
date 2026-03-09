package main

import (
	"fmt"
)

var phoneBook = make(map[string][]string) // name -> list of phone numbers

func main() {
	fmt.Println(phoneBook)
	Add("Johannes", "+4915126216542")
	fmt.Println(phoneBook)
}

func Add(name, phone string) {
	phoneBook[name] = append(phoneBook[name], phone)
}
