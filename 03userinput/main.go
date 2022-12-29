package main

import "bufio"
import "fmt"
import "os"

func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter something")

	//comma ok / error ok syntax
	input, _:=reader.ReadString('\n')

	fmt.Println(input)
	fmt.Printf("type : %T",input)
}