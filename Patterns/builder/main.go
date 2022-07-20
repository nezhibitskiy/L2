package main

import "fmt"

func main() {
	builder := NewComputerBuilder()
	comp := builder.RAM(5).MB("gigabyte").Build()
	fmt.Println(comp)
}
