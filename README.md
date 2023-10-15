# AskMe Discord Bot

**Author:** cyb3rguru

## Description

This Discord bot is a versatile AI-powered assistant that can help you with various tasks, from generating code to answering questions. It integrates with ChatGPT to provide intelligent responses. Whether you're a developer looking for code snippets or need answers to general queries, this bot has got you covered.

## Features

- **Hello Command**: Greet the bot and start a conversation.
- **Say Command**: Make the bot say something specific.
- **Info Command**: Learn more about the bot and its capabilities.
- **GPT Command**: Ask questions and get answers with ChatGPT.
- **Help Command**: Get a list of available commands and their descriptions.

## Prerequisites

Before you run the bot, you'll need:

- Go installed on your system.
- Discord Bot Token for authentication.
- OpenAI GPT-3 API Token for ChatGPT integration.

## Setup

1. Clone this repository.
2. Create a `.env` file with the following environment variables:

```plaintext
OPENAI_TOKEN=your_openai_token_here
BOT_TOKEN=your_discord_bot_token_here
```

3. Run the bot using the following command:

```bash
go run bot.go
```

## Deploying with docker

- Use the following command to deploy:

```bash
docker build -t askme-bot .
docker run askme-bot
```

## Usage

1. Invite the bot to your Discord server using this [link](https://discord.com/oauth2/authorize?client_id=1162986767872430080&permissions=8&scope=bot).
2. Use the available commands to interact with the bot.
3. Enjoy the bot's assistance in generating code, answering questions, and more.

## Examples

- Direct Message
  ![](https://imgur.com/c8NsKGi.png)
  ![](https://imgur.com/JTPDp5L.png)
  ![](https://imgur.com/KNTiocf.png)
- Group Message
  ![](https://imgur.com/wWiQqn8.png)

## Commands

- `!hello`: Greet the bot and start a conversation.
- `!say <message>`: Make the bot say something specific.
- `!info`: Learn more about the bot and its capabilities.
- `!gpt <query>`: Ask questions and get answers with ChatGPT.
- `!help`: Get a list of available commands and their descriptions.

## Contributing

Feel free to contribute to this project by creating pull requests, reporting issues, or suggesting improvements. Your feedback is valuable.

## License

This project is licensed under the [MIT License](LICENSE).

```
This README provides information about your Discord bot, its features, how to set it up, and how to use its commands. You can further customize it to include specific installation instructions, configuration details, and additional features.
```
