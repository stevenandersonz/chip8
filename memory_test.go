package main

import (
    "testing"
    "reflect"
)

func TestLoadReserved(t *testing.T) {
    memory := InitMemory()
    font := memory.reserved[:5]
    want := []uint8 { 0xF0, 0x90, 0x90, 0x90, 0xF0 }
    if !reflect.DeepEqual(font,want) { 
        t.Fatalf(`Reserved at 0:5 is %s wants %s`, font, want)
    }
}
func TestReadFromMemory(t *testing.T) {
    m := InitMemory()
    item := m.ReadFromMemory(0)
    var want uint8 = 240
    if item != want {
        t.Fatalf(`memory at 0 is %v, got %v`, want, item)
    }
}
func TestWriteToMemory(t *testing.T) {
    m := InitMemory()
    m.WriteToMemory(0x200, 0xF0)
    value := m.ReadFromMemory(0x200)
    if value != 0xF0 {
        t.Fatalf(`try to write 0xF0, got %v instead`, value)
    }
}
func TestReadBlockFromMemory (t *testing.T) {
    m := InitMemory()
    sprite := m.ReadBlockFromMemory(0, 5)
    want := []byte { 0xF0, 0x90, 0x90, 0x90, 0xF0 }
    if !reflect.DeepEqual(sprite,want) { 
        t.Fatalf(`Reserved at 0:5 is %s wants %s`, sprite, want)
    }
}
func TestWriteBlockFromMemory (t *testing.T) {
    m := InitMemory()
    newBlock := []byte{ 0x10, 0x20, 0x60, 0x30, 0x10 }
    m.WriteBlockToMemory(0x200, 0x205, newBlock)
    block := m.ReadBlockFromMemory(0x200, 0x205)
    if !reflect.DeepEqual(newBlock, block) {
        t.Fatalf(`try to write %v, got %v instead`, newBlock, block)
    }

}
