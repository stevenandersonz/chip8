package main
import (
    "time"
)
type registers struct {
    GP [16] byte //vx -> x == idx
    I uint16
    PC uint16
    SP uint16
    stackPtr uint16
    DT byte
    ST byte
    soundBuffer *bool
}
func InitRegisters () *registers {
    regs := new(registers)
    regs.I = 0
    regs.PC = 0x198
    regs.stackPtr = 0
    return regs
}

func (r *registers) IncrementStackPtr() {
    if r.stackPtr < 15 {
        r.stackPtr++
    }
}
func (r *registers) DecrementStackPtr() {
    if r.stackPtr > 0 {
        r.stackPtr--
    }
}
func (r *registers) SetI (address uint16) {
    r.I = address
}
func (r *registers) IncrementPC () {
    r.PC += 2
}
func (r *registers) WriteVx(vx uint8, value byte) {
    r.GP[vx] = value 
}
func (r *registers) ReadVx(vx uint8) byte {
    return r.GP[vx] 
}
func (r *registers) ReadGP(vIdx uint8) byte {
    return r.GP[vIdx] 
}
func (r *registers) WriteGP(vx uint8, value byte) {
    r.GP[vx] = value 
}
func (r *registers) AddToVx(vx uint8, value byte) {
    r.GP[vx] += value
}
func (r *registers) MoveVyToVx(vy uint8, vx uint8){
    r.GP[vx] = r.GP[vy]
}
func (r *registers) OrVxVy(vx byte, vy byte) {
    r.GP[vx] = r.GP[vx] | r.GP[vy]
}
func (r* registers) AndVxVy(vx byte, vy byte)  {
    r.GP[vx] = r.GP[vx] & r.GP[vy]
}
func (r* registers) XOrVxVy (vx byte, vy byte)  {
     r.GP[vx] = r.GP[vx] ^  r.GP[vy]
}
func (r* registers) AddVyVx (vy byte, vx byte)  {
    sum := uint16(r.GP[vx]) + uint16(r.GP[vy]) 
    //if overflows carry 1 to VF
    if sum > 255{
        r.GP[15] = 1
    } else {
        r.GP[vx] += r.GP[vy]
    }
}
func (r* registers) SubVyVx (vy uint8, vx uint8) {
    if r.GP[vx] > r.GP[vy] {
        r.GP[15] = 1
    } else {
        r.GP[15] = 0
    }
    r.GP[vx] = r.GP[vx] - r.GP[vy]
}
func (r *registers) ShiftRVx (vx byte) {
    if r.GP[vx] & 0x01 == 1 {
        r.GP[15] = 1
    } else {
        r.GP[15] = 0
    }
    r.GP[vx] = r.GP[vx] >> 1
}
func (r *registers) SubNVxVy (vx byte, vy byte) {
    if vy > vx {
        r.GP[15] = 1
    } else {
        r.GP[15] = 0
    }
    r.GP[vx] = r.GP[vy] - r.GP[vx]
}
func (r *registers) ShiftLVx (vx byte) {
    if r.GP[vx] & 0x1 == 1 {
        r.GP[15] = 1
    } else {
        r.GP[15] = 0
    }
    r.GP[vx] = r.GP[vx] << 1
}
func (r *registers) SkipNextInstruction (vx uint8, vy uint8) {
    if r.GP[vx] != r.GP[vy] {
        r.PC = r.PC << 1
    }
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
