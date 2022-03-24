package main
import (
    "time"
)
// Chip8 Registers
// programCounter -> store currently executing address
// generalPurpose -> usually referred as Vx 
// i -> generally used to store memory addresses
// stackPtr -> points to the top most element in the Stack
// delayTimer -> is active whenever the delay timer is non-zero
// soundTimer -> is active whenever the sound timer is non-zero
type Registers struct {
    ProgramCounter uint16 
    I uint16 
    GeneralPurpose [16] byte 
    StackPtr uint16 
    DelayTimer byte
    SoundTimer byte
    SoundBuffer *bool
}


// Get I Register  
func (r *Registers) GetI () uint16 {
    return r.I 
}
// Set I Register  
func (r *Registers) SetI (value uint16) {
    r.I = value
}
// Get Program Counter Register
func (r *Registers) GetPC () uint16 {
    return r.ProgramCounter
} 
// Set Program Counter Register
func (r *Registers) SetPC (address uint16) {
    r.ProgramCounter = address
}
// Increment Program Counter Register By 2
// Each instruccion is 4 bytes long or 2 Memory block
func (r *Registers) IncrementPC () {
    r.ProgramCounter += 2
}
// Get General Purpose Register at index idx
func (r *Registers) GetGP(idx uint8) byte {
    return r.GeneralPurpose[idx] 
}
// Set General Purpose Register at index idx
func (r *Registers) SetGP(idx uint8, value byte) {
    r.GeneralPurpose[idx] = value 
}
// Set General Purpose Register at index f
// Use when carry result must be set
func (r *Registers) SetVF(carry byte) {
    r.SetGP(15,carry)
}

func (r *Registers)UpdateClockTimers () {
    if r.DelayTimer > 0 {
        r.DelayTimer--
    }
}

func (r *Registers) RegisterClockLoop () {
    for {
        r.UpdateClockTimers()
        time.Sleep(time.Second/60)
    }
}
// Creates and return a Register Pointer
func InitRegisters () *Registers {
    r := new(Registers)
    r.SetPC(0x198)
    r.StackPtr = 0
    return r
}
