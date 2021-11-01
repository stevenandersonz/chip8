package main

//big endian start a 512
const start = 0x200 

type  memory struct {
    ram[3584] byte
    fonts[512] byte
}

func (m *memory) LoadMemory (address uint16, value byte) bool {
    if address < start {
        m.fonts[address] = value
    } else {
        m.ram[address - start] = value
    }
    return true
}

func (m *memory) ReadFromMemory (address uint16) byte {
    if address < start {
        return m.fonts[address] 
    } else {
        return m.ram[address - start]
    }
}

func InitMemory() *memory {
    newMem := new(memory)
    return newMem
}
