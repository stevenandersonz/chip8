package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"syscall/js"
	"time"

	"github.com/stevenandersonz/tree"
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
        return true
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
func initDisplay(p *cpu) js.Func {
    return js.FuncOf(func (this js.Value, args []js.Value) interface {} {
    jsDoc := js.Global().Get("document")
    DOMDocument := tree.TreeElement(jsDoc)
    displayUI := DOMDocument.GetElementById("chip8Display")
    if !p.display.rendered {
        p.display.Print(func (pixel bool, y int, x int) {
            id := "id=pixel-" + strconv.Itoa(y) + "-" + strconv.Itoa(x)
            pixelUI := DOMDocument.CreateElement("div", []string{id, "class=off"})
            displayUI.AppendChild(pixelUI)
        })
        p.display.rendered = true
    }
    return nil
    })
}
func refreshDisplay(p *cpu) js.Func {
    return js.FuncOf(func (this js.Value, args []js.Value) interface {} {
                jsDoc := js.Global().Get("document")
                DOMDocument := tree.TreeElement(jsDoc)
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
                return nil 
            })
}
func getKeyPress(p *cpu) js.Func {
    return js.FuncOf(func (this js.Value, args []js.Value) interface {} {
        keyASCII := args[0]
        p.keyboard.WriteKeyPress(keyASCII.String())
       time.sleep(time.Second / time.Duration(500) 
        js.Global().Get("console").Call("log", p.keyboard.lastKey)
        return nil
    })
}
func main () {
    cpu := InitCPU()
    js.Global().Set("loadROM", getROMWrapper(cpu))
    js.Global().Set("initDisplay", initDisplay(cpu))
    js.Global().Set("refreshDisplay", refreshDisplay(cpu))
    js.Global().Set("onKeypress", getKeyPress(cpu))
    <- make(chan bool)
}



