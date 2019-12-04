package getfilms

import (
	"encoding/json"
	"net/http"
	"time"
)

// Film represents a film
type Film struct {
	Nome   string `json:"nome"`
	Orari  string `json:"orari"`
	Durata string `json:"durata"`
}

var myClient = &http.Client{Timeout: 120 * time.Second}

// GetList call the API server and return an Array of
// Film structs
func GetList(url string, target *[]Film) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
