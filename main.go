package main

import (
    "github.com/deanveloper/karman"
    "os"
    "os/signal"
    "syscall"
)

func main() {
    bots := []Bot{karman.New()}

    for _, bot := range bots {
        go bot.Start()
    }

    // keep bots running until force closed
    sigChan := make(chan os.Signal)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
    <-sigChan

    for _, bot := range bots {
        go bot.Close()
    }
}

type Bot interface {
    Start()
    Close()
}
