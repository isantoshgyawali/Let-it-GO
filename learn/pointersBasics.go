package main 
import (
  "fmt"
)

func pointerData(x *string){
  *x = "table" //defrencing the pointer
}

func pointer(){
  item := "chair"
  pointer := &item

  fmt.Println("\nMemory address of item: ", pointer)
  fmt.Println("Value at address: ", *pointer)
  pointerData(pointer)
  fmt.Println("Updated Value at address: ", *pointer)

  it2 := "update"
  fmt.Println("\nMemory address of item: ", &it2)
  fmt.Println("Value at address: ", *&it2)
}
