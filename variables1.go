package main

import "fmt"

var c,python,java bool 

func main(){
	var i int 
	fmt.Println(i,c,python,java)
}

/* 
	var is the keyword used to create variables

	In Go , the type comes after the variable name

	variables can be declared inside or outside the functions

	Package-level variables(Global variables)
		var c,python ,java bool ->This line is outside main
			That means: 
				these variables belong to entire package 
				they are accessible anywhere inside this package 

	If we don't assign any value to variables go gives them default values
		For bool -> false
		For int ->0
		For string -> ""
		For float ->0.0
			Go never leaves variable uninitialized

	Function-level variable(Local variable)
		var i int -> declared inside main()
			exits only inside this function
				default value of int -> 0

*/
