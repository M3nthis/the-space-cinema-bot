package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	tb "gopkg.in/tucnak/telebot.v2"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	var (
		port      = os.Getenv("PORT")
		publicURL = os.Getenv("PUBLIC_URL")
		token     = os.Getenv("TOKEN")
	)

	webhook := &tb.Webhook{
		Listen:   ":" + port,
		Endpoint: &tb.WebhookEndpoint{PublicURL: publicURL},
	}

	pref := tb.Settings{
		Token:  token,
		Poller: webhook,
	}

	b, err := tb.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Sender, "Ciao!")
	})

	inlineBtn1 := tb.InlineButton{
		Unique: "filmList",
		Text:   "Guarda elenco dei film di oggi",
	}

	// Button
	b.Handle(&inlineBtn1, func(c *tb.Callback) {
		b.Respond(c, &tb.CallbackResponse{
			ShowAlert: false,
		})
		b.Send(c.Sender, "Elenco film")
	})

	inlineKeys := [][]tb.InlineButton{
		[]tb.InlineButton{inlineBtn1},
	}

	b.Handle("/lista_film", func(m *tb.Message) {
		b.Send(
			m.Sender,
			"Scegli l'azione",
			&tb.ReplyMarkup{InlineKeyboard: inlineKeys},
		)
	})

	b.Start()
}
