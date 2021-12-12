package main

type Keyboard struct {
    memory [16] bool
}
func isKeyInKeyboard(key byte) bool{
    return key > 0xF
}

func (k *Keyboard) IsKeyPressed (key byte) bool {
    if isKeyInKeyboard(key) {
        return false
    }
    return k.memory[key]
}

func (k *Keyboard) WriteKeyPress (key byte) {
    if isKeyInKeyboard(key) {
        k.memory[key] = true
    }
}

func InitKeyboard() *Keyboard {
    return new(Keyboard)
}

