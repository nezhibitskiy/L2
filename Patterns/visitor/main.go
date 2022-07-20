package main

import (
	"fmt"
	"log"
)

type visitor interface {
	visitForSquare(square *square)
	visitForCircle(circle *circle)
	visitForRectangle(rect *rectangle)

	calcWithoutType(shapes []shape)
}

type shape interface {
	getType() string
	accept(visitor visitor)
}

type circle struct {
}

func (c *circle) getType() string {
	return "circle"
}

func (c *circle) accept(visitor visitor) {
	visitor.visitForCircle(c)
}

type square struct {
}

func (s *square) getType() string {
	return "square"
}

func (s *square) accept(visitor visitor) {
	visitor.visitForSquare(s)
}

type rectangle struct {
}

func (r *rectangle) getType() string {
	return "rectangle"
}

func (r *rectangle) accept(visitor visitor) {
	visitor.visitForRectangle(r)
}

type areaCalc struct {
	area int
}

func (a *areaCalc) visitForSquare(s *square) {
	fmt.Println("area calculating for square")
}

func (a *areaCalc) visitForCircle(c *circle) {
	fmt.Println("area calculating for circle")
}

func (a *areaCalc) visitForRectangle(r *rectangle) {
	fmt.Println("area calculating for rectnagle")
}

func (a *areaCalc) calcWithoutType(shapes []shape) {
	for _, s := range shapes {
		switch s.getType() {
		case "circle":
			c, ok := s.(*circle)
			if !ok {
				log.Fatal("cannot cast shape to circle")
			}
			a.visitForCircle(c)
		case "square":
			s, ok := s.(*square)
			if !ok {
				log.Fatal("cannot cast shape to rectangle")
			}
			a.visitForSquare(s)
		case "rectangle":
			r, ok := s.(*rectangle)
			if !ok {
				log.Fatal("cannot cast shape to rectangle")
			}
			a.visitForRectangle(r)
		}
	}
}

func main() {
	c := circle{}
	s := square{}
	r := rectangle{}

	a := areaCalc{}

	c.accept(&a)
	s.accept(&a)
	r.accept(&a)

	a.calcWithoutType([]shape{&c, &s, &r})
}
