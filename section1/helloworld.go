package main


import (
	"fmt"
)

type User struct {
	Name string
	Age int
}

func main() {
	
	st := User{"name", 10}

	fmt.Printf("%v\n", st)
	fmt.Printf("%#v\n", st)
	fmt.Printf("%+v\n", st)
	fmt.Printf("%T\n", st)

	
}