package main

type Display struct {
    screen screen
    screenBuffer *screen
    rendered bool
    draw bool
}

func (d *Display) Clear () {
    for row := range d.screen {
        for col := range d.screen[row] {
            d.screen[row][col]=false
        }
    }
    d.Sync()
    d.draw = true
}

//func (d *Display) Sync () {
//    copy((*d.screenBuffer)[:], d.screen[:])
//}

func (d *Display) Sync () {
    for rowIdx:=0; rowIdx<32; rowIdx++ {
        for colIdx:=0; colIdx<64; colIdx++ {
            d.screenBuffer[rowIdx][colIdx] = d.screen[rowIdx][colIdx]
        }
    }
}

func InitDisplay (screenBuffer *screen) *Display {
    display := new(Display)
    display.screenBuffer = screenBuffer
    display.rendered = false
    display.Clear()
    return display
}
