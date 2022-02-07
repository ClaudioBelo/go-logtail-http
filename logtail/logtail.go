package logtail

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type LogTail struct {
	SourceToken string
}

var ApiURL = "https://in.logtail.com"

func (t LogTail) New(token string) *LogTail {
	t.SourceToken = token
	return &t
}

func (t *LogTail) sendMessage(msg string) error {
	body := struct {
		Message string `json:"message"`
	}{
		Message: msg,
	}

	json_bytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", ApiURL, bytes.NewBuffer(json_bytes))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", t.SourceToken))

	client := &http.Client{}

	if resp, err := client.Do(req); err != nil {
		return err
	} else {
		defer resp.Body.Close()
		return nil
	}
}
