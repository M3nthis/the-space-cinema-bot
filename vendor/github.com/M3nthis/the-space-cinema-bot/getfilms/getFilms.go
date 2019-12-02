package getfilms

import (
	"encoding/json"
	"net/http"
	"time"
)

type Film struct {
	Nome string `json:"nome"`
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func GetList(url string, target *[]Film) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
