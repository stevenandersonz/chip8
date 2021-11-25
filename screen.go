package main
import (
    "github.com/faiface/pixel"
    "github.com/faiface/pixel/pixelgl"
)

func run () {
    cfg := pixelgl.WindowConfig{
        Title: "CHIP 8",
        Bounds: pixel.R(0,0,1024,768),
    }
    win, err := pixelgl.NewWindow(cfg)
    if err != nil {
        panic(err)
    }

    for !win.Closed() {
        win.Update()
    }

}
func InitScreen () {
    pixelgl.Run(run)
}
