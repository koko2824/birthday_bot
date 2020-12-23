package spreadsheet

import (
  "encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

type SheetData []struct {
	Name     string `json:"name"`
	Slack    string `json:"slack"`
	Birthday string `json:"birthday"`
}

func GetData() (SheetData, error) {
	req, err := http.NewRequest("GET", os.Getenv("SPREADSHEET_URL"), nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := new(SheetData)
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return *data, nil
}
