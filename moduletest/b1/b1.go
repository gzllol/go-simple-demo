package main

import (
	"fmt"
	"github.com/gzllol/go-simple-demo/moduletest/a1"
	"github.com/gzllol/go-simple-demo/moduletest/a1/submod/b1"
)

func main() {
	fmt.Println("b1 version: v1.0.0")
	a1.A1()
	b1.A1B1()
}

