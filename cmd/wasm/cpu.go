package main 

import ( 
    "fmt" 
)

type cpu struct {
    m *memory
    regs *registers
    display *Display
    stack *Stack
}
func getInstructionChar(instruction uint16) string {
    return fmt.Sprintf("%04X", instruction)
}

func  InitCPU () *cpu {
    p := new(cpu)
    p.m = InitMemory()
    p.regs = InitRegisters()
    p.display = InitDisplay()
    p.stack = new(Stack)
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
func (p *cpu) Cycle () {
    opCode := getInstructionChar(p.FetchInstruction())
    p.Execute(opCode)
    p.regs.IncrementPC()
}

