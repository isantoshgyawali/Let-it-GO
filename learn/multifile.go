package main
import (
  "fmt"
)

var list = []string{"Java","Cpp","Go"}

func sayHelloFromAnotherFile(lan string) {
  fmt.Printf("Is %v fast?\n", lan)
}
