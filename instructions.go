package main

func (p *cpu) executeInstruction (instruction string) {
    switch instruction {
        case "0000":
            break
        case "00E0":
            // CLR
            break
        case "1NNN":
            // Jump
            break
        case "6XNN":
            // set Register VX
            break
        case "7XNN":
            //add value to register VX
            break
        case "ANNN": 
            // set index register I
            break
        case "DXYN":
            // display draw
            break
    }

}
