package main

import (
	"DiscordPancakeAuto/bot"
	"errors"
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
	"time"
)

var (
	r    = regexp.MustCompile("(\\d+|\\d{1,3}(,\\d{3})*)(\\.\\d+)?")
	remo = regexp.MustCompile("(<a?)?:\\w+:(\\d{18}>)?")
	rmen = regexp.MustCompile("<@!*&*[0-9]+>")
)

func main() {
	bot.InitBot()
	c := cron.New()
	c.AddFunc("@every 5m6s", doWork)
	c.AddFunc("@every 1m4s", doFishing)
	c.AddFunc("@every 8m2s", sellItems)
	c.AddFunc("@every 10m1s", doTrivia)
	c.Start()

	log.Println("Bot is now running.", bot.Client.State.User.Username, "Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func doWork() {
	log.Println("Starting to invoke .work command...")
	message, err := bot.Client.ChannelMessageSend(bot.ChannelId, ".work")
	if err != nil {
		log.Println("Something went wrong while doing the .work command", err)
		return
	}
	time.Sleep(1 * time.Second)
	messages, err := bot.Client.ChannelMessages(bot.ChannelId, 5, "", message.ID, "")
	if err != nil {
		log.Println("Error while getting messages after .work command", err)
		return
	}
	for _, v := range messages {
		if v.Author.ID != bot.PancakeId {
			continue
		}
		if strings.Contains(v.Content, bot.Client.State.User.ID) {
			if strings.Contains(v.Content, "cooked") {
				bot.Client.ChannelTyping(bot.ChannelId)
				time.Sleep(2 * time.Second)
				bot.Client.ChannelMessageSend(bot.ChannelId, ".dep all")
				log.Println("Successfully worked hard")
				break
			} else if strings.Contains(v.Content, "burnt") {
				log.Println("Burnt pancakes :(")
				replaced := rmen.ReplaceAllString(remo.ReplaceAllString(v.Content, ""), "")
				r := strings.Join(r.FindAllString(replaced, -1), "")
				bot.Client.ChannelMessageSend(bot.ChannelId, ".wd "+r)
				break
			}
		}
	}
	log.Println("Finishing to invoke .work command...")
}

func doFishing() {
	log.Println("Starting fishing...")
	message, err := bot.Client.ChannelMessageSend(bot.ChannelId, ".fish")
	if err != nil {
		log.Println("Something went wrong while doing the .fish command", err)
		return
	}
	time.Sleep(1 * time.Second)
	messages, err := bot.Client.ChannelMessages(bot.ChannelId, 5, "", message.ID, "")
	if err != nil {
		log.Println("Something went wrong while fishing", err)
		return
	}
	for _, v := range messages {
		if v.Author.ID != bot.PancakeId {
			continue
		}
		if strings.Contains(v.Content, "Your fishing rod broke!") {
			log.Println("Buying a fishing rod...")
			err = buyFishingRod()
			if err != nil {
				log.Println("Something went wrong while fishing", err)
				return
			}
			break
		} else if strings.Contains(v.Content, "You don't have a fishing rod") {
			log.Println("Buying a fishing rod...")
			err = buyFishingRod()
			if err != nil {
				log.Println("Something went wrong while fishing", err)
				return
			}
			doFishing()
			return
		}
	}
	log.Println("Stopped fishing...")
}
func doTrivia() {
	log.Println("Starting doing trivia...")
	message, err := bot.Client.ChannelMessageSend(bot.ChannelId, ".trivia hard")
	if err != nil {
		log.Println("Something went wrong while doing the .trivia hard command", err)
		return
	}
	time.Sleep(1 * time.Second)
	messages, err := bot.Client.ChannelMessages(bot.ChannelId, 5, "", message.ID, "")
	if err != nil {
		log.Println("Something went wrong while doing the .trivia hard command", err)
		return
	}

	for _, v := range messages {
		if v.Author.ID != bot.PancakeId {
			continue
		}
		if len(v.Embeds) > 0 {
			embed := v.Embeds[0]
			if strings.Contains(embed.Author.Name, bot.Client.State.User.Username) && strings.Contains(embed.Description, "[4]") {
				err = solveTrivia(1, 4)
				if err != nil {
					log.Println("Something went wrong while doing trivia", err)
				}
				break
			} else {
				err = solveTrivia(1, 2)
				if err != nil {
					log.Println("Something went wrong while doing trivia", err)
				}
				break
			}
		}

	}
	log.Println("Finished doing trivia...")
}
func solveTrivia(min int, max int) error {
	message, err := bot.Client.ChannelMessageSend(bot.ChannelId, fmt.Sprint(getRandom(min, max)))
	if err != nil {
		return err
	}
	time.Sleep(1 * time.Second)
	messages, err := bot.Client.ChannelMessages(bot.ChannelId, 5, "", message.ID, "")

	for _, v := range messages {
		if v.Author.ID != bot.PancakeId {
			continue
		}
		if strings.Contains(v.Content, "Correct!") {
			log.Println("Successfully solved trivia")
			bot.Client.ChannelMessageSend(bot.ChannelId, ".dep all")
			return nil
		} else {
			log.Println("Entered incorrect answer while doing trivia")
			return nil
		}
	}
	return errors.New("pancake did not give any response after doing trivia")
}

func getRandom(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func buyFishingRod() error {
	log.Println("Buying a fishing rod...")
	bot.Client.ChannelTyping(bot.ChannelId)
	time.Sleep(1 * time.Second)
	bot.Client.ChannelMessageSend(bot.ChannelId, ".wd 50")
	time.Sleep(1 * time.Second)
	message, err := bot.Client.ChannelMessageSend(bot.ChannelId, ".buy rod")
	if err != nil {
		return err
	}
	time.Sleep(1 * time.Second)
	messages, err := bot.Client.ChannelMessages(bot.ChannelId, 5, "", message.ID, "")
	for _, v := range messages {
		if v.Author.ID != bot.PancakeId {
			continue
		}
		if strings.Contains(v.Content, "Are you sure you want to purchase") {
			bot.Client.MessageReactionAdd(bot.ChannelId, v.ID, "check:502235451592278016")
			log.Println("Finishing a fishing rod purchase...")
			return nil
		}
	}
	return errors.New("pancake didn't give any response while buying a fishing rod")

}

func sellItems() {
	log.Println("Starting selling items...")
	message, err := bot.Client.ChannelMessageSend(bot.ChannelId, ".inv")
	if err != nil {
		log.Println("Something went wrong while checking the inventory", err)
		return
	}
	time.Sleep(1 * time.Second)
	messages, err := bot.Client.ChannelMessages(bot.ChannelId, 5, "", message.ID, "")
	if err != nil {
		log.Println("Something went wrong while trying to sell items", err)
		return
	}
	for _, v := range messages {
		if v.Author.ID != bot.PancakeId {
			continue
		}
		if len(v.Embeds) > 0 {
			embed := v.Embeds[0]
			if len(embed.Fields) > 5 {
				err = confirmSell()
				if err != nil {
					log.Println("Something went wrong while selling", err)
					return
				}
			}
		}
	}
	log.Println("Finished selling items...")
}

func confirmSell() error {
	bot.Client.ChannelTyping(bot.ChannelId)
	time.Sleep(1 * time.Second)
	message, err := bot.Client.ChannelMessageSend(bot.ChannelId, ".sell all")
	if err != nil {
		return err
	}
	time.Sleep(1 * time.Second)
	messages, err := bot.Client.ChannelMessages(bot.ChannelId, 5, "", message.ID, "")
	if err != nil {
		return err
	}
	for _, v := range messages {
		if v.Author.ID != bot.PancakeId {
			continue
		}
		if strings.Contains(v.Content, "Are you sure you want to sell") {
			bot.Client.MessageReactionAdd(bot.ChannelId, v.ID, "check:502235451592278016")
			bot.Client.ChannelTyping(bot.ChannelId)
			time.Sleep(1 * time.Second)
			bot.Client.ChannelMessageSend(bot.ChannelId, ".deposit all")
			return nil
		}
	}
	return errors.New("pancake didn't give any response while confirming sell")
}
