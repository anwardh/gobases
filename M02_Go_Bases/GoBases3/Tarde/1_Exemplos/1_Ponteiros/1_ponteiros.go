package main

import "fmt"

func main() {

	var p *int

	var v = new(int)

	p3 := &v

	fmt.Printf("%T %T %T", p, v, p3)

}
