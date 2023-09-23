package main

import (
	"fmt"
)

func main() {
	r := NewRegistry()
	r.printRegistry()
	r.trigger(0, true, false)
	r.trigger(2, true, false)
	r.printRegistry()
	r.trigger(0, false, false)
	r.trigger(2, false, false)
	r.printRegistry()
	//printTruthTable(NOR)
}

// && - AND
// || - OR
// ! - NOT
// != XOR

type logicalOperation func(bool, bool) bool

type memoryCell struct {
	Q  bool
	Q_ bool
}

type memoryRegistry struct {
	cells []memoryCell
}

func NewRegistry() *memoryRegistry {
	r := &memoryRegistry{
		cells: make([]memoryCell, 8),
	}
	for i := 0; i < len(r.cells); i++ {
		r.cells[i].Q = false
		r.cells[i].Q_ = true
	}
	return r
}

func NOR(x, y bool) bool {
	return !(x || y)
}

func printTruthTable(op logicalOperation) {
	fmt.Println(op(false, false))
	fmt.Println(op(false, true))
	fmt.Println(op(true, false))
	fmt.Println(op(true, true))
}

func (c *memoryCell) trigger(s bool, r bool) {
	if s != r {
		c.Q = s
		c.Q_ = r
	}
}

func (registry *memoryRegistry) trigger(index int, s bool, r bool) {
	if s != r {
		registry.cells[index].Q = s
		registry.cells[index].Q_ = r
	}
}

func (r memoryRegistry) printRegistry() {
	var binary string
	for _, v := range r.cells {
		if v.Q {
			binary += "1"
		} else {
			binary += "0"
		}
	}
	println(binary)
}
