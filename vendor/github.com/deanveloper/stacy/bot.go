package stacy

import (
    "github.com/bwmarrin/discordgo"
    "io/ioutil"
    "log"
)

type Stacy struct {
    dg  *discordgo.Session
    log *log.Logger
}

func New(logger *log.Logger) *Stacy {
    return &Stacy{log: logger}
}

func (b *Stacy) Start() {
    bytes, err := ioutil.ReadFile("STACY_SECRET")
    if err != nil {
        b.log.Println("Error reading secret key:", err)
        return
    }

    dg, err := discordgo.New("Bot " + string(bytes))
    if err != nil {
        b.log.Println("Error creating discord session:", err)
        return
    }

    b.dg = dg

    dg.AddHandler(b.ready)
}

func (b *Stacy) Stop() {
    b.dg.Close()
}
