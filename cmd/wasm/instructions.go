package main
import (
    "syscall/js"
    "strconv"
    "math/rand"
)
func convertStrToUint16(str string) uint16 {
    val, err := strconv.ParseUint(str,16,16)
    check(err)
    return uint16(val)
}
func splitInstruccion(opCode string) (uint8, uint16){
    code := convertStrToUint16(opCode)
    sysCode := uint8(code >> 12)
    value := uint16(code & 0x0FFF)
    return sysCode, value
}
func handleSystemInstruccion (instruccion uint16, p *cpu) {
        switch instruccion {
            case 0x000:
                break
            case 0x0E0:
                //p.display.Clear()
        }
}
func drawSprite (x uint8, y uint8, n uint8, p *cpu) {
    i := p.regs.I
    p.regs.GP[15] = 0
    spriteData := p.m.ReadBlockFromMemory(i, i+uint16(n))
    vy := p.regs.ReadGP(y)
    for _,sprite:=range(spriteData[:]) {
        vx := p.regs.ReadGP(x)
        mask:=byte(0x80)
        for mask > 0 {
            isBitOn := uint8(sprite & mask) > 0
            if isBitOn {
                isPixelOn := p.display.screen[vy][vx]
                if isPixelOn {
                    p.display.screen[vy][vx] =false 
                    p.regs.GP[15]=1
                } else {
                    p.display.screen[vy][vx] = true
                    p.regs.GP[15]=0
                }
            }
            vx = (vx + 1) % 64
            mask = mask >> 1
        }
        vy = (vy + 1) % 32
    }
}
func (p *cpu) Execute(instruction string) {
    sysCode, value := splitInstruccion(instruction)
    js.Global().Get("console").Call("log", instruction)
   
    switch sysCode {
        case 0x0:
            handleSystemInstruccion(value,p)
        case 0x1:
            // Jump
            p.regs.SetPC(value-uint16(2))
        case 0x2:
            //stack ptr ++
            p.regs.IncrementStackPtr()
            // Put PC address to the top of the stack
            p.stack.Push(p.regs.GetPC(), &p.regs.stackPtr)
            // PC == nnn
            p.regs.SetPC(value-uint16(2))
        case 0x3:
            //read x and compare vx to kk
            //if equal increase pc by 2
            x := uint8(value >> 8)
            kk := uint8(value&0x0FF)
            vx := p.regs.ReadGP(x)
            if vx == kk {
                p.regs.IncrementPC()
            }
        case 0x4:
             //read x and compare vx to kk
            //if not equal increase pc by 2
            x := uint8(value >> 8)
            kk := uint8(value&0x0FF)
            vx := p.regs.ReadGP(x)
            if vx != kk {
                p.regs.IncrementPC()
            }
        case 0x5:
             //read vx and vy compare them
            //if equal increase pc by 2
            x := uint8(value >> 8)
            y := uint8((value >> 4) & 0x00F)
            vx := p.regs.ReadGP(x)
            vy := p.regs.ReadGP(y)
            if vx == vy {
                p.regs.IncrementPC()
            }

        case 0x6:
            // set Register VX
            x := uint8(value >> 8)
            val := uint8(value)
            p.regs.WriteGP(x, val)
        case 0x7:
            //add value to register VX
            x := uint8(value >> 8)
            val := uint8(value)
            p.regs.AddToVx(x, val)
        case 0x8: 
           opCode := uint8((value & 0x00F))
           x := uint8(value >>8)
            y := uint8((value >> 4) & 0x00F)
           if opCode == 0x0 {
               p.regs.MoveVyToVx(y,x)
           } 
           if opCode == 0x1 {
            p.regs.OrVxVy(x,y) 
           } 
           if opCode == 0x2 {
               p.regs.AndVxVy(x,y)
           }
           if opCode == 0x3 {
               p.regs.XOrVxVy(x,y)
           }
           if opCode == 0x4 {
               p.regs.AddToVx(x,y)
           }
           if opCode == 0x5 {
               p.regs.SubVyVx(y,x)
           }
           if opCode == 0x6 {
               p.regs.ShiftRVx(x)
           }
           if opCode == 0x7 {
               p.regs.SubNVxVy(x,y)
           }
           if opCode == 0xE {
               p.regs.ShiftLVx(x)
           }
        case 0x9: 
            x := uint8(value >>8)
            y := uint8((value >> 4) & 0x00F)
            p.regs.SkipNextInstruction(x, y)

        case 0xA: 
            // set index register I
            p.regs.SetI(value)
        case 0xB:
            jumpTo := uint16(p.regs.ReadGP(0)) + value
            p.regs.SetPC(jumpTo)
        case 0xC:
            x := uint8(value >>8)
            kk := uint8(value&0x0FF)
            nRand := uint8(rand.Intn(256))
            p.regs.WriteGP(x,nRand&kk)
        case 0xD:
            // display draw
            p.display.draw = true
           x := uint8(value >> 8)
           y := uint8((value >> 4) & 0x00F)
           n := uint8((value & 0x00F))
           drawSprite(x,y,n,p)
        case 0xE:
           x := uint8(value >> 8)
           kk := uint8(value&0x0FF)
           vx := p.regs.ReadGP(x)
           if kk == 0x9E {
                if p.keyboard.m[vx]  {
                    p.regs.IncrementPC()
                }
            }
            if kk == 0xA1 {
                if !p.keyboard.m[vx]  {
                    p.regs.IncrementPC()
                }
            }
        case 0xF:
            x := uint8(value >> 8)
            kk := uint8(value&0x0FF)
            vx := p.regs.ReadGP(x)
            if kk == 0x07 {
                p.regs.WriteGP(x, p.regs.DT)
            }
            if kk == 0x0A {
                //halt execution until key is pressed
                key := p.keyboard.WaitForKeyPress()
                p.regs.WriteGP(x, key)
            }
            if kk == 0x15 {
                p.regs.DT = vx
            }
            if kk == 0x18 {
                break
            }
            if kk == 0x1E {
                p.regs.I += uint16(vx)
            }
            if kk == 0x29 {
                p.regs.I = uint16(vx)
            }
            if kk == 0x33 {
                bcd := [3]byte {vx/100, (vx/10)%10, vx/10} 
                start := p.regs.I
                end := p.regs.I+2
                p.m.WriteBlockToMemory(start,end, bcd[:])
            }
            if kk == 0x55 {
                p.m.WriteBlockToMemory(0, uint16(x), p.regs.GP[:x])
            }
            if kk == 0x65 {
                block := p.m.ReadBlockFromMemory(0, uint16(x))
                copy(p.regs.GP[:x], block)
            }


    }
}
