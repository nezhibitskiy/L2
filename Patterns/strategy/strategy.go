package main

import "fmt"

type B struct {
	eviction eviction
	data     []int
}

func (b *B) Evict() {
	b.eviction.evict(b)
}

func (b *B) AddValue(i int) {
	b.data = append(b.data, i)
}

type removeFirst struct {
}

func (r *removeFirst) evict(b *B) {
	b.data = b.data[1:]
}

type removeLast struct {
}

func (r *removeLast) evict(b *B) {
	b.data = b.data[:len(b.data)-1]
}

type eviction interface {
	evict(b *B)
}

func main() {
	b := B{eviction: &removeFirst{}, data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}}

	fmt.Println(b.data)
	b.Evict()
	fmt.Println(b.data)

	b.eviction = &removeLast{}

	b.Evict()
	fmt.Println(b.data)
}
