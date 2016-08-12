package main

import (
    "C"
    "unsafe"
    "syscall"
)

var (
    kernel32, _ = syscall.LoadLibrary("kernel32.dll")
    
)

func main(){
    defer syscall.FreeLibrary(kernel32)

    syscall.GetProcAddress()
}
