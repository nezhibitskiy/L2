package main

import "fmt"

type ExtraHardLib struct {
}

func (e *ExtraHardLib) Do1() {
	fmt.Println("Do 1")
}

func (e *ExtraHardLib) Do2() {
	fmt.Println("Do 2")
}

func (e *ExtraHardLib) Do3() {
	fmt.Println("Do 3")
}

func (e *ExtraHardLib) Do4() {
	fmt.Println("Do 4")
}

func (e *ExtraHardLib) Do5() {
	fmt.Println("Do 5")
}

type ExtraHardLibFacade struct {
	ExtraHardLib *ExtraHardLib
}

func (e *ExtraHardLibFacade) EasyDo() {
	e.ExtraHardLib.Do1()
	e.ExtraHardLib.Do3()
	e.ExtraHardLib.Do5()
}

func main() {
	e := ExtraHardLib{}
	ef := ExtraHardLibFacade{ExtraHardLib: &e}

	ef.EasyDo()
}
