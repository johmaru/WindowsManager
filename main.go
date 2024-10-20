package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Data   string `json:"data"`
	Config string `json:"config"`
}

func main() {

	existCheck()

	argumentsCheck()

}

func argumentsCheck() bool {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "help":
			fmt.Print("Usage: " + os.Args[0] + " [command]\n" + "Commands:\n" + "help - Show help\n")
		}
	} else {
		fmt.Print("Usage: " + os.Args[0] + "help\n")
		return true
	}
	return false
}

func existCheck() bool {
	_, err := os.ReadFile("config.json")
	if os.IsNotExist(err) {
		defaultConfig := Config{
			Data:   "Data/data.json",
			Config: "config.json",
		}

		file, err := os.Create("config.json")
		if err != nil {
			fmt.Println(err)
			return false
		}
		defer file.Close()
		encoder := json.NewEncoder(file)
		err = encoder.Encode(defaultConfig)
		if err != nil {
			fmt.Println(err)
			return false
		}
		return true
	} else if err != nil {
		return false
	}
	return true
}
