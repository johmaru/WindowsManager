package main

import (
	"encoding/json"
	"fmt"
	"os"
	"wm/config"
	"wm/core/args"
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
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "help":
			args.Help()
		case "ls":
			args.Ls()
		case "check":
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
			if len(os.Args) > 3 {
				switch os.Args[2] {
				case "data":

					args.Set(os.Args[3])
				default:
					args.ErrorHelp(args.SetError)
				}
			}

		case "list":
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
		defaultConfig := config.JConfig{
			Data: "Data",
		}

		file, err := os.Create("config.json")
		if err != nil {
			fmt.Println("Failed to create config.json:", err)
			return false
		}
		defer file.Close()
		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "    ")
		err = encoder.Encode(defaultConfig)
		if err != nil {
			fmt.Println("Failed to write to config.json:", err)
			return false
		}
	}

	file, err := os.ReadFile("config.json")
	if err != nil {
		fmt.Println("Failed to read config.json:", err)
		return false
	}

	data := config.JConfig{}
	err = json.Unmarshal(file, &data)
	if err != nil {
		fmt.Println("Failed to parse config.json:", err)
		return false
	}

	dataDirPath := data.Data
	// Data directory exists check
	_, err = os.Stat(dataDirPath)
	if os.IsNotExist(err) {
		err = os.Mkdir(dataDirPath, 0755)
		if err != nil {
			fmt.Println("Failed to create Data directory:", err)
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
			return false
		}
		defer file.Close()
		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "    ")
		err = encoder.Encode(map[string]interface{}{})
		if err != nil {
			fmt.Println("ExistCheckScope:Failed to write to data.json:", err)
			return false
		}
	}

	return true
}
