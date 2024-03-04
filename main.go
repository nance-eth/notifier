package main

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

const nanceEndpoint = "https://api.nance.app/"

var s *discordgo.Session

type config struct {
	Space   string `json:"space"`
	GuildId string `json:"guildId"`
}

func main() {
	// Read env variables and create discord session
	_, err := os.Stat(".env")
	if !os.IsNotExist(err) {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error loading .env file: %v\n", err)
		}
	}

	if discordToken := os.Getenv("DISCORD_TOKEN"); discordToken == "" {
		log.Fatalf("DISCORD_TOKEN is not set in environment or .env")
	}

	s, err = discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Fatalf("Error creating Discord session: %v\n", err)
	}
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	// Add slash command handlers
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	err = s.Open()
	if err != nil {
		log.Fatalf("Error opening Discord session: %v\n", err)
	}
	defer s.Close()

	// Read config file
	f, err := os.Open("config.json")
	if err != nil {
		log.Fatalf("Error opening config file: %v\n", err)
	}

	var config []config
	json.NewDecoder(f).Decode(&config)

	// Add spaces for each item in the config file
	for _, conf := range config {
		go addSpace(conf.Space, conf.GuildId)
	}

	// Block until the user interrupts the program
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt, syscall.SIGHUP)
	<-sc

	log.Println("Shutting down...")
}

// TODO: WIP
func addSpace(space string, guildId string) {
	spaceData, err := nanceSpace(space)
	if err != nil {
		log.Printf("Error fetching space data for %s (skipping guild): %v\n", space, err)
		return
	}

	// Register slash commands
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, command := range commands {
		reg, err := s.ApplicationCommandCreate(s.State.User.ID, guildId, command)
		if err != nil {
			log.Printf("Error registering command %s in guild %s (skipping guild): %v\n", command.Name, guildId, err)
			return
		}
		registeredCommands[i] = reg
	}

	// Defer command deletion
	defer func() {
		for _, v := range registeredCommands {
			err := s.ApplicationCommandDelete(s.State.User.ID, v.GuildID, v.ID)
			if err != nil {
				log.Panicf("Error deleting command '%s' from guild %s: %v\n", v.Name, guildId, err)
			}
		}
	}()
}
