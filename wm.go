package main

import (
	"encoding/json"
	"fmt"
	"os"
	"wm/config"
	"wm/core/args"
)

func main() {

	existCheck()

	argumentsCheck()

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
						}
					}
				default:
					args.ErrorHelp(args.ListError)
				}

			} else {
				args.ErrorHelp(args.ListError)
			}
		}

	} else {
		args.ErrorHelp(args.NormalError)
		return true
	}
	return false
}

func existCheck() bool {
	_, err := os.ReadFile("config.json")
	if os.IsNotExist(err) {
		defaultConfig := config.JConfig{
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
