package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Name struct {
	Fname string
	Lname string
}

func main() {
	var filename string
	var names []Name

	fmt.Print("Enter the name of the text file: ")
	fmt.Scanln(&filename)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read each line from the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line into first name and last name
		fields := strings.Fields(line)
		if len(fields) != 2 {
			fmt.Println("Imvalid line: ", line)
			continue
		}

		name := Name{
			Fname: fields[0],
			Lname: fields[1],
		}

		names = append(names, name)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, name := range names {
		fmt.Printf("First Name: %s, Last Name: %s\n", name.Fname, name.Lname)
	}

}
