package main

import (
    "os"

    "github.com/xaevman/crash"
    "github.com/xaevman/srvApp"
)

func main() {
    defer crash.HandleAll()
    srvApp.Init()
    os.Exit(srvApp.Run())
}
