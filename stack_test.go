package main

import (
    "testing"
)

func TestStackPush(t *testing.T) {
    stack := new(Stack)
    sp := uint16(0)
    stack.Push(0xFFFF, &sp)
    if stack[sp] == 0xFFFF && sp == uint16(1) {
        t.Fatalf(`Stack should be empty`)
    }
}
