# Slack AI Bot

A **Golang-based Slack AI bot** that responds to both **channel mentions** and **direct messages** using **OpenAI GPT**.  
Built with Slack **Socket Mode**, it demonstrates **modular AI integration** and real-time interactions.

---

## **Features**

- Responds to **@mentions** in Slack channels.  
- Handles **direct messages (DMs)** from users.  
- Integrates with **OpenAI GPT API** for intelligent, context-aware responses.  
- Modular design allows easy extension to additional AI capabilities.  
- Provides **debug logging** to monitor interactions.  

---

## **Tech Stack**

- **Language:** Go (Golang)  
- **Slack SDK:** [slack-go/slack](https://github.com/slack-go/slack)  
- **OpenAI GPT API:** [sashabaranov/go-openai](https://github.com/sashabaranov/go-openai)  
- **Environment Management:** `.env` file for tokens  
- **Deployment Mode:** Local development / Socket Mode  

---

## **Architecture**

## Architecture

+---------------+    +----------------+    +-------------------+    +------------------+    +------------------+
| User Input    |--->| Slack Event    |--->| AI Handler        |--->| OpenAI GPT       |--->| Slack Reply      |
| (Channel/DM)  |    | (mention/DM)   |    | slack_handler.go  |    | ai/gpt_client.go |    | Channel / DM     |
+---------------+    +----------------+    +-------------------+    +------------------+    +------------------+

### **Explanation**

- **User Input:** Messages from Slack users (channel mentions or DMs).  
- **Slack Event:** Captures events using Socket Mode (`app_mention` or `message.im`).  
- **AI Handler:** Processes Slack events, prepares input for GPT, and determines the response.  
- **OpenAI GPT API:** Generates context-aware responses using the GPT model.  
- **Slack Reply:** Sends the final response back to the channel or DM.  

---

## **Setup**

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
  app_mentions:read, chat:write, im:write
- Also subscribe to Bot Events: app_mention and message.im

  ## Usage

### Channel Mentions

1. Invite the bot to a channel:
    ```bash
    /invite @go_slack


2. Mention the bot:
    ```bash
    @go_slack Hello there!

    - The bot replies in the same channel.

### Direct Messages

1. Open a DM with the bot via **Apps → go_slack → Message**  
2. Type a message:
    ```bash
    Hello AI!

    - The bot replies directly in the DM.

---

## Demo

### Channel Mentions

![Channel Demo GIF](demo/channel_01.gif)  
![Channel Demo Screenshot](demo/channel_02.png)  

### Direct Messages

![DM Demo GIF](demo/dm_01.gif)  
![DM Demo Screenshot](demo/dm_02.png)  

### Terminal Output

    ```bash
    $ go run main.go
    Slack AI Bot started...
    Mention received: @go_slack Hello there!
    Direct message received: Hey bot, summarize AI


## Test Cases

| Context | Input                                  | Expected Response             |
| ------- | -------------------------------------- | ----------------------------- |
| Channel | `@go_slack Hello!`                     | Friendly greeting             |
| Channel | `@go_slack Summarize AI news.`         | Short AI news summary         |
| Channel | `@go_slack Write a poem about coding.` | 2–3 line poem                 |
| Channel | `@go_slack Solve 45*12`                | `540`                         |
| DM      | `Hi bot!`                              | Friendly greeting             |
| DM      | `Explain quantum computing simply.`    | Beginner-friendly explanation |
| DM      | `Can you help me plan my day?`         | Sample schedule / task plan   |


Screenshots/GIFs for each test case are included in the demo/ folder.

## Future Work

- Add Slash Commands for quick interactions.

- Integrate additional AI tools for sentiment analysis, code generation, or data queries.

- Add persistent memory to handle ongoing conversations.

- Deploy on cloud server for 24/7 availability.
