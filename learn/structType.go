package main 
import (
  "fmt"
)

type bill struct {
  name string 
  items map[string]float64
  tip float64
}

func newBill(name string) bill {
 b := bill{
  name: name,
  items: map[string]float64{"cake": 24.66, "Pie": 52.75},
  tip: 100,
 }
 return b
}

func bills() {
	myBill := newBill("san's bill ")

	fmt.Println(myBill.name, myBill.items,myBill.tip)
	format := myBill.format()
	fmt.Printf("\n-------------------------\n%v\n",format)
}


/** 
* These type of functions can only be called from
* object associated with the given type 
*
* here any object associated with bill type 
* can only access the prinData() method
*/
//example-method with type bill
func (b bill) printData(){
  fmt.Println(b.name, "is found of items ",b.items)
}
 
// formating the outputs
func (b bill) format()string {
   fs := "Your Bill: \n"
   var total float64 = 0 

   /**
   * listing items from map
   * 
   * using %-10v will give the 10 character gap to the right and %10v
   * will give 10 character gap to the left side  
   */

   //Another thing is Sprintf adds the formatted string to the string
   //and store it and then it needs to be printed seperately using prinln or pri   //ntf
   for k,v := range b.items {
	   fs += fmt.Sprintf("%-10v ...$%v\n", k+":",v)
	   total += v
   }
   fs += fmt.Sprintf("------------------------\n")
   //toal
   fs += fmt.Sprintf("%-10v ...$%0.2f\n", "total:",total)
   return fs
}
