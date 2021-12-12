package main

import (
    "fmt"
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
    cpu := InitCPU()
    rom, romSize := openFile("./roms/IBM_test.ch8")
    cpu.LoadProgram(*rom, romSize)
    for cpu.regs.PC < 0xFFD {
        cpu.Cycle()
    }
    cpu.display.Print()
}



