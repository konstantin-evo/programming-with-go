package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name string
	var address string

	fmt.Print("Enter a name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter an address: ")
	fmt.Scanln(&address)

	data := map[string]string{
		"name": name,
		"address": address,
	}

	jsonData, err := json.Marshal(data)
	if (err != nil) {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(jsonData))
}
