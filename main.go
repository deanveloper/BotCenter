package main

import (
    "github.com/deanveloper/karman"
    "github.com/deanveloper/stacy"
    "log"
    "os"
    "os/signal"
    "syscall"
)

func main() {
    bots := []Bot {
        karman.New(log.New(os.Stdout, "[Karman]", log.Ldate | log.Ltime)),
        stacy.New(log.New(os.Stdout, "[Stacy]", log.Ldate | log.Ltime)),
    }

    for _, bot := range bots {
        go bot.Start()
    }

    // keep bots running until force closed
    sigChan := make(chan os.Signal)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
    <-sigChan

    for _, bot := range bots {
        go bot.Stop()
    }
}

type Bot interface {
    Start()
    Stop()
}
