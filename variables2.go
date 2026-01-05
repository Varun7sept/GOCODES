package main

import "fmt"

var i,j int =1,2 //package level variable

func main(){
	var c, python,java=true,false,"no!"
	fmt.println(i,j,c,python,java)
}

/* Variables with initializers
	What if I want to give values to variables at the time I create them?
		When you declare a variable using var,you are allowed to give it a starting value immediately
		Each variable can have its own value

	If an initializer is present, the type can be omitted; the variable will take the tye of the initializer

	If you giva a value,Gocan figure out the type.
	If you don't give a value ,you must specify the type
*/	