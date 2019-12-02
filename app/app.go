package app

import (
	"log"
	"os"

	"github.com/M3nthis/the-space-cinema-bot/getfilms"
	tb "gopkg.in/tucnak/telebot.v2"
)

//StartApp starts the app
func StartApp() {
	var (
		port      = os.Getenv("PORT")
		publicURL = os.Getenv("PUBLIC_URL")
		fetchURL  = os.Getenv("FETCH_URL")
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
		b.Send(m.Chat, "Ciao!")
	})

	b.Handle("/lista_film", func(m *tb.Message) {
		b.Send(m.Chat, "Carico l'elenco dei film...")
		b.Send(m.Chat, loadFilms(fetchURL))
	})

	b.Start()
}

func loadFilms(url string) string {
	films := []getfilms.Film{}
	err := getfilms.GetList(url, &films)
	if err != nil {
		log.Println(err)
		return "Errore nel caricamento"
	}

	resp := ""
	for _, film := range films {
		resp += (film.Nome + "\n\n")
	}
	return resp
}
