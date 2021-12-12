package main

//big endian start a 512
const ramStartAt  = 0x200 

type  memory struct {
    ram[3584] byte
    reserved[512] byte
}
func calculateRAMOffset(address uint16) uint16 {
    return ramStartAt- address 
}
func isAddressReserved(address uint16) bool {
    return address < ramStartAt
}
func (m *memory) LoadProgram (program []byte, programSize uint16) {
    m.WriteBlockToMemory(ramStartAt, programSize, program)
}

func (m *memory) ReadFromMemory (address uint16) byte {
    if isAddressReserved(address) {
        return m.reserved[address] 
    } else {
        return m.ram[address - ramStartAt]
    }
}
func (m *memory) WriteToMemory(address uint16, value byte) bool {
    if isAddressReserved(address){
        return false
    }
    m.ram[calculateRAMOffset(address)]= value
    return true
}
func (m *memory) WriteBlockToMemory(start uint16, stop uint16, data []byte) bool {
    if start < 0x200 {
        return false
    }
    dst := m.ram[calculateRAMOffset(start):calculateRAMOffset(stop)]
    copy(dst, data)
    return true
}
func (m *memory) ReadBlockFromMemory (start uint16, stop uint16) [] byte {
    buffer := make([] byte, stop-start)
    var reservedStart, reservedStop, ramStart, ramStop uint16
    if isAddressReserved(start) {
        reservedStart = start
    }else {
        ramStart = calculateRAMOffset(start)
    }
    if isAddressReserved(stop){
        reservedStop = stop
    }else {
        ramStop = calculateRAMOffset(stop)
    }
    separator := reservedStop - reservedStart
    copy(buffer[:separator], m.reserved[reservedStart:reservedStop])
    copy(buffer[separator:], m.ram[ramStart:ramStop])
    return buffer
}

func (m *memory) LoadReserved () {
    m.reserved = [512] uint8 {
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
}
func InitMemory() *memory {
    newMem := new(memory)
    newMem.LoadReserved()
    return newMem
}
