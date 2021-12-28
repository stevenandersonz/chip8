package main
import (
    "time"
    "syscall/js"
)

type Keyboard struct {
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
    "41": 0xA,
    "42": 0xB,
    "43": 0xC,
    "44": 0xD,
    "45": 0xE,
    "46": 0xF,
    "255": 0xFF,
} 
func isKeyInKeyboard(key string) byte{
    return keyMap[key]
}


func (k *Keyboard) WriteKeyPress (key string) {
    js.Global().Get("console").Call("log", key)
    keyCode := isKeyInKeyboard(key)
    js.Global().Get("console").Call("log", keyCode)
    k.lastKey = keyCode
}
func (k *Keyboard) WaitForKeyPress() byte {
    for {
        if k.lastKey != 0xFF {
            return k.lastKey
        }
        time.Sleep(time.Second / time.Duration(500))
    }
}

func InitKeyboard() *Keyboard {
    k :=  new(Keyboard)
    k.lastKey = 0xFF
    return k
}

