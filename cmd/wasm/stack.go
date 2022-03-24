package main

type Stack [16]uint16


func (s *Stack) Pop(stackPtr *uint16)  (uint16) {
    top := s[*stackPtr-1]
    if *stackPtr > 0 {
        (*stackPtr)--
    }
    return top
}

func (s *Stack) Push(address uint16, stackPtr *uint16)  {
    if *stackPtr < 15 { 
        (*stackPtr)++
        s[*stackPtr-1] = address 
    }
}

func (r *Registers) IncrementStackPtr() {
    if r.StackPtr < 15 {
        r.StackPtr++
    }
}
func (r *Registers) DecrementStackPtr() {
    if r.StackPtr > 0 {
        r.StackPtr--
    }
}
