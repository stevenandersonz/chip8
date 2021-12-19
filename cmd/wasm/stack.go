package main

type Stack [16]uint16


func (s *Stack) Pop(stackPtr *uint16)  (uint16) {
    top := s[*stackPtr]
    if *stackPtr > 0 {
        (*stackPtr)--
    }
    return top
}

func (s *Stack) Push(address uint16, stackPtr *uint16)  {
    if *stackPtr < 15 { 
        s[*stackPtr] = address 
        (*stackPtr)++
    }
}

func (r *Registers) IncrementStackPtr() {
    if r.stackPtr < 15 {
        r.stackPtr++
    }
}
func (r *Registers) DecrementStackPtr() {
    if r.stackPtr > 0 {
        r.stackPtr--
    }
}
