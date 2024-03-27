package main

import (
	"fmt"
	"log"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("TOKEN", "token of the bot")
	os.Setenv("CHANNEL_ID", "C04J7QW95U4")
	api := slack.New(os.Getenv("TOKEN"))
	channeArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"book.pdf"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channeArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("file uploaded %v", file.Name, file.URL)
	}
}
