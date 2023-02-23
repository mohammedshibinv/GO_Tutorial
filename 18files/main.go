package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	//create file
	file, err := os.Create("./file.txt")
	checkNilErr(err)

	defer file.Close()

	//wright to file,Use WrightString method from io package,retuns length if operation successfull
	content := "Hello its me shibin"
	len, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("Lenth of writtern content is", len)

	//read file, data is readed in byte format, need to convert into string
	dataBytes, err := ioutil.ReadFile("./file.txt")
	checkNilErr(err)
	fmt.Println("The content of file is \n", string(dataBytes))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
