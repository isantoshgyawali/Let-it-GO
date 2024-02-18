package main 
import (
	"fmt"
	"math"
)

type rect struct {
  width float64
  height float64
}

type circle struct {
  radius float64
}

func (r *rect) area() float64{
	areaOfRect := r.width * r.height
  return areaOfRect
}

func (r *rect) identifier() string{
  return "rectangle"
}

func (c *circle) area() float64{
	areaOfCircle := math.Pi * c.radius * c.radius
  return areaOfCircle 
}

func (c *circle) identifier() string{
  return "circle"
}

/** 
* As we can see,
* we are making the same method for different types 
* or say in this case "shapes" 
*
* We should be allowed to call the area for different shapes
* even if we got the slices of shapes like:
*
* shape := []<someType>{circleObj,rectObj}
* 
* and then iterate or do whatever in it to achieve something like : 
* shape[0].area() or shape[1].area()
* 
* but we don't know what to assign in types so: "WE DEFINE INTERFACES"
* and that will allow us to 
* group any types of data even if we don't know it's specific undelying type
*/
type shape interface {
  area() float64 // here any types with "area method ie. area()" that returns float64 is of 
                 // type shape then it can:
                 // use this "shapes" as a upper level type to reference it's methods[:w
  identifier() string
}

func interfaces(){
    circ := circle{4.5}   //creating a object of type circle
    rec := rect{12,13.5} //creating a object of type rect
    
    fmt.Printf("area of cirle: %v \narea of rectangle : %v \n", circ.area(),rec.area())

    
    //NOW WHAT WE ARE ALLOWED AFTER USING INTERFACE IS : 
    fmt.Println("---------------------------------------------")
    shapes := []shape{&circ,&rec}

    for i, v := range shapes{

	    fmt.Println(shapes[i]==v)      // true : we can use them back and fourth
	                                   // just added to get used to with go :)
	    name := shapes[i].identifier() // returns "circle" or "rectangle"
	    area := v.area() 

	    fmt.Printf("area of %s : %f\n",name,area)
	    fmt.Printf("type of %s : %T\n\n",name,v)
    }
}


/** AND POINT TO BE NOTED : WE ARE NOT LIMITED TO USE ONLY ONE INTERFACE ON OBJECTS */
