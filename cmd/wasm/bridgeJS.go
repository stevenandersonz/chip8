package main
import (
    "syscall/js"
    "encoding/json"
    "fmt"
    "os"
    "bytes"
    "strconv"
    "time"
)

func ConsoleLog (log interface {} ) {
    console := js.Global().Get("console")
    console.Call("log", fmt.Sprintf("%v", log))
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
func shouldDraw(p *cpu) js.Func {
    return js.FuncOf(func (this js.Value, args []js.Value) interface {} {
        draw := p.display.draw
        p.display.draw = false
        return draw 
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


func getROMWrapper (emu *Emulator) js.Func {
    getROMFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        array := args[0]
		buffer := make([]uint8, array.Get("byteLength").Int())
		js.CopyBytesToGo(buffer, array)
		reader:= bytes.NewReader(buffer)
        programSize,err  := reader.Read(buffer)
        check(err)
		emu.cpu.LoadProgram([]byte(buffer), uint16(programSize))
        go RunChip8(emu)
        return true
    })

    return getROMFunc

}
func RunChip8(emu *Emulator) {
    for emu.cpu.registers.GetPC() < 0xFFD {
        if(emu.state != "PAUSED"){
            emu.cpu.Cycle()
        }
        time.Sleep(time.Second / time.Duration(emu.cpu.GetClockSpeed()))
    }
}
func getKeyPress(p *cpu) js.Func {
    return js.FuncOf(func (this js.Value, args []js.Value) interface {} {
        keyASCII := args[0]
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
func getAllPixel(pixels *screen) js.Func {
    return js.FuncOf(func (this js.Value, args []js.Value) interface {} {
        strPixels, err := json.Marshal(pixels) 
        check(err)
        return string(strPixels)
   })
}

func setEmulatorState(emu *Emulator) js.Func {
    return js.FuncOf(func (this js.Value, args []js.Value) interface {} {
        newState := args[0].String()
        emu.state = newState
        return nil
   })
}

func getCPUState (p *cpu) map[string]interface{} {
        registers, err := json.Marshal(p.registers)
        check(err)
        instruction := GetInstructionChar(p.FetchInstruction())
        state := make(map[string]interface{})
        state["registers"] = string(registers)
        state["instruction"] = instruction
        return state
}
func executeNextInstruction(emu *Emulator) js.Func {
    return js.FuncOf(func (this js.Value, args []js.Value) interface {} {
        stateT0 := getCPUState(emu.cpu)
        emu.cpu.Cycle()
        stateT1 := getCPUState(emu.cpu)
        state := make(map[string]interface{})
        state["state0"] = stateT0
        state["state1"] = stateT1
        return state
   })
}
func InitC8API (emu *Emulator) {
    api := NewAPI("chip8")
    api.Add("getPixel",getPixel(emu.screen))
    api.Add("getScreen",getAllPixel(emu.screen))
    api.Add("shouldDraw",shouldDraw(emu.cpu))
    api.Add("onKeyPress",getKeyPress(emu.cpu))
    api.Add("getClockRate",getClockSpeed(emu.cpu))
    api.Add("increaseClockRate",increaseClockSpeed(emu.cpu))
    api.Add("nextInstruction",executeNextInstruction(emu))
    api.Add("decreaseClockRate",decreaseClockSpeed(emu.cpu))
    api.Add("loadRom",getROMWrapper(emu))
    api.Add("setEmulatorState",setEmulatorState(emu))
    api.Publish()
}


