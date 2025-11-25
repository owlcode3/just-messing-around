package main

import (
	"fmt"

	"github.com/owlcode3/jma/add"
	"github.com/owlcode3/jma/minus"
)

func main() {
	fmt.Println(add.Add(800, 900))
	fmt.Println(minus.Minus(3000, 900))
}
