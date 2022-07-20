package main

type Computer struct {
	MB  string
	RAM int
}

type ComputerBuilderI interface {
	RAM(int) ComputerBuilderI
	MB(string) ComputerBuilderI

	Build() Computer
}

type ComputerBuilder struct {
	ram int
	mb  string
}

func NewComputerBuilder() *ComputerBuilder {
	return &ComputerBuilder{}
}

func (c *ComputerBuilder) RAM(ram int) ComputerBuilderI {
	c.ram = ram
	return c
}

func (c *ComputerBuilder) MB(mb string) ComputerBuilderI {
	c.mb = mb
	return c
}

func (c *ComputerBuilder) Build() Computer {
	return Computer{
		MB:  c.mb,
		RAM: c.ram,
	}
}
