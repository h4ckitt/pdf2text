package bot

import (
	"github.com/h4ckitt/goTelegram"
	"log"
	"memo/app"
	"net/http"
	"os"
	"strings"
)

var bot goTelegram.Bot

func Initialize(token, port string) {

	var err error

	bot, err = goTelegram.NewBot(token)

	if err != nil {
		log.Fatalln(err)
	}

	bot.SetHandler(handler)

	if port == "" {
		port = "8080"
	}

	//log.Println("Starting Server On PORT: ", port, "......")

	//log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%s", port), http.HandlerFunc(bot.UpdateHandler)))
}

func Handle(w http.ResponseWriter, r *http.Request) {
	bot.UpdateHandler(w, r)
}

func handler(update goTelegram.Update) {
	switch update.Type {
	case "document":
		names := strings.Split(update.Message.File.FileName, ".")
		if strings.ToLower(names[len(names)-1]) != "pdf" {
			bot.SendMessage("Only PDF Files Are Supported At The Moment", update.Message.Chat)
			return
		}

		log.Println("downloading file")
		if err := bot.GetFile(update.Message.File.FileID, update.Message.File.FileName); err != nil {
			log.Println(err)
			bot.SendMessage("An error occurred while processing your request", update.Message.Chat)
			return
		}

		log.Println("downloaded file")

		defer os.Remove(update.Message.File.FileName)

		contents, err := os.ReadFile(update.Message.File.FileName)

		if err != nil {
			bot.SendMessage("An error occurred while processing your file, please try again later", update.Message.Chat)
			return
		}

		result, err := app.ProcessDoc(contents)

		if err != nil {
			bot.SendMessage("An error occurred while processing your file, please try again later", update.Message.Chat)
			return
		}

		bot.SendMessage(result, update.Message.Chat)
	}
}
