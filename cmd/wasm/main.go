package main

type screen [32][64]bool
type Emulator struct {
    state string
    cpu *cpu
    screen *screen
}
func check(e error) {
    if e != nil {
        panic(e)
    }
}
func InitEmulator () *Emulator {
    emu := new(Emulator)
    var screenBuffer screen
    emu.cpu = InitCPU(&screenBuffer)
    emu.screen = &screenBuffer
    emu.state = "RUNNING"
    return emu
}

func main () {
    emu := InitEmulator() 
    InitC8API(emu)
    <- make(chan bool)
}



