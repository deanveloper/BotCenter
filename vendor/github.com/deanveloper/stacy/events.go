package stacy

import (
    "fmt"
    "github.com/bwmarrin/discordgo"
    "strings"
)

func (b *Stacy) ready(s *discordgo.Session, ev *discordgo.Ready) {
    err := s.UpdateStatus(0, "Statistic Simulator")
    if err != nil {
        b.log.Println("Error while readying:", err)
    } else {
        b.log.Println("I'm ready to check some statistics!")
    }
}

func (b *Stacy) handleCommand(s *discordgo.Session, ev *discordgo.MessageCreate) {
    if !strings.HasPrefix(ev.Message.Content, "!stats") {
        return
    }

    if ev.Message.Content == "!stats list" {
        s.ChannelMessageSend(ev.ChannelID, "**Games list:**")
        msg := ""
        for game := range Games {
            msg += game + ", "
        }
        msg = msg[0:len(msg) - 2] // remove trailing comma
        s.ChannelMessageSend(ev.ChannelID, "Games list:")
        return
    }

    if len(ev.Mentions) > 1 {
        s.ChannelMessageSend(ev.ChannelID, "Error: You can only get the stats of 1 person at a time.")
        return
    }

    args := strings.Split(ev.Message.Content, " ")[1:]
    if len(args) < 2 {
        b.sendUsage(ev.ChannelID, ev.Author.ID)
        return
    }

    game := Games[args[0]]
    if game == 0 {
        s.ChannelMessageSend(ev.ChannelID, "Error: That game hasn't been added! Use `!stats list` to get a list.")
        return
    }

    err := b.runOnGame(game, ev.Mentions[0], args[2:])
    if err != nil {
        s.ChannelMessageSend(ev.ChannelID, "Error: " + err.Error())
        return
    }
}

func (b *Stacy) sendUsage(chanId, userId string) {
    _, err := b.dg.ChannelMessageSend(chanId, fmt.Sprintf("<!%s> Usage: `!stats <game> <@user> [-help]`", userId))
    if err != nil {
        b.log.Println("Error sending usage message:", err)
    }
}
