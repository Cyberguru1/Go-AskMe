// main file
// written by cyb3rguru
// Date: 2023-8-22

package main

import (
	"Go-AskMe/botFile"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	// load up the env file

	if err := godotenv.Load(); err != nil {
		fmt.Println("Error reading .env file", err)
	}

	bot.OPENAPI_TOKEN = os.Getenv("OPENAPI_TOKEN")
	bot.BOT_TOKEN = os.Getenv("DISCORD_BOT_TOKEN")

	log.Println("Loaded Api keys")

	// Startup the bot
	bot.Run()
}
