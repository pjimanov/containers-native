package main

import (
    "fmt"
    "net"
    "os"
)

func main() {
    if _, err := net.Listen("tcp", ":8080"); err != nil {
        fmt.Fprintln(os.Stdout, err)
        os.Exit(2)
    }
    fmt.Println("listening")
}

