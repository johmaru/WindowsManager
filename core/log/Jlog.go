package Jlog

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
	"wm/config"

	"go.uber.org/zap"
)

type LogLevel int

const (
	Info LogLevel = iota
	Warning
	Error
	Debug
	Critical
)

func InitLog() (*zap.SugaredLogger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	sugar := logger.Sugar()
	return sugar, nil
}

func Log(level LogLevel, message string, logger *zap.SugaredLogger) bool {
	Date := time.Now().Format("2006-01-02")
	file, err := os.ReadFile("config.json")
	if err != nil {
		fmt.Println("Failed to read config.json:", err)
		return false
	}
	data := config.JConfig{}
	err = json.Unmarshal(file, &data)
	if err != nil {
		fmt.Println("Failed to unmarshal config.json:", err)
		return false
	}

	logPath := data.Log + "/" + Date + ".log"

	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		file, err := os.Create(logPath)
		if err != nil {
			fmt.Println("Failed to create log file:", err)
			return false
		}
		defer file.Close()
		file.WriteString("Log file created at: " + time.Now().Format("2006-01-02 15:04:05") + "\n" + "\n" + "\n")
	}

	logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Failed to open log file:", err)
		return false
	}
	defer logFile.Close()
	switch level {
	case Info:
		_, err := logFile.WriteString(time.Now().Format("2006-01-02 15:04:05 ") + "-INFO: " + message + "\n")
		if err != nil {
			fmt.Println("Failed to write to log file:", err)
			return false
		}
	case Warning:
		_, err := logFile.WriteString(time.Now().Format("2006-01-02 15:04:05 ") + "WARNING: " + message + "\n")
		if err != nil {
			fmt.Println("Failed to write to log file:", err)
			return false
		}
	case Error:
		_, err := logFile.WriteString(time.Now().Format("2006-01-02 15:04:05 ") + "ERROR: " + message + "\n")
		if err != nil {
			fmt.Println("Failed to write to log file:", err)
			return false
		}
	case Debug:
		_, err := logFile.WriteString(time.Now().Format("2006-01-02 15:04:05 ") + "DEBUG: " + message + "\n")
		if err != nil {
			fmt.Println("Failed to write to log file:", err)
			return false
		}
	case Critical:
		_, err := logFile.WriteString(time.Now().Format("2006-01-02 15:04:05 ") + "CRITICAL: " + message + "\n")
		if err != nil {
			fmt.Println("Failed to write to log file:", err)
			return false
		}
	default:
		fmt.Println("Unknown log level")
	}
	return true
}
