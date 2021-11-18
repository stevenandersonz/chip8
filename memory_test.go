package main

import (
    "testing"
)

func TestLoadReserved(t *testing.T) {
    memory := InitMemory()
    item := memory.reserved[0]
    if item != 0xF0 {
        t.Fatalf(`Reserved At 0 = %x it should be 0xF0`, item)
    }
}
