package vm

type Slot interface{}

type Stack struct {
	slots []Slot
	top   uint64
}

func NewStack() *Stack {
	return &Stack{[]Slot{}, 0}
}

func (s *Stack) Clear() {
	s.slots = []Slot{}
}

func (s *Stack) Top() uint64 {
	return s.top
}

func (s *Stack) Push(obj ...Slot) {
	s.slots = append(s.slots, obj...)
	s.top += uint64(len(obj))
}

func (s *Stack) Pop() Slot {
	s.top--
	var obj = s.slots[s.top]
	s.slots[s.top] = nil
	return obj
}

func (s *Stack) Set(idx uint64, obj Slot) {
	s.slots[idx] = obj
}

func (s *Stack) Get(idx uint64) Slot {
	return s.slots[idx]
}

func (s *Stack) Copy(src uint64, dst uint64) {
	s.slots[dst] = s.slots[src]
}
