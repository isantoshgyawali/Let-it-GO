package main

import (
  "fmt"
  "strings"
)

func first(){
  fmt.Println("Hello from first")
}
 
func dataTypes(){
  var hello string = "Hello from dataTypes"
  var no int = 19

  fmt.Println(hello)
  fmt.Println("I am number :", no)
}

func colonUse(){
  /** 
  * This ":=" is only applicable inside of the function
  */
  hello := "this is colon hello"
  fmt.Println(hello)
}

func bitsAndMem(){
  /** 
  * We can set the specific int size so that 
  * we make the integer not exceed certain value 
  * like in this numOne will throw error as it's range is -128 to 127
  */

  // var numOne int8 = 215 
  // fmt.Println(numOne)
  var numTwo int64 = 240000
  fmt.Println(numTwo);

  // uint lets create a variable that can't be -ve
  var pos uint8 = 249 // I can assign 249 in uint 8 cause now -ve no. are not included so range changed to 0 to 255
  fmt.Println(pos);
}
 
func printFormat(){
  fmt.Print("no new line there. ")
  fmt.Println("new line there")

  number := 10
  name := "user"

  /** 
  * Various ways to insert the variables inside the print statement 
  * Formatted Strings 
  * %v -> for the variables
  * %q -> for the quotes [works only with the strings type]
  * %T -> stores the type of the variables
  * %f -> for the floats and also %0.3f can be passed that means 3 digits after the decimal point
  */
  fmt.Printf("Here is the %v from %q \nwe can see that the number is of type %T\n",number,name,number)
  fmt.Printf("this is float : %0.3f\n",33.2223232)

  /** 
  * we can save the string inside a variable
  * using the Sprintf
  */
  var savedString string = fmt.Sprintf("This is the save String")
  fmt.Println(savedString)
}

func arrayAndSlices(){
  //Arrays - Sizes for array is fixed [ appending new element is restricted ]
  var arr [3]int = [3]int{10,20,30}
  fmt.Println("value array second element in narr[] is: ",arr[1],"\nlength of the array arr is:",len(arr))

  var name [4]string = [4]string{"Alex","Hales","Hashim","Amla"}
  fmt.Println(name)

  //Slices (use arrays under the hood) - Sizes for array is dynamic and can be updated 
  slice := []int{10,20,30,40}
  fmt.Println(slice)

  slice = append(slice,50)
  fmt.Println(slice)

  //Slice_Ranges 
  rangeOne := name[:]    // Includes Everything 
  rangetwo := name[2:]   // includes from position 2 to end 
  rangethree := name[:3] // from beginning to position 2 don't include 3 
  rangefour := name[0:3] // includes from pos 1 to pos 2 don't include 3

  fmt.Println(rangeOne, rangetwo, rangethree, rangefour)

  //Converting the array to slice and appending to it
  sliceName := name[:]
  sliceName = append(sliceName, "faf")
  fmt.Println(sliceName)
}

func loops(){
  x := 0

  /** 
  * In GO "for loops" works as: 
  * -- while loops 
  * -- typical for 
  * -- for in 
  */
  //using forLoop as a whileloop
  for x < 5 {
    fmt.Println(x);
    x++
  } 
  fmt.Println("-----------------------")

 //using normal for loop
 for i := 0; i<=5; i++{
    fmt.Println(i);
 }
 fmt.Println("-----------------------")
 name := [4]string{"Alex","Hales","Hashim","Amla"}
 for j := 0; j<len(name); j++{
	 fmt.Println("name : ",name[j])
 }

 //using range
 for index, value := range name {
	 fmt.Printf("Name is  %v from index %v\n",value,index)
 }

 // -- if don't want to use the index and similarly in the other scenario
 for _, value := range name {
	 fmt.Printf("Name : %v\n",value)
 }
}

/** function and Parameters */
func funcOne(name string){
  fmt.Printf("hello!, %v\n",name)

}
func funcTwo(n []string, f func(string)){
  for _,v := range n {
     f(v)
  }
}

/** 
* first func add takes the slice of integer n and then returns a int sum
* func mult calls the func add for the given no of times and print product
* basically what all this functions do is : take repeatation times and numbers
* add those no. in arrays and do it for given no. of times
*/
func add(n []int) int {
  sum := 0
  for _, num := range n {
    sum += num
  }
  return sum
}

func mult(numbers []int , times int, sumfun func([]int) int){
	product := 1
	for i := 0 ; i < times ; i++ {
		product *= sumfun(numbers)
	}
	fmt.Println(product)
}

func initialsReturn(n string) (string, string) /**returning two string values*/ {
   s := strings.ToUpper(n)
   names := strings.Split(s, " ")

   var initials []string
   for _, v := range names {
     initials = append(initials, v[:1]) //getting the initial letter of the given strings
   }

   // if there are the two strings return initials of both else return initial at 0 and "_"
   if len(initials) > 1 {
	return initials[0], initials[1] // return the initials of all pos 
   }

   return initials[0], "_"
}

func yourValue(){
  value := "shyam"
  
  updateValue(value) // at ./passByValue.go
  fmt.Println(value) // don't changes value
   
  value = updateValue(value) // updates the value but consumes more space as it
                             // made a copy outside the scope
  fmt.Println(value)
}

func main(){
  //first()
  //dataTypes()
  //colonUse()
  //bitsAndMem()
  //printFormat()
  //arrayAndSlices()
  //loops()
  //funcOne("user")
  //funcTwo([]string{"Hello","user","how","are","you"}, funcOne)

  /** asking for numbers , repeatation , function that calls add */
  //mult([]int{3,4}, 2, add)

  /** passing the two string */
  //this way we can store the multiple returned values in just one Line in go
  //firstOne , secondOne := initialsReturn("hello there")
  //thirdOne, fourthOne := initialsReturn("what")

  //fmt.Println(firstOne, secondOne)
  //fmt.Println(thirdOne, fourthOne)

  /** 
  * message from another file 
  *
  * though you don't have to import the multiple files under the 
  * same package to run the multifile program but they have to be runned 
  * at same time like: "$ go run basics.go multifile.go"
  */
  //for _, v := range list{
  //  sayHelloFromAnotherFile(v)
  //}

  /** 
  * notes : you might assume as I did : if there are multiple files 
  * do I have to run by typing each file 
  * yes but using *.go (pretty straightforward and obvious) and what if 
  * files are divided into multiple dir , then : "$ go run $DIR_NAME"
  * how cool is that : go will find all the go file inside dir and subDir 
  * and run them at once 
  */

  // -- As I know how to use multiple file from this point
  // -- each program will have there own file
  //usingMap() // at ./map.go
  //insertMap() // at ./learninput.go
  
  //yourValue()

  //pointer()
  //bills()

  interfaces() //./interfaces.go
}

