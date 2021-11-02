package main

func InitStack () [16] *uint16 {
    var stack [16] *uint16
    return stack
}


func Pop(stack[16] *uint16, stack_ptr *uint16)  uint16 {
    *stack_ptr--
    return *stack[*stack_ptr]
}

func Push(stack[16] *uint16, stack_ptr *uint16, address uint16) {
    *stack[*stack_ptr] = address
    *stack_ptr++
}
