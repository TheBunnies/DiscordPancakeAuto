package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
)

const (
	PancakeId = "239631525350604801"
)

var (
	ChannelId = ""
	Token     = ""
	Client    = &discordgo.Session{}
)

func InitBot() {
	err := loadEnv()
	if err != nil {
		log.Fatalln(err)
	}
	client, err := discordgo.New(Token)
	if err != nil {
		log.Fatalln(err)
	}
	Client = client
	client.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36"
	client.Identify.Intents = discordgo.IntentsAll
	client.Identify.Properties = discordgo.IdentifyProperties{
		OS:     "Windows",
		Device: "Web",
	}
	err = client.Open()
	if err != nil {
		log.Fatalln(err)
	}

}

func loadEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}
	ChannelId = os.Getenv("CHANNEL_ID")
	Token = os.Getenv("TOKEN")
	return nil
}
