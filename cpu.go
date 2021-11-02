package main 

import ( 
    "fmt" 
)

type cpu struct {
    m *memory
    regs *registers
    display *[64][32] bool
    stack *[16] uint16
}


func NewProcessor () *cpu {
    cpu := new(cpu)
    cpu.m := InitMemory()
    cpu.regs := InitRegisters()
    cpu.stack := InitStack()

}
