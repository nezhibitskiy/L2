package main

import "fmt"

type TaskInterface interface {
	Do() bool
}

type TaskScheduler struct {
	Tasks []TaskInterface
}

func (t *TaskScheduler) RegisterTask(task TaskInterface) {
	t.Tasks = append(t.Tasks, task)
}

func (t *TaskScheduler) TriggerEvent() {
	for _, task := range t.Tasks {
		if task.Do() {
			return
		}
	}
	fmt.Println("cannot do this task")
}

type Task1 struct {
	counter int
}

func (t *Task1) Do() bool {
	t.counter++
	if t.counter < 2 {
		fmt.Println("Task 1 done")
		return true
	}
	return false
}

type Task2 struct {
	counter int
}

func (t *Task2) Do() bool {
	t.counter++
	if t.counter < 3 {
		fmt.Println("Task 2 done")
		return true
	}
	return false
}

type Task3 struct {
	counter int
}

func (t *Task3) Do() bool {
	t.counter++
	if t.counter < 4 {
		fmt.Println("Task 3 done")
		return true
	}
	return false
}

func main() {
	var t TaskScheduler
	t.RegisterTask(&Task1{})
	t.RegisterTask(&Task2{})
	t.RegisterTask(&Task3{})

	t.TriggerEvent()
	t.TriggerEvent()
	t.TriggerEvent()
	t.TriggerEvent()
	t.TriggerEvent()
	t.TriggerEvent()
	t.TriggerEvent()
}
