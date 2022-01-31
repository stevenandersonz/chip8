package main


type Display struct {
    screen [32][64]bool
    screenJS []interface{}
    screenBuffer *[32][64]bool
    draw bool
    rendered bool
}

func (d *Display) Clear () {
    for row := range d.screen {
        for col := range d.screen[row] {
            d.screen[row][col]=false
        }
    }
    d.Sync()
}
func (d *Display) Sync () {
    copy((*d.screenBuffer)[:], d.screen[:])
}
func (d *Display) Print (drawPixel func(bool,int, int)) {
    for row := range *d.screenBuffer {
        for col := range (*d.screenBuffer)[row] {
            pixel := (*d.screenBuffer)[row][col]
            drawPixel(pixel,row,col)
        }
    }
}

func InitDisplay (screenBuffer *[32][64]bool) *Display {
    display := new(Display)
    display.screenBuffer = screenBuffer
    display.rendered = false
    display.Clear()
    return display
}
