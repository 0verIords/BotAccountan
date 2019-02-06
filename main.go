package main

import (
	"log"
	"os"
	"strings"

	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/rylio/ytdl"
)

func youtubeDown(url string) {
	vid, _ := ytdl.GetVideoInfo(url)
	file, _ := os.Create(vid.ID + ".mp3")

	defer file.Close()
	vid.Download(vid.Formats[13], file)
}

func getVideoID(url string) (string, bool) {
	if strings.Contains(url, "=") {
		urls := strings.Split(url, "=")
		if len(urls[len(urls)-1]) == 11 {
			return urls[len(urls)-1], true
		}
	}
	if strings.Contains(url, "/") {
		urls := strings.Split(url, "/")
		if len(urls[len(urls)-1]) == 11 {
			return urls[len(urls)-1], true
		}

	}
	return "fail", false

}

func main() {

	bot, err := tgbotapi.NewBotAPI(os.Getenv("token"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		link := update.Message.Text
		user := update.Message.Chat.ID
		finish := make(chan bool)

		go func() {
			log.Printf("[%s] %s", update.Message.From.UserName, link)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

			if strings.Contains(link, "youtu.be") || strings.Contains(link, "https://www.youtube.com/watch?v") {
				go func() {
					for {
						select {
						case <-finish:
							close(finish)
							return
						default:
							msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
							time.Sleep(time.Second * 30)
							msg.Text = "file is download"
							bot.Send(msg)
						}
					}
				}()
				msg.Text = "music on the way"
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)
				id, allow := getVideoID(link)

				if allow {
					youtubeDown(link)
					newVideo := tgbotapi.NewAudioUpload(user, id+".mp3")
					if _, err := bot.Send(newVideo); err != nil {
						log.Panic(err)
					}
					finish <- true
				} else {
					msg.Text = "incorrect link"
					bot.Send(msg)
					finish <- true
				}
			} else {
				msg.Text = "paste only the link to music"
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)
				finish <- true
			}
		}()
	}
}
