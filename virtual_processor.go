package main

func main() {
	r1, r2 := NewRegistry(), NewRegistry()
	r1.setByteNumber("00000001")
	r2.setByteNumber("11111111")
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

func Sum(a *memoryRegistry, b *memoryRegistry) *memoryRegistry {
	r3 := NewRegistry()
	pCurrent := false
	var pNext bool
	for i := 7; i >= 0; i-- {
		var s bool
		if !pCurrent {
			s = a.getBit(i) != b.getBit(i)
			pNext = a.getBit(i) && b.getBit(i)
		} else {
			s = a.getBit(i) == b.getBit(i)
			pNext = a.getBit(i) || b.getBit(i)
		}
		r3.setBit(i, s)
		pCurrent = pNext
	}
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

func (registry *memoryRegistry) setByteNumber(byteNumber string) {
	if isByteNumber(byteNumber) {
		for i := 0; i < 8; i++ {
			if byteNumber[i] == '0' {
				registry.setBit(i, false)
			} else {
				registry.setBit(i, true)
			}
		}
	}
}

func isByteNumber(byteNumber string) bool {
	if len(byteNumber) != 8 {
		return false
	}
	for _, bit := range byteNumber {
		if bit != '0' && bit != '1' {
			return false
		}
	}
	return true
}

func (registry *memoryRegistry) getBit(index int) bool {
	return registry.cells[index]
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
