package main

import (
	"fmt"
	"log"
	"strings"

	"time"
	ge "github.com/katsew/go-getenv"
	"github.com/bwmarrin/discordgo"
	"net/http"
	"github.com/katsew/kawaii-bot/discord/pkg/resp"
	"encoding/json"
	"math/rand"
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
	discord.LogLevel = discordgo.LogDebug
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

		h := ge.GetEnv("TARGET_API_HOST", "heartcatch")
		p := ge.GetEnv("TARGET_API_PORT", "5000")
		req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%s/images", h, p), nil)
		if err != nil {
			sendMessage(s, c, "Sorry...Could not complete your request... ;(")
			return
		}
		q := req.URL.Query()
		q.Add("q", msg)
		req.URL.RawQuery = q.Encode()

		client := http.Client{ Timeout: 10 * time.Second }
		res, err := client.Get(req.URL.String())
		resJson := new(resp.GoogleAPIResponse)
		err = json.NewDecoder(res.Body).Decode(resJson)
		if err != nil {
			sendMessage(s, c, "Sorry...Could not complete your request... ;(")
			return
		}
		if len(resJson.Items) > 0 {
			count := len(resJson.Items) - 1
			rand.Seed(time.Now().UnixNano())
			idx := rand.Intn(count)
			item := resJson.Items[idx]
			sendMessage(s, c, item.Link)
		} else {
			sendMessage(s, c, "Sorry...Could not complete your request... ;(")
		}
	}
}

func sendMessage(s *discordgo.Session, c *discordgo.Channel, msg string) {
	_, err := s.ChannelMessageSend(c.ID, msg)

	log.Println(">>> " + msg)
	if err != nil {
		log.Println("Error sending message: ", err)
	}
}
