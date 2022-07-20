package main

import (
	"fmt"
	"math/rand"
	"time"
)

type DefaultInterface interface {
	GetType() string
	PrintInfo()
}

type DefaultStruct struct{ value int }

type Struct0 struct {
	DefaultStruct
}

type Struct1 struct {
	DefaultStruct
	extraStruct1Field int
}

type Struct2 struct {
	DefaultStruct
	extraStruct2Field bool
}

func NewStruct0() *Struct0 {
	return &Struct0{DefaultStruct{value: 0}}
}

func NewStruct1() *Struct1 {
	return &Struct1{DefaultStruct{value: 1}, rand.Int() % 2}
}

func NewStruct2(hasExtraField bool) *Struct2 {
	return &Struct2{DefaultStruct{value: 2}, hasExtraField}
}

func (Struct0) GetType() string { return "Struct0" }

func (Struct1) GetType() string { return "Struct1" }

func (Struct2) GetType() string { return "Struct2" }

func (s Struct0) PrintInfo() {
	fmt.Printf("%s; DefaultValue %d\n", s.GetType(), s.DefaultStruct.value)
}

func (s Struct1) PrintInfo() {
	fmt.Printf("%s, DefaultValue %d, ExtraField %d\n", s.GetType(), s.DefaultStruct.value, s.extraStruct1Field)
}

func (s Struct2) PrintInfo() {
	fmt.Printf("%s, DefaultValue %d, ExtraField: %v\n", s.GetType(), s.DefaultStruct.value, s.extraStruct2Field)
}

func FabricMethod() DefaultInterface {
	rand.Seed(time.Now().UnixNano())
	switch rand.Int() % 3 {
	case 0:
		return NewStruct0()
	case 1:
		return NewStruct1()
	case 2:
		return NewStruct2(rand.Int()&1 == 1)
	}
	return nil
}

func main() {
	const N = 10
	computers := make([]DefaultInterface, N)
	for i := 0; i < N; i++ {
		computers[i] = FabricMethod()
	}
	for _, v := range computers {
		v.PrintInfo()
	}
}
