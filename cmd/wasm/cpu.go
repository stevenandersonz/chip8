package main 

import ( 
    "fmt" 
)

type cpu struct {
    m *memory
    clockSpeed uint16
    registers *Registers
    display *Display
    keyboard *Keyboard
    stack *Stack
    lastOpcode string
    data chan string
}
func getInstructionChar(instruction uint16) string {
    return fmt.Sprintf("%04X", instruction)
}


func  InitCPU (screenBuffer *screen) *cpu {
    p := new(cpu)
    p.clockSpeed = 500
    p.m = InitMemory()
    p.registers = InitRegisters()
    p.display = InitDisplay(screenBuffer)
    p.keyboard = InitKeyboard()
    p.stack = new(Stack)
    go p.registers.RegisterClockLoop()
    return p
}
func (p *cpu) GetClockSpeed () uint16 {
    return p.clockSpeed
}
func (p *cpu) DecreaseClockSpeed ()  {
    if p.clockSpeed < 200 {
        return 
    }
    p.clockSpeed = p.clockSpeed - 100
}
func (p *cpu) IncreaseClockSpeed ()  {
    if p.clockSpeed > 900 {
        return 
    }
    p.clockSpeed = p.clockSpeed + 100
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
    p.lastOpcode = opCode
    p.Execute(opCode)
}


