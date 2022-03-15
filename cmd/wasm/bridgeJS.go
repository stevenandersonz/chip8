package main
import (
    "syscall/js"
    "fmt"
    "os"
    "bytes"
    "strconv"
    "time"
)

func ConsoleLog (log string) {
    console := js.Global().Get("console")
    console.Call("log", fmt.Sprintf(log))
}

type API struct {
	name string
	entryPoint js.Func
    functions map[string] interface {}
}

func (api *API) Add (fnName string, fn js.Func) {
	api.functions[fnName] = fn
}

// Publish the API and make it accesible from the DOM as a function
// API.name will be use as the name of the function
// calling foo() where foo is API.name will return an object 
// Properties in the object are map by API.Add
func (api *API) Publish () {
	var publicFunctions = make(map[string]interface{})
	for jsName, fn:=range(api.functions){
		publicFunctions[jsName] = fn
	}
	api.entryPoint = js.FuncOf(func (this js.Value, args []js.Value) interface {} {	
		return publicFunctions
	})
	js.Global().Set(api.name, api.entryPoint)
}

func NewAPI (name string) *API {
	api := new(API)
	api.name = name
	api.functions = make(map[string]interface{})
	return api
}

func getPixel(display *screen) js.Func {
    return js.FuncOf(func (this js.Value, args []js.Value) interface {} {
        x:= args[0]
        y:= args[1]
        pixel := display[y.Int()][x.Int()]
        return pixel
    })
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
        ConsoleLog(p.lastOpcode)
        
        p.keyboard.WriteKeyPress(strconv.Itoa(keyASCII.Int()))
        return nil
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

func InitC8API (emu *Emulator) {
    api := NewAPI("chip8")
    api.Add("getPixel",getPixel(emu.screen))
    api.Add("onKeyPress",getKeyPress(emu.cpu))
    api.Add("getClockRate",getClockSpeed(emu.cpu))
    api.Add("increaseClockRate",increaseClockSpeed(emu.cpu))
    api.Add("decreaseClockRate",decreaseClockSpeed(emu.cpu))
    api.Add("loadRom",getROMWrapper(emu.cpu))
    api.Publish()
}


