package main

import (
    "bufio"
    "fmt"
    "github.com/deanveloper/karman"
    "github.com/deanveloper/xkcdnews"
    "log"
    "os"
    "strings"
    "time"
)

func main() {
    bots := map[string]Bot{
        "karman": karman.New(log.New(os.Stdout, "[Karman]", log.Ldate|log.Ltime)),
        "xkcdnews": xkcdnews.New(log.New(os.Stdout, "[XKCDNews]", log.Ldate|log.Ltime)),
    }

    // start all bots
    for _, bot := range bots {
        go bot.Start()
    }

    // stop all bots when we stop
    defer func() {
        fmt.Println("Stopping bots gracefully...")

        done := make(chan string)
        timeout := time.After(5 * time.Second)

        for key, bot := range bots {
            go func(key string, bot Bot) {
                bot.Stop()
                done <- key
            }(key, bot)
        }

        for i := 0; i < len(bots); i++ {
            select {
            case s := <-done:
                bots[s] = nil
            case <-timeout:
                for key := range bots {
                    fmt.Println("Timed out:", key)
                }
            }
        }
    }()

    // keep bot running until "stop" is typed
    scan := bufio.NewScanner(os.Stdin)
    for scan.Scan() {
        input := scan.Text()
        if input == "stop" {
            break
        }
        split := strings.Split(input, " ")
        bot := bots[split[0]]
        if bot != nil {
            cmdr, ok := bot.(Commander)
            if ok {
                cmdr.Command(split[1:])
            } else {
                fmt.Println("That bot does not take commands")
            }
        } else {
            fmt.Println("Command not found:", split[0])
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
