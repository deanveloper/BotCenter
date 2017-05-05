package stacy

import "net/url"

func (b *Stacy) getOverwatch(battletag, platform string) {
    link, err := url.Parse("https://owapi.net/")
    if err != nil {
        // should never happen, but just in case
        b.log.Println("Error parsing url:", err)
        return
    }

    link.RawQuery = "platform=" + platform
}