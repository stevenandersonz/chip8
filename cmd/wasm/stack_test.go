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

func TestStackPop(t *testing.T) {
    stack := new(Stack)
    sp := uint16(0)
    stack.Push(0xFFFF, &sp)
    top := stack.Pop(&sp)
    if top == 0xFFFF && sp == uint16(0) {
        t.Fatalf(`Stack should be empty`)
    }
}
func TestStackPushAfterFull(t *testing.T) {
    stack := new(Stack)
    sp := uint16(0)
    for i:=0; i<16; i++ {
        stack.Push(0xFFFF, &sp)
    }
    if stack[sp] == 0xFFFF && sp == uint16(15) {
        t.Fatalf(`Stack should be empty`)
    }
}


