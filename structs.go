/*
	A sruct is just a way to group related data together

	Definining a struct
		type person struct{
			name string
			age int 
		}
			type person struct{...}
				you are creating a new data type called person
			Inside it
				name -> string
				age -> int
			
	Creating a struct value (basic way)
		fmt.Println(person{"Bob",20})
			person{"Bob",20} creates a new person
			Order matters :
				"Bob"->name
				 20->age
			Output ->{Bob 20}

	Creating a struct using field names(recommended)
		fmt.Println(person{name:"Alice",age:30})
			Order does not matter
			Code is more readable
			Safer when structs grow bigger
		Output -> {Alice 30}

	Zero values in structs(very important)
		fmt.Println(person{name:"Fred"})
			you didn't give age ,so Go uses zero value(default value)
		Output-> {Fred 0}
		GO never leaves garbage values. Every field gets a default.

	Getting a pointer to a struct(&)
		fmt.Println(&person{name: "Ann",age: 40})
			& means "give me the address"
			This returns a pointer to the struct
		Ouptut-> &{Ann 40}
		struct can be used with pointers safely in Go

	
	