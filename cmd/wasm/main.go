package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"syscall/js"
	"time"
)
func check(e error) {
    if e != nil {
        panic(e)
    }
}
func openFile (path string) (*[]byte, uint16) {
    rom, err := os.Open(path)
    check(err)
    program := make([]byte, 0xFFF)
    programSize,err := rom.Read(program)
    check(err)
    return &program, uint16(programSize)
}

func getROMWrapper (p *cpu) js.Func {
    getROMFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        array := args[0]
		buffer := make([]uint8, array.Get("byteLength").Int())
		js.CopyBytesToGo(buffer, array)
		reader:= bytes.NewReader(buffer)
        programSize,err  := reader.Read(buffer)
        check(err)
		p.LoadProgram([]byte(buffer), uint16(programSize))
        go RunChip8(p)
        return true
    })

    return getROMFunc

}
func RunChip8(p *cpu) {
    doc:= js.Global().Get("document")
    instructionsList := doc.Call("getElementById", "instructions")
    gpReg := doc.Call("getElementById", "gp-reg")
    iReg := doc.Call("getElementById", "i-reg")
    pcReg := doc.Call("getElementById", "pc-reg")
    stackPtr := doc.Call("getElementById", "stack-ptr")
    dtReg := doc.Call("getElementById", "dt-reg")
    n := 0
    for p.registers.GetPC() < 0xFFD {
        p.Cycle()
        if instructionsList.Truthy() {
            if n == 5 {
                instructionsList.Set("innerHTML", "")
                n=0
            }
            instruction:= doc.Call("createElement", "li")
            text:= doc.Call("createTextNode", p.lastOpcode)
            instruction.Call("append", text)
            instructionsList.Call("append", instruction)
            gpReg.Set("innerHTML", fmt.Sprintf("GP: [ %v ]", p.registers.generalPurpose[:]))
            iReg.Set("innerHTML", fmt.Sprintf("I: [ %v ]", p.registers.GetI()))
            pcReg.Set("innerHTML", fmt.Sprintf("PC: [ %v ]", p.registers.GetPC()))
            stackPtr.Set("innerHTML", fmt.Sprintf("Stack Ptr: [ %v ]", p.registers.stackPtr))
            dtReg.Set("innerHTML", fmt.Sprintf("DT: [ %v ]", p.registers.GetI()))
        }
        n++
        time.Sleep(time.Second / time.Duration(p.GetClockSpeed()))
    }
}
func getKeyPress(p *cpu) js.Func {
    return js.FuncOf(func (this js.Value, args []js.Value) interface {} {
        keyASCII := args[0]
        p.keyboard.WriteKeyPress(strconv.Itoa(keyASCII.Int()))
        return nil
    })
}
func getPixel(p *cpu) js.Func {
    return js.FuncOf(func (this js.Value, args []js.Value) interface {} {
        x:= args[0]
        y:= args[1]
        pixel := p.display.screenBuffer[y.Int()][x.Int()]
        return pixel
   })
}
func getClockSpeed(p *cpu) js.Func {
    return js.FuncOf(func (this js.Value, args []js.Value) interface {} {
        return p.GetClockSpeed()
    })
}
func increaseClockSpeed(p *cpu) js.Func {
    return js.FuncOf(func (this js.Value, args []js.Value) interface {} {
        p.IncreaseClockSpeed()
        return nil 
    })
}
func decreaseClockSpeed(p *cpu) js.Func {
    return js.FuncOf(func (this js.Value, args []js.Value) interface {} {
        p.DecreaseClockSpeed()
        return nil
    })
}
func getAllPixel(p *cpu) js.Func {
    return js.FuncOf(func (this js.Value, args []js.Value) interface {} {
        screen := make([] interface {},32)
        return screen
   })
}

func main () {
    var screenBuffer[32][64] bool
    cpu := InitCPU(&screenBuffer)
    js.Global().Set("loadROM", getROMWrapper(cpu))
    js.Global().Set("onKeypress", getKeyPress(cpu))
    js.Global().Set("getPixel", getPixel(cpu))
    js.Global().Set("getClockSpeed", getClockSpeed(cpu))
    js.Global().Set("increaseClockSpeed", increaseClockSpeed(cpu))
    js.Global().Set("decreaseClockSpeed", decreaseClockSpeed(cpu))
    <- make(chan bool)
}



