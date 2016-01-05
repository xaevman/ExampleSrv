package main

import (
    "github.com/xaevman/crash"
    "github.com/xaevman/srvApp"
)

func main() {
    defer crash.HandleAll()
    srvApp.Init()
}
