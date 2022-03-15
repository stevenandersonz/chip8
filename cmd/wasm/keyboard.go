package main

type Keyboard struct {
    lastKey byte
    key chan byte
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
    keyCode := isKeyInKeyboard(key)
    k.lastKey = keyCode
    if(keyCode != 0xFF){
        k.key <- keyCode
    }
}

func InitKeyboard() *Keyboard {
    k :=  new(Keyboard)
    k.lastKey = 0xFF
    k.key = make(chan byte)
    return k
}

