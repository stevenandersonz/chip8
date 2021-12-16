package main

type Keyboard struct {
    m[16] bool
    lastKey uint8
    state string
}
func isKeyInKeyboard(key byte) bool{
    return key > 0xF
}

func (k *Keyboard) IsKeyPressed (key byte) bool {
    if isKeyInKeyboard(key) {
        return false
    }
    return k.m[key]
}

func (k *Keyboard) WriteKeyPress (key byte) {
    if isKeyInKeyboard(key) {
        k.m[key] = true
        k.lastKey = key
        k.state = "keyDown"
    }
}
func (k *Keyboard) WaitForKeyPress() byte {
    for {
        if k.state == "keyDown" {
            k.state = "idle"
            return k.lastKey
        }
    }
}

func InitKeyboard() *Keyboard {
    return new(Keyboard)
}

