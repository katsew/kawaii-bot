package main

import (
	"fmt"
	"log"
	"strings"

	"time"
	ge "github.com/katsew/go-getenv"
	"github.com/bwmarrin/discordgo"
)

var (
	Token string
	BotName string
	TargetChannelId string
	stopBot         = make(chan bool)
)

func init() {
	Token = fmt.Sprintf("Bot %s", ge.GetEnv("BOT_TOKEN", "").String())
	BotName = fmt.Sprintf("<@%s>", ge.GetEnv("BOT_ID", "").String())
	TargetChannelId = ge.GetEnv("TARGET_CHANNEL_ID", "").String()
}

func main() {

	discord, err := discordgo.New()
	discord.Token = Token
	if err != nil {
		fmt.Println("Error logging in")
		fmt.Println(err)
	}

	discord.AddHandler(onMessageCreate)

	err = discord.Open()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Listening...")
	<-stopBot
	return
}

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		log.Println("Error getting channel: ", err)
		return
	}
	if c.ID != TargetChannelId {
		log.Printf("This is not a target channel: %s", c.Name)
		return
	}
	fmt.Printf("%20s %20s %20s > %s\n", m.ChannelID, time.Now().Format(time.Stamp), m.Author.Username, m.Content)

	switch {
	case strings.HasPrefix(m.Content, BotName):
		msg := strings.TrimPrefix(m.Content, BotName)
		sendMessage(s, c, "OK, wait a minutes!")
		log.Printf("Your channelID: %s", c.ID)
		time.Sleep(3 * time.Second)
		sendMessage(s, c, msg)
	}
}

func sendMessage(s *discordgo.Session, c *discordgo.Channel, msg string) {
	_, err := s.ChannelMessageSend(c.ID, msg)

	log.Println(">>> " + msg)
	if err != nil {
		log.Println("Error sending message: ", err)
	}
}
