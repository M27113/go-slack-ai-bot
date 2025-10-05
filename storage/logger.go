package storage

import (
    "log"
    "os"
)

func LogMessage(user, message, response string) {
    f, err := os.OpenFile("messages.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Println(err)
        return
    }
    defer f.Close()
    logLine := user + " | " + message + " | " + response + "\n"
    if _, err := f.WriteString(logLine); err != nil {
        log.Println(err)
    }
}
