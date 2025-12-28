/*
	What is a map in Go?
		A map is like a dictionary or a phonebook
			You have a key
			You have a value
			Using the key , you find the value
		Example
			Name->Phone Number
			"Varun"->7217607148
		key->value

	How to create a map
		m:=make(map[string]int)
			map[string]int 
				key type ->string
				value type ->int
			make ->allocates memory(creates the map)
		You must use make before using a map

	Adding key-value pairs
		m["k1"]=7
		m["k2"]=13
			"k1"is the key
			 7 is the value =>(k1->7)

	Printing a map
		fmt.Println("map:",m)
			Output -> map:map[k1:7 k2:13]
		Go prints maps like 
			map[key:value key:value]
		Order is not guaranteed : Maps are unordered.

	Getting a value from a map
		v1:=m["k1"]
		fmt.Println("v1:",v1)
			output-> v1:7

	What if the key does not exist 
		v3:=m["k3"]
		fmt.println("v3:",v3)
			output-> v3:0
		why 0?
			because :
				value type is int 
				zero value(default value) of int is 0

	Length of a map
		fmt.Println("len:",len(m))
			output -> len:2
		len(m)=number of key-value pairs

	Deleting a key
		delete(m,"k2")
		fmt.Println("map:",m)
			Now map becomes: 
				k1->7
			delete does not give error if key doesn't exist

	Clearing the entire map
		clear(m)
		fmt.Println("map:",m)
			Output-> map:map[]
		clear removes all key-value pairs

	Checking if a key exists
		_, prs:=m["k2"]
		fmt.Println("prs:",prs)
			output-> prs:false
		When you do:
			value,ok:=m[key]
				value->value for key
				ok->true if key exists,else false
			Here 
				"k2" was deleted
				So prs =false
			_(blank identifier)means: 
				"I don't care about the value ,just tell me if it exists"

	Creating a map with values directly
		n :=map[string]int{
		"foo":1,
		"bar":2,
		}
		This is shortcut syntax.
		Same as :
			n:=make(map[string]int)
			n["foo"]=1
			n["bar"]=2
		
	Comparing maps
		n2:=map[sting]int{"foo:1","bar":2}
		if maps.Equal(n,n2){
			fmt.Println("n==n2")
		}
		This checks : Same keys , Same values
		
	Important map rules
		Allowed
			keys must be comparable
				string,int,bool
			values can be anything
		Not allowed
			slices,maps,functions as keys

	Mental Model
		Map
			|-key->value
			|-key->value
			|-key->value
		No order
		Fast lookup
		Keys must be unique
	
	*/
package main

import "fmt"

func main() {
    scores := make(map[string]int)

    scores["Alice"] = 90
    scores["Bob"] = 85

    fmt.Println(scores)

    fmt.Println("Alice score:", scores["Alice"])

    if val, ok := scores["Charlie"]; ok {
        fmt.Println(val)
    } else {
        fmt.Println("Charlie not found")
    }
}

	
