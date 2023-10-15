// bot file
// written by cyb3rguru

package bot

import (
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/sashabaranov/go-openai"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strings"
)

// Initialize tokens

var (
	OPENAPI_TOKEN string
	BOT_TOKEN     string
)

func chatgpt(message string) string {

	c := openai.NewClient(OPENAPI_TOKEN)

	resp, err := c.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    "user",
					Content: message,
				},
			},
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	return resp.Choices[0].Message.Content
}

func Run() {
	// create new Discord client session

	discord, err := discordgo.New("Bot " + BOT_TOKEN)
	if err != nil {
		log.Fatal(err)
	}

	// Add event handler for general message

	discord.AddHandler(newMessage)

	// Open Session
	discord.Open()
	defer discord.Close()

	log.Println("Bot running...")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

}

func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {

	// Check if the message content contains a mention of a bot.

	botMentionID := discord.State.User.ID

	if strings.Contains(message.Content, botMentionID) {
		// Split the message content into parts.
		response := fmt.Sprintf("Thanks for mentioning me, %s! use `!help` to see available commands",
			message.Author.Mention())
		discord.ChannelMessageSend(message.ChannelID, response)
	}

	// ignore messages from bot to prevent feedback loops.
	if message.Author.ID == discord.State.User.ID {
		return
	}

	// same as above, safe check
	if message.Author.Bot {
		return
	}

	// check if the message was sent in a direct message (DM).
	if message.GuildID == "" && !strings.Contains(message.Content, "!"){
		response := fmt.Sprintf("Hello there!,\nUse `!help` for more information")
		discord.ChannelMessageSend(message.ChannelID, response)
	}

	// Check if the message content is a help command.
	if strings.HasPrefix(message.Content, "!help") {
		// Send a message with information about available commands.
		helpMessage := "Here are the available commands:\n" +
			"1. `!hello`: Say hello to the bot.\n" +
			"2. `!say <message>`: Make the bot say something.\n" +
			"3. `!info`: Get information about the bot.\n" +
			"4. `!gpt <query>` : Ask a question and get an answer with chatgpt.\n" +
			"5. `!help`: Display this help message."

		discord.ChannelMessageSend(message.ChannelID, helpMessage)
	}

	switch {
	case strings.Contains(message.Content, "!hello"):
		greet_response := []string{
			"Hello! How can I assist you today?",
			"Hey there! What can I do for you?",
			"Greetings! What's on your mind?",
			"Hi! I'm here to help. What do you need?",
		}

		response := fmt.Sprintf("%s"+greet_response[rand.Intn(len(greet_response))],
			message.Author.Mention())

		discord.ChannelMessageSend(message.ChannelID, response)

	case strings.Contains(message.Content, "!say"):
		// Split the message content into parts.
		parts := strings.Fields(message.Content)

		// join the rest of the message
		request := strings.Join(parts[1:], " ")

		discord.ChannelMessageSend(message.ChannelID, request)

	case strings.Contains(message.Content, "!info"):
		response := "Welcome to our Discord bot! 🤖 " +
			"This bot is your AI companion for asking questions and generating answers." +
			"Whether you need assistance with programming, seeking answers to your queries," +
			"or want to quickly generate code snippets, our bot is here to help you. " +
			"Feel free to ask questions, request code samples, or engage in insightful conversations" +
			"with our AI-powered assistant. Just type your questions, and let's get started on a journey of discovery and code creation!"

		discord.ChannelMessageSend(message.ChannelID, response)

	case strings.Contains(message.Content, "!gpt"):
		// Split the message content into parts.
		parts := strings.Fields(message.Content)

		// join the rest of the message
		request := strings.Join(parts[1:], " ")

		// send users request to chatgpt
		response := chatgpt(request)

		discord.ChannelMessageSend(message.ChannelID, response)

	}

}
