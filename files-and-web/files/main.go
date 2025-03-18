package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := "./files-and-web/files/message.txt"
	file, err := os.Create(fileName)
	defer file.Close()
	checkError(err)
	length, err := io.WriteString(file, "Learn Go and change your life")
	fmt.Printf("Message has been sent with %v characters\n", length)
	readFile(fileName)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(fileName string) {
	data, err := os.ReadFile(fileName)
	checkError(err)
	fmt.Println("Read message:", string(data))
}