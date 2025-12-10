package main

import (
    "fmt"
    "runtime/debug"
)

func main() {
    fmt.Println("Hello, World!")

    if info, ok := debug.ReadBuildInfo(); ok {
        fmt.Println("— Build Info —")
        fmt.Println("Go toolchain:", info.GoVersion)
        fmt.Println("Main module:", info.Main.Path)
        fmt.Println("Main version:", info.Main.Version)
        for _, s := range info.Settings {
            switch s.Key {
            case "vcs", "vcs.revision", "vcs.time", "vcs.modified":
                fmt.Printf("%s = %s\n", s.Key, s.Value)
            }
        }
    } else {
        fmt.Println("No build info found.")
    }
}

