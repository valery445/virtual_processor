package main

func main() {
	r1, r2 := NewRegistry(), NewRegistry()
	r1.setBit(7, true)
	r2.setBit(7, true)
	r1.printRegistry()
	r2.printRegistry()
	r3 := Sum(r1, r2)
	r3.printRegistry()
}

// && - AND
// || - OR
// ! - NOT
// != XOR

type memoryRegistry struct {
	cells []bool
}

func Sum(r1 *memoryRegistry, r2 *memoryRegistry) *memoryRegistry {
	r3 := NewRegistry()
	r3.cells[6] = r1.cells[7] && r2.cells[7]
	r3.cells[7] = r1.cells[7] != r2.cells[7]
	return r3
}

func NewRegistry() *memoryRegistry {
	r := &memoryRegistry{
		cells: make([]bool, 8),
	}
	for i := 0; i < len(r.cells); i++ {
		r.cells[i] = false
	}
	return r
}

func NOR(x, y bool) bool {
	return !(x || y)
}

func (registry *memoryRegistry) setBit(index int, b bool) {
	registry.cells[index] = b
}

func (r memoryRegistry) printRegistry() {
	var binary string
	for _, bit := range r.cells {
		if bit {
			binary += "1"
		} else {
			binary += "0"
		}
	}
	println(binary)
}
