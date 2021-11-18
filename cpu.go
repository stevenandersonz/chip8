package main 

import ( 
    "fmt" 
)

type cpu struct {
    m *memory
    regs *registers
    display *[64][32] bool
    stack *[16] uint16
}
func getInstructionChar(instruction uint16) string {
    return fmt.Sprintf("%04X", instruction)
}

func  NewProcessor () *cpu {
    p := new(cpu)
    p.m = InitMemory()
    p.regs = InitRegisters()
    return p
}

func (p *cpu) FetchInstruction () uint16 {
    mostSignificantByte := p.m.ReadFromMemory(p.regs.GetPC())
    leastSignificantByte := p.m.ReadFromMemory(p.regs.GetPC()+1)
    return uint16(mostSignificantByte) <<8 + uint16(leastSignificantByte)
}

func (p *cpu) LoadProgram (program []byte, programSize uint16) {
    p.m.LoadProgram(program, programSize)
}


