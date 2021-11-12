package main
import (
    "time"
)
type registers struct {
    GP [16] byte
    I uint16
    PC uint16
    SP uint16
    DT byte
}
func InitRegisters () *registers {
    regs := new(registers)
    regs.I = 0
    regs.PC = 0
    regs.SP = 0
    return regs
}

func (regs *registers) UpdateClockTimers () {
    if regs.DT > 0 {
        regs.DT--
    }
}

func (regs *registers) RegisterClockLoop () {
    for {
        regs.UpdateClockTimers()
        time.Sleep(time.Second/60)
    }
}

func (regs *registers) GetPC () uint16 {
    return regs.PC
} 

func (regs *registers) SetPC (address uint16) {
    regs.PC = address
}
