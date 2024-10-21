package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Data string `json:"data"`
}

func main() {

	existCheck()

	argumentsCheck()

}

func argumentsCheck() bool {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "help":
			fmt.Print("Usage: " + os.Args[0] + " [command]\n\n" + "Commands:\n" + "help - Show help\n" + "set <options> - Set the file path\n" + "check <otions> - Check the file path\n")
		case "ls":
			dir, err := os.Getwd()
			if err != nil {
				fmt.Println(err)
				return false
			}
			entries, err := os.ReadDir(dir)
			if err != nil {
				fmt.Println(err)
				return false
			}
			for _, entry := range entries {
				fmt.Println(entry.Name())
			}
		case "set":
			if len(os.Args) > 3 {
				switch os.Args[2] {
				case "data":
					file, err := os.ReadFile("config.json")
					if err != nil {
						fmt.Println(err)
						return false
					}
					data := Config{}
					json.Unmarshal(file, &data)
					if err != nil {
						fmt.Println("Failed to parse JSON:", err)
						return false
					}
					data.Data = os.Args[3]

					jsonBytes, err := json.MarshalIndent(data, "", "    ")
					if err != nil {
						fmt.Println("Failed to marshal JSON:", err)
						return false
					}

					jsonStr := string(jsonBytes)

					err = os.WriteFile("config.json", []byte(jsonStr), 0644)
					if err != nil {
						fmt.Println("Failed to write to config.json:", err)
						return false
					}
					return true

				default:
					fmt.Print("Usage: " + os.Args[0] + " set <options> <path>\n")
				}

			} else {
				fmt.Print("Usage: " + os.Args[0] + " set <options> <path>\n")
				return true
			}
		case "check":
			if len(os.Args) > 2 {
				switch os.Args[2] {
				case "data":
					file, err := os.ReadFile("config.json")
					if err != nil {
						fmt.Println(err)
						return false
					}
					data := Config{}
					json.Unmarshal(file, &data)
					if err != nil {
						fmt.Println("Failed to parse JSON:", err)
						return false
					}
					fmt.Println(data.Data)
					return true
				}
			} else {
				fmt.Print("Usage: " + os.Args[0] + " check <options>\n")
				return true
			}

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
			Data: "Data/data.json",
		}

		file, err := os.Create("config.json")
		if err != nil {
			fmt.Println(err)
			return false
		}
		defer file.Close()
		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "    ")
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
