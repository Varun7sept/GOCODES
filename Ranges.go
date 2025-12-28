/*			
	In Go, range is used with for loops to iterate over elements of data structure one by one
	you can use range with:
		Arrays
		Slices
		Maps
		Strings
	Example: range with slice(most commom)
		nums:=[]int{1,2,3}
		sum:=0
		for _,num:=range nums{
			sum+=num
		}
		fmt.Println("sum:,sum")
			nums:=[]int{1,2,3}
				This is a slice of integers
				Values are 2,3,4
			for _,num:=range nums
				for ->loop keyword
				range nums ->Go goes through nums one element at a time
				_,num ->values returned by range

	What does range return ?
		When you use range on a slice or array ,GO gives you:
			index,value
		
	Why is _ used?
		for _,num:=range nums
			_is called the blank identifier
			It means :"I don't care about this value"
		Here:
			We don't need the index
			We only care about num
		So Go says:
			"Okay ,I'll throw the index away."
		If you wrote this instead:
			for i,num:=range nums{
			fmt.Println(i,num)
			}
			Output would be :
				0 2
				1 3 
				2 4

	Example 2 : Using index condition
		for i,num:=range nums{
		if num==3{
			fmt.Println("index:",i)
			}
		}
			Output -> index:1
		
	Example 3 : range with a map
		kvs :=map[string]string{
			"a":"apple",
			"b":"banana",
		}
		for k,v :=range kvs{
			fmt.Printf("%s -> %s\n",k,v)
		}
		Important rule for maps
			When ranging over a map ,GO gives
				key,value
		Only keys
			for k:=range kvs{
				fmt.Println("key:",l)
			}
				output -> key:a key:b
		Only values
			for _,v:=range kvs{
				fmt.Println("value:",v)
			}

	Example 4: range with strings
		for i,c:=range "go"{
			fmt.Println(i,c)
		}
		What happens here?
			Strings are made of charaters
			Go treats them as Unicode characters(runes)
				output
					0 103  g->ASCII value 103
					1 111  o->ASCII value 111
			For characters
				fmt.Println(i,string(c))
					output
						0 g
						1 0

	What range returns
		| Data Type     | Returns      |
		| ------------- | ------------ |
		| Slice / Array | index, value |
		| Map           | key, value   |
		| String        | index, rune  |
		| Channel       | value        |

	Mental model (remember this 🔥)
		range = “Give me each item from this collection, one by one”	
