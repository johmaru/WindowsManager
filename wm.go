package main

import (
	"encoding/json"
	"fmt"
	"os"
	"wm/config"
	"wm/core/args"
	Jlog "wm/core/log"
)

func main() {

	if !argumentsCheck() {
		fmt.Println("Initialization failed.")
		os.Exit(1)
	}
	if !ExistCheck() {
		fmt.Println("Initialization failed.")
		os.Exit(1)
	}

}

func argumentsCheck() bool {
	logger, err := Jlog.InitLog()
	if err != nil {
		fmt.Println("Failed to initialize log:", err)
		return false
	}
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "help":
			Jlog.Log(Jlog.Info, "Help command executed.", logger)
			args.Help()
		case "ls":
			Jlog.Log(Jlog.Info, "Ls command executed.", logger)
			args.Ls()
		case "check":
			Jlog.Log(Jlog.Info, "Check command executed.", logger)
			if len(os.Args) > 2 {
				switch os.Args[2] {
				case "data":
					args.Check()

				case "path":
					if len(os.Args) > 3 {
						args.PathCheck(os.Args[3])
					} else {
						args.ErrorHelp(args.CheckError)
					}
				}
			} else {
				args.ErrorHelp(args.CheckError)
				return true
			}

		case "set":
			Jlog.Log(Jlog.Info, "Set command executed.", logger)
			if len(os.Args) > 3 {
				switch os.Args[2] {
				case "data":

					args.Set(os.Args[3])
				default:
					args.ErrorHelp(args.SetError)
				}
			}

		case "list":
			Jlog.Log(Jlog.Info, "List command executed.", logger)
			if len(os.Args) > 2 {
				switch os.Args[2] {
				case "options":
					if len(os.Args) > 3 {
						switch os.Args[3] {
						case "help":
							args.HelpList()

						case "set":
							args.SetList()

						case "check":
							args.CheckList()

						case "ls":
							args.LsList()

						case "list":
							args.List()

						case "add":
							args.AddList()
						}
					}
				default:
					args.ErrorHelp(args.ListError)
				}

			} else {
				args.ErrorHelp(args.ListError)
			}

		case "add":
			Jlog.Log(Jlog.Info, "Add command executed.", logger)
			if len(os.Args) > 2 {
				switch os.Args[2] {
				case "ap":
					if len(os.Args) > 4 {
						err := args.ApAdd(os.Args[3], os.Args[4])
						if !err {
							fmt.Println(err)
						}
					} else {
						args.ErrorHelp(args.AddError)
					}
				}
			}
		}

	} else {
		args.ErrorHelp(args.NormalError)
		return true
	}
	return true
}

func ExistCheck() bool {
	// config.json exists check
	_, err := os.Stat("config.json")
	if os.IsNotExist(err) {
		if err != nil {
			fmt.Println("Failed to initialize log:", err)
			Jlog.Log(Jlog.Error, "Failed to initialize log.", nil)
			return false
		}

		defaultConfig := config.JConfig{
			Data: "Data",
			Log:  "Log",
		}

		file, err := os.Create("config.json")
		if err != nil {
			fmt.Println("Failed to create config.json:", err)
			Jlog.Log(Jlog.Error, "Failed to create config.json.", nil)
			return false
		}
		defer file.Close()
		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "    ")
		err = encoder.Encode(defaultConfig)
		if err != nil {
			fmt.Println("Failed to write to config.json:", err)
			Jlog.Log(Jlog.Error, "Failed to write to config.json.", nil)
			return false
		}
	}

	// config.json Object check
	file, err := os.ReadFile("config.json")
	if err != nil {
		fmt.Println("Failed to read config.json:", err)
		Jlog.Log(Jlog.Error, "Failed to read config.json.", nil)
		return false
	}

	data := config.JConfig{}
	err = json.Unmarshal(file, &data)
	if err != nil {
		fmt.Println("Failed to parse config.json:", err)
		Jlog.Log(Jlog.Error, "Failed to parse config.json.", nil)
		return false
	}

	LogDirPath := data.Log
	// Log directory exists check
	_, err = os.Stat(LogDirPath)
	if os.IsNotExist(err) {
		err = os.Mkdir(LogDirPath, 0755)
		if err != nil {
			fmt.Println("Failed to create Log directory:", err)
			Jlog.Log(Jlog.Error, "Failed to create Log directory.", nil)
			return false
		}
	}

	if data.Data == "" {
		file, err := os.OpenFile("config.json", os.O_RDWR, 0644)
		if err != nil {
			fmt.Println("Failed to open config.json:", err)
			Jlog.Log(Jlog.Error, "Failed to open config.json.", nil)
			return false
		}
		defer file.Close()
		data.Data = "Data"
		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "    ")
		err = encoder.Encode(data)
		if err != nil {
			fmt.Println("Failed to write to config.json:", err)
			Jlog.Log(Jlog.Error, "Failed to write to config.json.", nil)
			return false
		}
	}
	if data.Log == "" {
		file, err := os.OpenFile("config.json", os.O_RDWR, 0644)
		if err != nil {
			fmt.Println("Failed to open config.json:", err)
			Jlog.Log(Jlog.Error, "Failed to open config.json.", nil)
			return false
		}
		defer file.Close()
		data.Log = "Log"
		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "    ")
		err = encoder.Encode(data)
		if err != nil {
			fmt.Println("Failed to write to config.json:", err)
			Jlog.Log(Jlog.Error, "Failed to write to config.json.", nil)
			return false
		}
	}

	dataDirPath := data.Data
	// Data directory exists check
	_, err = os.Stat(dataDirPath)
	if os.IsNotExist(err) {
		err = os.Mkdir(dataDirPath, 0755)
		if err != nil {
			fmt.Println("Failed to create Data directory:", err)
			Jlog.Log(Jlog.Error, "Failed to create Data directory.", nil)
			return false
		}
	}

	// Data/data.json exists check
	dataFilePath := dataDirPath + "/data.json"
	_, err = os.Stat(dataFilePath)
	if os.IsNotExist(err) {
		file, err := os.Create(dataFilePath)
		if err != nil {
			fmt.Println("Failed to create data.json:", err)
			Jlog.Log(Jlog.Error, "Failed to create data.json.", nil)
			return false
		}
		defer file.Close()
		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "    ")
		err = encoder.Encode(map[string]interface{}{})
		if err != nil {
			fmt.Println("Failed to write to data.json:", err)
			Jlog.Log(Jlog.Error, "Failed to write to data.json.", nil)
			return false
		}
	}

	return true
}
