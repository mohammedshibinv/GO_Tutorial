/*
Basic types:
string
bool
integer - uint8,uint64,int8,int64,uintptr
floating - float32,float64
complex

Other types:
array
slices
maps
struct
pointers
*/

package main

import "fmt"

const LoginToken string = "siodh4w9yfwe"//var is public=>starting letter is capital

func main(){
	var username string = "shibin"
	fmt.Printf("Type of %s in %T \n",username,username)

	var isLoggedin bool= false
	fmt.Printf("Type of %t in %T \n",isLoggedin,isLoggedin)

	var smallval uint8= 255
	fmt.Printf("Type of %d in %T \n",smallval,smallval)

	var smallfloat float32 = 255.5380934097943773947
	fmt.Printf("Type of %f in %T \n",smallfloat,smallfloat)

	//dafault value
	//The default value is zero, doesnt put any garbage value
	var val uint8
	fmt.Printf("Type of %d in %T \n",val,val)

	//implicit type
	var websit = "abc.com"
	fmt.Printf("Type of %s in %T \n",websit,websit)

	//no var style , only allowed inside func,
	count := 100
	fmt.Printf("Type of %d in %T \n",count,count)
}
