package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files in Golang")
	content := "This should go into the file - LCO"
	file, err := os.Create("./mygolangfile.txt")
	checkNilError(err)
	length, err := io.WriteString(file, content)
	checkNilError(err)
	fmt.Printf("The length of the file is: %v\n", length)
	defer file.Close()
	readFile("./mygolangfile.txt")
}

func readFile(fileName string) {
	data, err := ioutil.ReadFile(fileName) // The data comes in bytes format
	checkNilError(err)
	fmt.Println("The actual data from the file is : \n", data)
	fmt.Println("The string data from the file is : \n", string(data))
}

func checkNilError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
