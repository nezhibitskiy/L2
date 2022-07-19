package main

import (
	"fmt"
	"sync"
	"time"
)

// Реализовать функцию, которая будет объединять один или более done-каналов в single-канал, если один из его составляющих каналов закроется.

func main() {
	ch1, ch2, ch3, ch4 := make(chan interface{}), make(chan interface{}), make(chan interface{}), make(chan interface{})
	channels := []chan interface{}{ch1, ch2, ch3, ch4}
	chdata := [][]interface{}{{1, 1, 1, 1}, {"2", 2.2, "2", "2"}, {"3, 3, 3", 3, 3, 3}, {-4, "4"}}
	for i, _ := range chdata {
		go workerWriter(channels[i], chdata[i])
	}
	for v := range Or(ch1, ch2, ch3, ch4) {
		fmt.Println(v)
	}

}

func Or(channels ...<-chan interface{}) <-chan interface{} {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	done := make(chan struct{})
	for _, ch := range channels {
		go workerListen(ch, done)
	}
	go func() {
		for {
			select {
			case <-done:
				wg.Done()
				return
			}
		}
	}()
	wg.Wait()
	return merge(channels)
}

func workerWriter(ch chan<- interface{}, data []interface{}) {
	defer close(ch)
	for i, _ := range data {
		ch <- data[i]
	}
}

func workerListen(ch <-chan interface{}, done chan struct{}) {
	for {
		select {
		case <-done:
			return
		case _, ok := <-ch:
			if !ok {
				done <- struct{}{}
			}
		default:
			time.Sleep(1 * time.Second)
		}
	}
}

func merge(chs []<-chan interface{}) <-chan interface{} {
	merged := make(chan interface{})
	go func() {
		wg := new(sync.WaitGroup)
		wg.Add(len(chs))
		for _, ch := range chs {
			go func(ch <-chan interface{}) {
				defer wg.Done()
				for data := range ch {
					merged <- data
				}
			}(ch)
		}
		wg.Wait()
		close(merged)
	}()
	return merged
}
