package main
import (
)

func updateValue(x string) string {
  x = "ram" //this will create a copy of the argument here but not update the value there
  return x 
}
