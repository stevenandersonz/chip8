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
    programCounter uint16
    i uint16
    generalPurpose [16] byte
    stackPtr uint16
    delayTimer byte
    soundTimer byte
    soundBuffer *bool
}


// Get I Register  
func (r *Registers) GetI () uint16 {
    return r.i 
}
// Set I Register  
func (r *Registers) SetI (value uint16) {
    r.i = value
}
// Get Program Counter Register
func (r *Registers) GetPC () uint16 {
    return r.programCounter
} 
// Set Program Counter Register
func (r *Registers) SetPC (address uint16) {
    r.programCounter = address
}
// Increment Program Counter Register By 2
// Each instruccion is 4 bytes long or 2 Memory block
func (r *Registers) IncrementPC () {
    r.programCounter += 2
}
// Get General Purpose Register at index idx
func (r *Registers) GetGP(idx uint8) byte {
    return r.generalPurpose[idx] 
}
// Set General Purpose Register at index idx
func (r *Registers) SetGP(idx uint8, value byte) {
    r.generalPurpose[idx] = value 
}
// Set General Purpose Register at index f
// Use when carry result must be set
func (r *Registers) SetVF(carry byte) {
    r.SetGP(15,carry)
}

func (r *Registers)UpdateClockTimers () {
    if r.delayTimer > 0 {
        r.delayTimer--
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
    r.stackPtr = 0
    return r
}
