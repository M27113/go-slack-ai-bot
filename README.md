# 🤖 SlackAI Assistant | Your Smart Workspace Bot
⚡ Boost Slack productivity with an AI bot that answers, summarizes, and plans your tasks.

## 📖 Project Overview

**SlackAI Assistant** is a **smart, context-aware AI bot** designed for Slack workspaces.  
It can interact with users via **channel mentions** and **direct messages**, providing:  

- 💬 Real-time responses to greetings, questions, and commands.  
- 📚 Summaries of articles, movies, anime plots, or any text input.  
- 🧠 Intelligent task planning and daily schedule suggestions.  
- 🎯 Motivational quotes and guidance for users.  
- 🔗 A modular architecture, making it easy to add new AI tools or integrations.  

The bot is built using **GoLang** for Slack integration, **OpenAI GPT** for AI reasoning, and **Slack Socket Mode** for real-time event handling.  
It is ideal for showcasing AI-powered productivity features and learning how to integrate AI into messaging platforms.


## 🛠 Tech Stack

- 🟢**GoLang** for Slack integration  
- 🟣**OpenAI GPT** for AI reasoning & responses  
- ⚡**Slack API (Socket Mode)** for real-time events

---

## 📂 Project Structure

    go-slack-ai-bot/
    ├── main.go                # Entry point: Slack event handling & server setup
    ├── handler/
    │   └── slack_handler.go   # Handles Slack events & slash commands
    ├── ai/
    │   └── gpt_client.go      # Communicates with OpenAI API
    ├── storage/
    │   └── logger.go          # Logs messages, responses, timestamps
    ├── config/
    │   └── config.go          # Stores API keys, tokens, and configs
    ├── go.mod                 # Golang module
    ├── demo/                  # Screenshots for README/demo
    └── README.md              # Project description, setup instructions


## 🏗 **Architecture**

    
    +---------------+    +----------------+    +-------------------+    +------------------+    +------------------+
    | User Input    |--->| Slack Event    |--->| AI Handler        |--->| OpenAI GPT       |--->| Slack Reply      |
    | (Channel/DM)  |    | (mention/DM)   |    | slack_handler.go  |    | ai/gpt_client.go |    | Channel / DM     |
    +---------------+    +----------------+    +-------------------+    +------------------+    +------------------+

### 📝 **Explanation**

- **User Input:** Messages from Slack users (channel mentions or DMs).  
- **Slack Event:** Captures events using Socket Mode (`app_mention` or `message.im`).  
- **AI Handler:** Processes Slack events, prepares input for GPT, and determines the response.  
- **OpenAI GPT API:** Generates context-aware responses using the GPT model.  
- **Slack Reply:** Sends the final response back to the channel or DM.  

---

## ⚙️ **Setup**

1. **Clone the repository**

    ```bash
    git clone https://github.com/M27113/go-slack-ai-bot.git
    cd go-slack-ai-bot

2. **Create a .env file in the root directory:**
    ```bash
    SLACK_BOT_TOKEN=xoxb-your-bot-token
   
    SLACK_APP_TOKEN=xapp-your-app-token
   
    OPENAI_API_KEY=sk-your-openai-key


4. **Install Go dependencies**
    ```bash
    go mod tidy


5. **Run the bot**
    ```bash
    go run main.go

**Note**: 
- ⚠️ Ensure your Slack app has the following Bot Token Scopes:
  - app_mentions:read, chat:write, im:write
- Also subscribe to Bot Events: app_mention and message.im

## 💻 Usage

### 💬 Channel Mentions

1. Invite the bot to a channel:
    ```bash
    /invite @go_slack


2. Mention the bot:
    ```bash
    @go_slack Hello there!

    - The bot replies in the same channel.

### ✉️ Direct Messages

1. Open a DM with the bot via **Apps → go_slack → Message**  
2. Type a message:
    ```bash
    Hello AI!

    - The bot replies directly in the DM.

---

## 🎬 Demo / Output

### App in Action – ✉️ Direct Messages

### 💬 Chat with the Bot

Open a direct message with the bot in Slack and see it respond in real-time:

- *Bot responding to a simple greeting.*
  
![DM Hello](/demo/hi.png)  

- *Bot providing a movie plot summary in a DM.*

![DM The Matrix Summary](/demo/matrix.png)  

- *Bot helping plan tasks for the day.*
  
![DM Plan Day](/demo/plan.png) 

### 💬 Channel Mentions

- *Bot responding to a simple greeting in a Slack channel.*
- *Bot providing a concise summary of the "Demon Slayer" anime plot in the channel.*
  
![Channel Demon Slayer Summary](/demo/ch_hi.png)  

- *Bot explaining machine learning models briefly in response to a user query in the channel.*
  
![Channel ML Models](/demo/ch_ml.png)  

### 🖥 Terminal Output

    ```bash
    $ go run main.go
    Slack AI Bot started...
    Mention received: @go_slack Hello there!
    Direct message received: Hey bot, summarize AI

## 🧪 Test Cases

| Context      | Input                                | Expected Response                            |
|--------------|--------------------------------------|----------------------------------------------|
| Channel      | `hi`                                 | Friendly greeting                            |
| Channel      | `Summarize "Demon Slayer" anime plot`| Short summary of "Demon Slayer"              |
| Channel      | `Can you tell me about ML models`    | Brief explanation of machine learning models |
| Direct Msg   | `hello`                              | Friendly greeting                            |
| Direct Msg   | `Hey bot, can you summarize AI?`     | Concise AI summary                           |
| Direct Msg   | `Summarize "The Matrix" movie plot.` | Short summary of "The Matrix"                |
| Direct Msg   | `Can you help me plan my day?`       | Sample day plan / task suggestions           |
| Direct Msg   | `Give me motivation quote`           | Short motivational quote                     |

- Screenshots for each test case are included in the `demo/` folder.
  
---

## ✍️ Contributing
- Fork the repo, create a branch, implement features, and submit a pull request.  
- Ensure all changes are tested in a Slack workspace.  
- Add proper logging and error handling for new features.

---

## 🔮 Future Work

- Add Slash Commands for quick interactions.

- Integrate additional AI tools for sentiment analysis, code generation, or data queries.

- Add persistent memory to handle ongoing conversations.

- Deploy on cloud server for 24/7 availability.



