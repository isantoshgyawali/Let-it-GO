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
  items: map[string]float64{},
  tip: 0,
 }
 return b
}

func bills() {
	myBill := newBill("san's bill ")

	fmt.Printf("Items in bill")
	fmt.Printf("\n%v \n %v \n %v \n",myBill.name, myBill.items,myBill.tip)


	myBill.updateItems("coffee", 15.63)
	myBill.updateItems("pasta", 37.63)
	myBill.updateItems("noodles", 43.63)
	myBill.updateTip(12.57)

	format := myBill.format()
	fmt.Printf("-------------------------\n%v\n",format)
}


/** 
* These type of functions can only be called from
* object associated with the given type 
*
* here, any objects associated with bill type 
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

   //adding tips
   fs += fmt.Sprintf("%-10v ...$%0.2f\n", "total:",b.tip)
   fs += fmt.Sprintf("------------------------\n")

   //toal
   fs += fmt.Sprintf("%-10v ...$%0.2f\n", "total:",total+b.tip)
   return fs
}

// -----------------------------------------------------------
// -----------------------------------------------------------

/** 
* Learning to use pointers with struct 
*
* Appending the data to the map and
* adding updating the tip
**/
func (b *bill) updateTip(tip float64){

  /**
  * created a copy of bill so updating without using 
  * refrence of bill here will not update the bill itself
  *
  * using pointers will allow you to save space as each time you
  * call pointers method the copy of object is not created which is
  * if the method is complex and large, the data underlying it will not be 
  * created at another memory location saving memory and copying data is 
  * more time complex than saving just memory location which ultimately improves
  * execution time of the program
  */
  b.tip = tip
}
 
func (b *bill) updateItems(name string , price float64){
  b.items[name] = price 
}
