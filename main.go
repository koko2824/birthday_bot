package main

import (
	"birthday_bot/spreadsheet"
	"fmt"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		return
	}

	sheetData, err:= spreadsheet.GetData()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(sheetData)
}
