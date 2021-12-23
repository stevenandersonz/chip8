package main

type Keyboard struct {
    m[16] bool
    lastKey byte
    state string
}
var keyMap = map[string] byte{ 
    "30": 0,
    "31": 1,
    "32": 2,
    "33": 3,
    "34": 4,
    "35": 5,
    "36": 6,
    "37": 7,
    "38": 8,
    "39": 9,
    "41": 10,
    "42": 11,
    "43": 12,
    "44": 13,
    "45": 14,
    "46": 15,
} 
func isKeyInKeyboard(key string) byte{
    return keyMap[key]
}

func (k *Keyboard) IsKeyPressed (key string) bool {
    if isKeyInKeyboard(key) ==0 {
        return false
    }
    return true
}

func (k *Keyboard) WriteKeyPress (key string) {
    keyCode := isKeyInKeyboard(key)
    if keyCode !=0 {
        k.lastKey = keyCode
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

