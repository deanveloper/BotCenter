package main

import (
    "bufio"
    "fmt"
    "github.com/deanveloper/karman"
    "log"
    "os"
    "strings"
)

func main() {
    bots := map[string]Bot {
        "karman": karman.New(log.New(os.Stdout, "[Karman]", log.Ldate | log.Ltime)),
    }

    for _, bot := range bots {
        go bot.Start()
    }

    defer func() {
        fmt.Println("Stopping bots gracefully...")
        for _, bot := range bots {
            bot.Stop()
        }
    }()

    // keep bot running until "stop" is typed
    scan := bufio.NewScanner(os.Stdin)
    for scan.Scan() {
        input := scan.Text()
        if input == "stop" {
            break
        }
        bot := bots[input]
        if bot != nil {
            cmdr, ok := bot.(Commander)
            if ok {
                cmdr.Command(strings.Split(input, " ")[1:])
            } else {
                fmt.Println("That bot does not take commands")
            }
        }
    }
}

type Bot interface {
    Start()
    Stop()
}

type Commander interface {
    Command(args []string)
}