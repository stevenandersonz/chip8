package main
import (
    "fmt"
)

type Display struct {
    screen [64][32]bool
}

func (d *Display) Clear () {
    for row := range d.screen {
        for col := range d.screen[row] {
            d.screen[row][col]=false
        }
    }
}
func (d *Display) Print () {
    for row := range d.screen {
        for col := range d.screen[row] {
            fmt.Printf(` %v `,d.screen[row][col])
        }
        fmt.Printf(`\n`)
    }
}



func InitDisplay () *Display {
    display := new(Display)
    display.Clear()
 //   display.Print()
    return display
}
