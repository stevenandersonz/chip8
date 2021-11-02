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
    m := InitMemory()
    m.LoadMemory(0x204, 0xF0)
    fmt.Println(m.ReadFromMemory(0x1))
    loadRom("./IBM_test.ch8")
}


