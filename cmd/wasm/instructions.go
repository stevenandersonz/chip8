package main
import (
    "strconv"
    "math/rand"
)
type InstructionVars struct {
    x uint8
    y uint8
    n uint8
    nn uint8
    nnn uint16
}

func convertStrToUint16(str string) uint16 {
    val, err := strconv.ParseUint(str,16,16)
    check(err)
    return uint16(val)
}
func splitInstruccion(instruction string) (uint8, InstructionVars){
    instructionHex := convertStrToUint16(instruction)
    opCode := instructionHex >> 12
    nnn := uint16(instructionHex & 0x0FFF)
    vars := InstructionVars {
        nnn:nnn, 
        nn:uint8(nnn & 0x00FF),
        n:uint8((nnn & 0x00F)),
        y:uint8((nnn >> 4) & 0x00F),
        x:uint8(nnn>>8),
    }
    return uint8(opCode), vars
}
func handleSystemInstruccion (instruccion uint16, p *cpu) {
        switch instruccion {
            case 0x000:
                break
            case 0x0EE:
                addr:= p.stack.Pop(&(p.registers.stackPtr))
                p.registers.SetPC(addr)
            case 0x0E0:
                p.display.Clear()
        }
}
func AddTo(vx byte, value byte) byte {
    return vx + value
}
func Or(vx byte, vy byte) byte {
   return vx | vy
}
func And(vx byte, vy byte) byte  {
    return vx & vy
}
func XOr (vx byte, vy byte) byte  {
     return vx ^  vy
}
func Add (vx byte, vy byte) (uint8, byte) {
    sum := uint16(vx) + uint16(vy) 
    //if overflows carry 1 to VF
    carry := byte(0x0)
    if sum > 255{
        carry = 0x1
    } 
    return vx + vy, carry
}
func Sub (vx byte, vy byte) (byte, byte) {
    carry := byte(0x0)
    if vx > vy {
        carry = 0x1
    }     
    return vx - vy, carry
}
func Shr (vx byte) (byte, byte) {
    carry := byte(0x0)
    if vx & 0x0F == 0x1 {
        carry = 0x1
    }
    return vx >> 1, carry
}
func SubN (vx byte, vy byte) (byte, byte){
    carry := byte(0x0)
    if vy > vx {
        carry = 0x1
    }
    return vy - vx, carry
}
func Shl (vx byte) (byte, byte) {
    carry := byte(0x0)
    if vx >> 7  == 0x1 {
        carry = byte(0x1)
    } 
   return vx << 1, carry
}
func Sne (vx byte,  vy byte) bool {
    return vx != vy
}
func drawSprite (x uint8, y uint8, n uint8, p *cpu) {
    i := p.registers.GetI()
    p.registers.SetVF(0)
    spriteData := p.m.ReadBlockFromMemory(i, i+uint16(n))
    vy := p.registers.GetGP(y)%32
    for _,sprite:=range(spriteData[:]) {
        vx := p.registers.GetGP(x)%64
        mask:=byte(0x80)
        for mask > 0 {
            isBitOn := uint8(sprite & mask) > 0
            if isBitOn {
                isPixelOn := p.display.screen[vy][vx]
                if isPixelOn {
                    p.display.screen[vy][vx] =false 
                    p.registers.SetVF(1)
                } else {
                    p.display.screen[vy][vx] = true
                    p.registers.SetVF(0)
                }
            }
            vx = (vx + 1) % 64
            mask = mask >> 1
        }
        vy = (vy + 1) % 32
    }
}

func (p *cpu) Execute(instruction string) {
    opCode, vars := splitInstruccion(instruction)
    x := vars.x
    y := vars.y
    n := vars.n
    nn := vars.nn
    nnn := vars.nnn
    vx := p.registers.GetGP(x)
    vy := p.registers.GetGP(y)
    switch opCode {
        case 0x0:
            handleSystemInstruccion(nnn,p)
        case 0x1:
            // Jump
            p.registers.SetPC(nnn-uint16(2))
        case 0x2:
            // Put PC address to the top of the stack
            
            p.stack.Push(p.registers.GetPC(), &(p.registers.stackPtr))
            // PC == nnn
            p.registers.SetPC(nnn-uint16(2))
        case 0x3:
            //read x and compare vx to nn
            //if equal increase pc by 2
            if vx == nn {
                p.registers.IncrementPC()
            }
        case 0x4:
             //read x and compare vx to kk
            //if not equal increase pc by 2
            if vx != nn {
                p.registers.IncrementPC()
            }
        case 0x5:
             //read vx and vy compare them
            //if equal increase pc by 2
            if vx == vy {
                p.registers.IncrementPC()
            }

        case 0x6:
            // set Register VX
            p.registers.SetGP(x, nn)
        case 0x7:
            //add value to register VX
            p.registers.SetGP(x, AddTo(vx, nn))
        case 0x8: 
           if n == 0x0 {
               p.registers.SetGP(x, vy)
           } 
           if n == 0x1 {
               p.registers.SetGP(x, Or(vx, vy))
           } 
           if n == 0x2 {
               p.registers.SetGP(x, And(vx, vy))
           }
           if n == 0x3 {
               p.registers.SetGP(x, XOr(vx, vy))
           }
           if n == 0x4 {
               val, carry := Add(vx, vy)
               p.registers.SetGP(x, val)
               p.registers.SetVF(carry)
           }
           if n == 0x5 {
               val, carry := Sub(vx, vy)
               p.registers.SetGP(x, val)
               p.registers.SetVF(carry)
           }
           if n == 0x6 {
               val, carry := Shr(vx)
               p.registers.SetGP(x, val)
               p.registers.SetVF(carry)
           }
           if n == 0x7 {
               val, carry := SubN(vx, vy)
               p.registers.SetGP(x, val)
               p.registers.SetVF(carry)
           }
           if n == 0xE {
               val, carry := Shl(vx)
               p.registers.SetGP(x, val)
               p.registers.SetVF(carry)
           }
        case 0x9: 
            if(Sne(vx, vy)) {
                p.registers.IncrementPC()
            }
        case 0xA: 
            // set index register I
            p.registers.SetI(nnn)
        case 0xB:
            jumpTo := uint16(p.registers.GetGP(0)) + nnn
            p.registers.SetPC(jumpTo)
        case 0xC:
            nRand := uint8(rand.Intn(255))
            p.registers.SetGP(x,nRand&nn)
        case 0xD:
            // display draw
            p.display.draw = true
           drawSprite(x,y,n,p)
        case 0xE:
           if nn == 0x9E {
                if p.keyboard.m[vx]  {
                    p.registers.IncrementPC()
                }
            }
            if nn == 0xA1 {
                if !p.keyboard.m[vx]  {
                    p.registers.IncrementPC()
                }
            }
        case 0xF:
            if nn == 0x07 {
                p.registers.SetGP(x, p.registers.delayTimer)
            }
            if nn == 0x0A {
                //halt execution until key is pressed
                key := p.keyboard.WaitForKeyPress()
                p.registers.SetGP(x, key)
            }
            if nn == 0x15 {
                p.registers.delayTimer = vx
            }
            if nn == 0x18 {
                break
            }
            if nn == 0x1E {
                p.registers.SetI(uint16(vx) + p.registers.GetI())
            }
            if nn == 0x29 {
                p.registers.SetI(uint16(vx))
            }
            if nn == 0x33 {
                bcd := [3]byte {vx/100, (vx/10)%10, vx/10} 
                start := p.registers.GetI()
                end := p.registers.GetI()+2
                p.m.WriteBlockToMemory(start,end, bcd[:])
            }
            if nn == 0x55 {
                p.m.WriteBlockToMemory(0, uint16(x), p.registers.generalPurpose[:x])
            }
            if nn == 0x65 {
                block := p.m.ReadBlockFromMemory(0, uint16(x))
                copy(p.registers.generalPurpose[:x], block)
            }


    }
}
