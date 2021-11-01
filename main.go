package main

import (
    "fmt"
    "reflect"
    "encoding/hex"
    "os"
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
type Registers struct {
    pc uint16
    vx uint16
    i uint16
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
func main () {

    var ram[3584] byte 
    m := InitMemory()
    m.LoadMemory(0x204, 0xF0)
    fmt.Println(m.ReadFromMemory(0x204))
    fonts := [512] uint8 {
		0xF0, 0x90, 0x90, 0x90, 0xF0, // 0
        0x20, 0x60, 0x20, 0x20, 0x70, // 1
        0xF0, 0x10, 0xF0, 0x80, 0xF0, // 2
        0xF0, 0x10, 0xF0, 0x10, 0xF0, // 3
        0x90, 0x90, 0xF0, 0x10, 0x10, // 4
        0xF0, 0x80, 0xF0, 0x10, 0xF0, // 5
        0xF0, 0x80, 0xF0, 0x90, 0xF0, // 6
        0xF0, 0x10, 0x20, 0x40, 0x40, // 7
        0xF0, 0x90, 0xF0, 0x90, 0xF0, // 8
        0xF0, 0x90, 0xF0, 0x10, 0xF0, // 9
        0xF0, 0x90, 0xF0, 0x90, 0x90, // A
        0xE0, 0x90, 0xE0, 0x90, 0xE0, // B
        0xF0, 0x80, 0x80, 0x80, 0xF0, // C
        0xE0, 0x90, 0x90, 0x90, 0xE0, // D
        0xF0, 0x80, 0xF0, 0x80, 0xF0, // E
        0xF0, 0x80, 0xF0, 0x80, 0x80,  // F
    } 
    var screen[64][32] bool
    var stack[16] uint16
    var stack_ptr uint16
    var sound_timer uint8
    var delay_timer uint8
    
    fmt.Println("Works!")
    fmt.Println(reflect.TypeOf(ram))
    fmt.Println(reflect.TypeOf(fonts))
    fmt.Println(reflect.TypeOf(screen))
    fmt.Println(reflect.TypeOf(stack))
    fmt.Println(reflect.TypeOf(stack_ptr))
    fmt.Println(reflect.TypeOf(sound_timer))
    fmt.Println(reflect.TypeOf(delay_timer))
    loadRom("./IBM_test.ch8")
}

func PopStack (stack[16] *uint16, stack_ptr *uint16)  uint16 {
    *stack_ptr--
    return *stack[*stack_ptr]
}

func PushStack (stack[16] *uint16, stack_ptr *uint16, address uint16) {
    *stack[*stack_ptr] = address
    *stack_ptr++
}
