package main

import "fmt"

type command interface {
	execute()
}

// --- Device ---

type device interface {
	on()
	off()
}

// --- Button ---

type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}

// -- On Command ---

type onCommand struct {
	device device
}

func (c *onCommand) execute() {
	c.device.on()
}

// -- On Command ---

type offCommand struct {
	device device
}

func (c *offCommand) execute() {
	c.device.off()
}

type TV struct {
}

func (t *TV) on() {
	fmt.Println("Device turned on")
}

func (t *TV) off() {
	fmt.Println("Device turned off")
}

func main() {
	tv := &TV{}

	onCmd := onCommand{device: tv}
	offCmd := offCommand{device: tv}
	onButton := &button{command: &onCmd}
	offButton := &button{command: &offCmd}

	onButton.press()
	offButton.press()
}
