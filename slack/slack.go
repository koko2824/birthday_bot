package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func SendMessage(text string) error {
	message := map[string]string{
		"text": text,
	}

	by, err := json.Marshal(message)
	req, err := http.NewRequest("POST", os.Getenv("SLACK_CH_WEBHOOK"), bytes.NewBuffer(by))
	if err != nil {
		return err
	}

	client := new(http.Client)
	_ , err = client.Do(req)
	if err != nil {
		return err
	}

	fmt.Println("success!")
	return nil
}
