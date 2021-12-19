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
        go RunChip8(p)
        return nil
    })

    return getROMFunc

}
func RunChip8(p *cpu) {
    clockSpeed := uint64(500)
    for p.registers.GetPC() < 0xFFD {
        p.Cycle()
        time.Sleep(time.Second / time.Duration(clockSpeed))
    }
}
func RunGraphics(p *cpu) {
    jsDoc := js.Global().Get("document")
    DOMDocument := tree.TreeElement(jsDoc)
    displayUI := DOMDocument.GetElementById("chip8Display")
    for {
        if !p.display.rendered {
            p.display.Print(func (pixel bool, y int, x int) {
                id := "id=pixel-" + strconv.Itoa(y) + "-" + strconv.Itoa(x)
                pixelUI := DOMDocument.CreateElement("div", []string{id, "class=off"})
                displayUI.AppendChild(pixelUI)
            })
            p.display.rendered = true
        } else {
            if p.display.draw {
                p.display.Print(func (pixel bool, y int,x int){
                    id := "pixel-" + strconv.Itoa(y) + "-" + strconv.Itoa(x)
                    pixelUI := DOMDocument.GetElementById(id)
                    if pixel {
                       js.Value(pixelUI).Set("className","on")
                    } else {
                       js.Value(pixelUI).Set("className","off")
                    }

                })
                    p.display.draw =false
            }
        }
            time.Sleep(time.Second / time.Duration(500))
    }
}
func main () {
    cpu := InitCPU()
    js.Global().Set("loadROM", getROMWrapper(cpu))
    go RunGraphics(cpu)
    <- make(chan bool)
}



