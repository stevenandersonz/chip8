package main 

import ( 
    "fmt" 
)

type cpu struct {
    m *memory
    registers *Registers
    display *Display
    keyboard *Keyboard
    stack *Stack
    lastOpcode string
}
func getInstructionChar(instruction uint16) string {
    return fmt.Sprintf("%04X", instruction)
}

func  InitCPU (screenBuffer *[32][64] bool) *cpu {
    p := new(cpu)
    p.m = InitMemory()
    p.registers = InitRegisters()
    p.display = InitDisplay(screenBuffer)
    p.keyboard = InitKeyboard()
    p.stack = new(Stack)
    go p.registers.RegisterClockLoop()
    return p
}

func (p *cpu) FetchInstruction () uint16 {
    mostSignificantByte := p.m.ReadFromMemory(p.registers.GetPC())
    leastSignificantByte := p.m.ReadFromMemory(p.registers.GetPC()+1)
    return uint16(mostSignificantByte) <<8 + uint16(leastSignificantByte)
}

func (p *cpu) LoadProgram (program []byte, programSize uint16) {
    p.m.LoadProgram(program, programSize)
}
func (p *cpu) Cycle () {
    p.registers.IncrementPC()
    opCode := getInstructionChar(p.FetchInstruction())
    p.Execute(opCode)
    p.lastOpcode = opCode
}


