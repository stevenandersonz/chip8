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
   //         x := opCode[1:2]
 //           value := opCode[3:5]
     //       p.regs.WriteVx(x, value)
            break
        case 0x7:
            //add value to register VX
            break
        case 0xA: 
            // set index register I
            address, err := strconv.ParseUint(opCode[1:4], 16, 32)
            check(err)
            p.regs.SetI(uint16(address))
        case 0xD:
            // display draw
            break
    }
}
