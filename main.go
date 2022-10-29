package main

import (
	"fmt"
	"log"
	"memo/bot"
	"net/http"
	"os"
	"strings"
)

var (
	// ocrClient *gosseract.Client
	SECRET = os.Getenv("TG_SECRET")
)

func main() {
	/*client := gosseract.NewClient()

	defer client.Close()

	client.SetImage("file.jpg")
	client.Languages = []string{"eng"}
	fmt.Println(client.Version())

	text, err := client.Text()

	fmt.Println(text, " ", err)*/

	PORT := os.Getenv("PORT")

	bot.Initialize(os.Getenv("BOT_API_TOKEN"), PORT)

	fmt.Println("Starting Server On Port ", PORT, " .......")

	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%s", PORT), http.HandlerFunc(router)))
}

func router(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.Header.Get("X-Telegram-Bot-Api-Secret-Token"), SECRET) {
		bot.Handle(w, r)
	}
}
