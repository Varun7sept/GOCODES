package main

import(
	"fmt"
	"math"
)
func main(){
	fmt.println(math.pi)//unexported in pi ,p is small
	fmt.pritnln(math.Pi)//exported Pi starts with capital P
}

/*
	I Go ,whether you can use something from another package depends on whether
	it name starts with a capital letter or not.const

	If a name starts with a captial letter -> it is exported(visible to other packages)
	If a name starts with a lowercase letter -> it is not exported

	When importing a package ,you can refer only to its exported names

	This rule applies to everything 
		variables , functions , structs ,methods ,constants
