package main
import (
    "syscall/js"
    "fmt"
)


func ConsoleLog (log string) {
    console := js.Global().Get("console")
    console.Call("log", fmt.Sprintf(log))
}

