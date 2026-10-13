package main

import (
    "fmt"
    "os"
    "path/filepath"
    "syscall"
)

func main() {
    dllPath, err := filepath.Abs(
        filepath.Join("target", "debug", "rust_native_bridge.dll"),
    )
    if err != nil {
        fmt.Fprintln(os.Stderr, "DLL path error:", err)
        os.Exit(1)
    }

    library := syscall.NewLazyDLL(dllPath)

    if err := library.Load(); err != nil {
        fmt.Fprintln(os.Stderr, "DLL load failed:", err)
        os.Exit(1)
    }

    helloWorld := library.NewProc("hello_world")

    if err := helloWorld.Find(); err != nil {
        fmt.Fprintln(os.Stderr, "Function lookup failed:", err)
        os.Exit(1)
    }

    helloWorld.Call()
}
