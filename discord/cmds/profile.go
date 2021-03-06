package cmds

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/mxrk/Emilia/database"
)

func GetXP(ds *discordgo.Session, dm *discordgo.Message, content []string) {
	user := dm.Mentions
	var ret string
	if len(user) == 0 {
		ret = fmt.Sprintf("You have %v xp.", database.ReturnXP(dm.Author.ID))
	} else {
		xp := database.ReturnXP(user[0].ID)
		if xp == "" {
			ret = fmt.Sprintf("%s doesn't have experience points.", user[0].Username)
		} else {
			ret = fmt.Sprintf("%s has %v xp", user[0].Username, xp)
		}

	}
	ds.ChannelMessageSend(dm.ChannelID, ret)
}

//GetLevel answers with the first mentioned user or the author
func GetLevel(ds *discordgo.Session, dm *discordgo.Message, content []string) {
	user := dm.Mentions
	var ret string
	if len(user) == 0 {
		ret = fmt.Sprintf("You're level %v", database.GetLevel(dm.Author.ID))
	} else {
		level := database.GetLevel(user[0].ID)
		if level == "" {
			ret = fmt.Sprintf("%s doesn't have a level yet.", user[0].Username)
		} else {
			ret = fmt.Sprintf("%s is level %v", user[0].Username, level)
		}

	}
	ds.ChannelMessageSend(dm.ChannelID, ret)
}

//GetCoins returns the level of the current user or first mentioned user
func GetCoins(ds *discordgo.Session, dm *discordgo.Message, content []string) {
	user := dm.Mentions
	var ret string
	if len(user) == 0 {
		ret = fmt.Sprintf("You have %v coins!", database.GetCoins(dm.Author.ID))
	} else {
		coins := database.GetCoins(user[0].ID)
		if coins == "" {
			ret = fmt.Sprintf("%s doesn't have a single coin yet.", user[0].Username)
		} else {
			ret = fmt.Sprintf("%s has %v coins", user[0].Username, coins)
		}

	}
	ds.ChannelMessageSend(dm.ChannelID, ret)
}
