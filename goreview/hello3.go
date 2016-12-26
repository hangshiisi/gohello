package main

import (
	"fmt"
	"github.com/hangshiisi/gohello/goreview/even"
)

func main() {
	i := 5
	fmt.Printf("Is %d even? \n %v \n", i, even.Even(i))
}
