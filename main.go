package main

import (
	"birthday_bot/slack"
	"birthday_bot/spreadsheet"
	"fmt"
	"github.com/joho/godotenv"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		return
	}

	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		fmt.Println(err)
		return
	}
	nowTime := time.Now().In(loc)
	nowTime = time.Date(2020, 6, 10, 0, 0, 0, 0, loc)
	_, m, d := nowTime.Date()
	sheetData, err := spreadsheet.GetData()

	var message string
	for _, v := range sheetData {
		fmt.Println(v.Birthday)
		t, err := time.Parse("2006-01-02 15:04:05 -0700", v.Birthday)
		if err != nil {
			fmt.Println(err)
			return
		}

		if t.Month() == m && t.Day() == d {
			message += fmt.Sprintf("<@%s>さん", v.Slack)
		}
	}

	if message != "" {
		res := fmt.Sprintf("今日は%sの誕生日です！\nみんなでお祝いしましょう！", message)
		fmt.Println(res)
		slack.SendMessage(res)
		return
	}

	fmt.Println("今日誕生日の人はいません。")
}
