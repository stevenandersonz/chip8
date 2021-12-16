package main

import (
   "fmt"
   "time"
    "encoding/hex"
    "os"
    "github.com/stevenandersonz/tree"
    "syscall/js"
    "bytes"
    "strconv"
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
func displayOP (rom *[]byte, romSize uint16) {
    fmt.Println("bytes:", romSize)
    encodedString := hex.EncodeToString(*rom)
    fmt.Println("Encoded Hex String: ", encodedString)
}
func loadRom (path string){
    rom, romSize := openFile(path)
    displayOP(rom, romSize)
}
func getDisplay (screen [32][64]bool) ([32][64]bool) {
    return screen
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
        RunChip8(p)
        return nil
    })

    return getROMFunc

}
func RunChip8(p *cpu) {
    clockSpeed := uint64(500)
    for p.regs.PC < 0xFFD {
        js.Global().Get("console").Call("log", p.lastInstruction)
        p.Cycle()
        time.Sleep(time.Second / time.Duration(clockSpeed))
    }
}
func RunGraphics(p *cpu) {
    jsDoc := js.Global().Get("document")
    DOMDocument := tree.TreeElement(jsDoc)
    displayUI := DOMDocument.GetElementById("chip8Display")
    for {

        if p.display.draw {
            js.Value(displayUI).Set("innerHTML", "")
            p.display.Print(func (pixel bool, x int,y int){
                id := "id=pixel-" + strconv.Itoa(y) + "-" + strconv.Itoa(x)
                if pixel {
                    pixelUI := DOMDocument.CreateElement("div", []string{id, "style=background-color: pink; width:10px; height:10px;"})
                    displayUI.AppendChild(pixelUI)
                } else {
                    pixelUI := DOMDocument.CreateElement("div", []string{id, "style=background-color: black; width:10px; height:10px;"})
                    displayUI.AppendChild(pixelUI)
                }

            })
                p.display.draw =false
        }
                time.Sleep(time.Second / time.Duration(clockSpeed))
    }
}
func getDisplayWrapper (p *cpu) js.Func {
    getDisplayFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        jsDoc := js.Global().Get("document")
        if !jsDoc.Truthy() {
            return "unable to get doc"
        }
        DOMDocument := tree.TreeElement(jsDoc)
        displayUI := DOMDocument.GetElementById("chip8Display")
        js.Value(displayUI).Set("innerHTML", "")
        p.display.Print(func (pixel bool, x int,y int){
            id := "id=pixel-" + strconv.Itoa(y) + "-" + strconv.Itoa(x)
            if pixel {
                pixelUI := DOMDocument.CreateElement("div", []string{id, "style=background-color: pink; width:10px; height:10px;"})
                displayUI.AppendChild(pixelUI)
            } else {
                pixelUI := DOMDocument.CreateElement("div", []string{id, "style=background-color: black; width:10px; height:10px;"})
                displayUI.AppendChild(pixelUI)
            }

        })
        return nil
    })
    return getDisplayFunc
}

  //  rom, romSize := openFile("./roms/IBM_test.ch8")
   // cpu.LoadProgram(*rom, romSize)
 //   for cpu.regs.PC < 0xFFD {
   //     cpu.Cycle()
    //}
    //cpu.display.Print()
 
func main () {
    cpu := InitCPU()
    js.Global().Set("getChip8Display", getDisplayWrapper(cpu))
    js.Global().Set("loadROM", getROMWrapper(cpu))
//   go RunGraphics(cpu)
    <- make(chan bool)
}



