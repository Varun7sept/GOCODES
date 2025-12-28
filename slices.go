/*
	In Go, a slice is like a flexible list .package gocodes
		* An array has a fixed size
		* A slice can grow and shrink
	Example
		Array
			var z[3]=[3]int{10,20,30}
				size is fixed 
				cannot add more elements
		Slice
			var z[]=[]int{10,20,30}
				size is not fixes
				can grow using append

		Creating a slice (Three common ways)

			Method 1 : Slice literal
				s:=[]int{1,2,3}
					creates a slice with values
					Length=3
				
			Method 2 : Using make
				s:=make([]int,5)
					creates a slice of length 5
					All values are initialized to 0

			Mehtod 3 : From an array
				arr :=[5]int{10,20,30,40,50}
				s:=arr[1:4]
					Result[20 30 40]
					Slice points to the same array , not a copy

		Length and Capacity
			Every slice has two numbers
				len(slice)
					Number of elements currently in the slice
				cap(slice)
					How many elements it can grow without creating a new array
			Example
				s:=make([]int,3,5)
				fmt.Println(len(s)) // 3
				fmt.Printlin(cap(s)) // 5
			Length : how much is filled
			Capacity : Total space (filled + empty)

		Zero-value Slice(nil slice)
			var s []int
				This is a nil slice
				length=0
				capacity=0
				no memory allocated yet
			fmt.Println(s==nil) // true


		Adding Elements -> append()

			Basic Example
				s:=[]int{1,2}
				s=append(s,3)
					result : [1 2 3]

			Adding multiple values 
				s=append(s,4,5,6)
		
			* append may create new array internally 
				That's why ,you must assign it back
					s=append(s,value)


		What happens when capacity is full ? 
			s:=make([]int,2,2)
			s=append(s,10)
				old capacity was full
				Go creates a new bigger array
				Copies old values
				Return a new slice
				This is automatic. You don't manage memory
			
		Copying Slices
			Wrong assumption
				a:=[]int{1,2,3}
				b:=a
					This does not copy data
					Both point to the same array
					change one -> affects the other
			Correct way -> copy()
				a:=[]int{1,2,3}
				b:=make([]int,len(a))
				copy(b,a)
					Now a and b are independent

		Slicing a Slice(sub slices)
			s:=[]int{10,20,30,40,50}
			s[start,end]
				start->included
				end->excluded
			Examples
				s[1:4] // [20 30 40]
				s[:3] //  [10 20 30]
				s[2:]//   [30 40 50]

		Capacity behavior with slicing
			s:=[]int{1,2,3,4,5}
			t:=s[1:3]
				t is [2,3]
				but cap(t)==4,
			Because capacity is counted from the start index to the end of original array

		Multidimensional Slices (slice of slices)
			matrix :=make([][]int,3)
			Each row must be created separately:
				for(i:=0;i<3;i++){
					matrix[i]=make([]int,4)
				}
				Result [
				[0 0 0 0]
				[0 0 0 0]
				[0 0 0 0]
				]
			In Go,rows can have different lengths (unlike arrays)

		Printing slices
			Slices and arrays look similar when printed:
				fmt.Println(s)
					output:[1 2 3]
				But internally:
					Array->fixed
					Slice->descriptor(pointer+len+cap)

		Mental Model
			A slice contains:
				1. Pointer->where data starts
				2. Length->how many elements visible
				3. Capacity->how many elements possible

		When should you use slices?
			Almost always
			use arrays only when:
				Fixed size
				Performance critical
				Low level code



*/