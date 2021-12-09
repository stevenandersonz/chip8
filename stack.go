package main

type Stack [16]uint16


func (s *Stack) Pop(stackPtr *uint16)  (uint16, bool) {
    top := s[*stackPtr]
    (*stackPtr)--
    return top, true
}
func (s *Stack) Push(address uint16, stackPtr *uint16) (bool) {
    s[*stackPtr] = address 
    (*stackPtr)++
    return true
}
