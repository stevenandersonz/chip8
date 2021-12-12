package main
import (
    "fmt"
)

type Display struct {
    screen [32][64]bool
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
            pixel := d.screen[row][col]
            if pixel {
                fmt.Printf("*")
            } else {
                fmt.Printf(" ")
            }
        }
        fmt.Printf("\n")
    }
}



func InitDisplay () *Display {
    display := new(Display)
    display.Clear()
    return display
}
