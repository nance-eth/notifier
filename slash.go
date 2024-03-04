package main

import "github.com/bwmarrin/discordgo"

var commands = []*discordgo.ApplicationCommand{
	{
		Name:        "subscribe",
		Description: "Subscribe to voting alerts",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "address",
				Description: "The Ethereum address or ENS name to receive voting alerts for",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionBoolean,
				Name:        "dm-alerts",
				Description: "Receive voting alerts via DM",
				Required:    false,
			},
		},
	},
	{
		Name:        "add-address",
		Description: "Add an address to receive voting alerts for",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "address",
				Description: "The Ethereum address or ENS name to receive voting alerts for",
				Required:    true,
			},
		},
	},
	{
		Name:        "unsubscribe",
		Description: "Unsubscribe from all voting alerts",
	},
}

var commmandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"subscribe":   func(s *discordgo.Session, i *discordgo.InteractionCreate) {},
	"add-address": func(s *discordgo.Session, i *discordgo.InteractionCreate) {},
	"unsubscribe": func(s *discordgo.Session, i *discordgo.InteractionCreate) {},
}
