package main

import (
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/rylio/ytdl"
)

func youtubeDown(url string) {
	vid, _ := ytdl.GetVideoInfo(url)
	file, _ := os.Create(vid.ID + ".mp3")

	defer file.Close()

	vid.Download(vid.Formats[17], file)

}

func getRightLink(link string) string {
	correctLink := strings.Split(link, " ")
	for i := 0; i < len(correctLink); i++ {
		if strings.Contains(correctLink[i], "https") {
			return correctLink[i]
		}

	}
	return "fail"
}

func getVideoID(url string) (string, bool) {
	if strings.Contains(url, "=") {
		youtubeDown(url)
		url := strings.Split(url, "=")
		return url[len(url)-1], true
	} else if strings.Contains(url, "/") {

		youtubeDown(url)
		url := strings.Split(url, "/")
		return url[len(url)-1], true

	} else {
		return "fail", false
	}
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
		{
			if update.Message == nil {
				continue
			}

			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

			if strings.Contains(update.Message.Text, "youtu.be") || strings.Contains(update.Message.Text, "youtube") {
				msg.Text = "music on the way"
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)
				id, allow := getVideoID(getRightLink(update.Message.Text))

				if allow {
					newVideo := tgbotapi.NewAudioUpload(update.Message.Chat.ID, id+".mp3")
					if _, err := bot.Send(newVideo); err != nil {
						log.Panic(err)
					}
				}
			} else {
				msg.Text = "paste only the link to music"
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)
			}
		}
	}
}
