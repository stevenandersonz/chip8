type registers struct {
    vx uint8
    vf uint8
    I uint16
    PC uint16
    SP uint16
}
func InitRegisters () *registers {
    regs := new(registers)
    regs.vx = 0
    regs.vf = 0
    regs.I = 0
    regs.PC = 0
    regs.SP = 0
    return regs
}
