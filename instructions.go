package main
import (
    "fmt"
    "strconv"
)

func (p *cpu) Execute(opCode string) {
    fmt.Println(opCode)
    code := opCode[:1]
    if code == "0"{
        switch opCode {
            case "0000":
                break
            case "00E0":
                // CLR
                p.display.Clear()
        }
    }
    switch code {
        case "1":
            // Jump
            address, err := strconv.ParseUint(opCode[1:5], 16, 32)
            check(err)
            p.regs.SetPC(uint16(address))
        case "6":
            // set Register VX
            break
        case "7":
            //add value to register VX
            break
        case "A": 
            // set index register I
            address, err := strconv.ParseUint(opCode[1:4], 16, 32)
            check(err)
            p.regs.SetI(uint16(address))
        case "D":
            // display draw
            break
    }
}
