package main
import (
    "fmt"
    "strconv"
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
    fmt.Printf("SYSCODE: %X\n",sysCode)
    fmt.Printf("VAL: %X\n",value)
    return sysCode, value
}
func handleSystemInstruccion (instruccion uint16, p *cpu) {
        switch instruccion {
            case 0x000:
                break
            case 0x0E0:
                // CLR
                fmt.Println("clear")
                p.display.Clear()
        }
}
func drawSprite (x uint8, y uint8, n uint8, p *cpu) {
    vx := p.regs.ReadVx(x) & 63
    vy := p.regs.ReadVx(y) & 31
    i := p.regs.I
    p.regs.GP[15] = 0
    for j:=uint8(0); j < n; j++ {
        spriteData:= p.m.ReadFromMemory(i)
        mask:=byte(0x80)
        for mask > 0 {
            isBitOn := uint8(spriteData & mask) > 0
            if isBitOn {
                isPixelOn := p.display.screen[vy][vx]
                if isPixelOn {
                    p.display.screen[vy][vx] =false 
                    p.regs.GP[15]=15
                } else {
                    p.display.screen[vy][vx] = true
                }
            }
            vx += 1
            mask = mask >> 1
        }
        vy +=1
    }
}
func (p *cpu) Execute(opCode string) {
    sysCode, value := splitInstruccion(opCode)
    
    switch sysCode {
        case 0x0:
            handleSystemInstruccion(value,p)
        case 0x1:
            // Jump
            p.regs.SetPC(value)
        case 0x6:
            // set Register VX
            x := uint8(value >> 8)
            val := uint8(value)
            p.regs.WriteVx(x, val)
        case 0x7:
            //add value to register VX
            x := uint8(value >> 8)
            val := uint8(value)
            p.regs.AddToVx(x, val)
            break
        case 0xA: 
            // set index register I
            p.regs.SetI(value)
        case 0xD:
            // display draw
            x := uint8(value >> 8)
            y := uint8((value & 0x0F) >> 4)
            n := uint8((value & 0x00F))
            
            drawSprite(x,y,n,p)
            break
    }
}
