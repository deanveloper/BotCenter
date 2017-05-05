package stacy

import "github.com/bwmarrin/discordgo"

// enum of games
const (
    _         = iota    // assign iota 0 to _ so that overwatch = 1
    OVERWATCH
)

var Games = map[string]int{
    "overwatch": OVERWATCH,
    "ow":        OVERWATCH,
}

func (b *Stacy) runOnGame(game int, user *discordgo.User, extra []string) error {

    switch game {
    case 1:
        resp, err := b.dg.Request("GET", discordgo.EndpointUserConnections(user.ID), struct{}{})
        b.log.Println(string(resp))
        return err
    }

    return nil
}
