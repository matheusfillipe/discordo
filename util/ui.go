package util

import (
	"regexp"
	"strings"

	"github.com/ayntgl/discordgo"
)

var (
	boldRegex          = regexp.MustCompile(`(?m)\*\*(.*?)\*\*`)
	italicRegex        = regexp.MustCompile(`(?m)\*(.*?)\*`)
	underlineRegex     = regexp.MustCompile(`(?m)__(.*?)__`)
	strikeThroughRegex = regexp.MustCompile(`(?m)~~(.*?)~~`)
)

func ParseMarkdown(md string) string {
	var res string
	res = boldRegex.ReplaceAllString(md, "[::b]$1[::-]")
	res = italicRegex.ReplaceAllString(res, "[::i]$1[::-]")
	res = underlineRegex.ReplaceAllString(res, "[::u]$1[::-]")
	res = strikeThroughRegex.ReplaceAllString(res, "[::s]$1[::-]")

	return res
}

func ChannelToString(c *discordgo.Channel) string {
	var repr string
	switch c.Type {
	case discordgo.ChannelTypeGuildText, discordgo.ChannelTypeGuildNews:
		repr = "#" + c.Name
	case discordgo.ChannelTypeGuildVoice:
		repr = "🔊" + c.Name
	case discordgo.ChannelTypeDM, discordgo.ChannelTypeGroupDM:
		if len(c.Recipients) == 1 {
			rp := c.Recipients[0]
			repr = rp.Username + "#" + rp.Discriminator
		} else {
			rps := make([]string, len(c.Recipients))
			for i, r := range c.Recipients {
				rps[i] = r.Username + "#" + r.Discriminator
			}

			repr = strings.Join(rps, ", ")
		}
	}

	return repr
}

func HasKeybinding(ks []string, k string) bool {
	for _, repr := range ks {
		if repr == k {
			return true
		}
	}

	return false
}
