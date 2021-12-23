package main


type Display struct {
    screen [32][64]bool
    draw bool
    rendered bool
}

func (d *Display) Clear () {
    for row := range d.screen {
        for col := range d.screen[row] {
            d.screen[row][col]=false
        }
    }
}
func (d *Display) Print (drawPixel func(bool,int, int)) {
    for row := range d.screen {
        for col := range d.screen[row] {
            pixel := d.screen[row][col]
            drawPixel(pixel,row,col)
        }
    }
}

func InitDisplay () *Display {
    display := new(Display)
    display.rendered = false
    display.Clear()
    return display
}
