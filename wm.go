package main

import (
	"encoding/json"
	"fmt"
	"os"
	"wm/config"
	"wm/core/args"
)

func main() {

	argumentsCheck()

}

func argumentsCheck() bool {
	err := existCheck()
	if !err {
		return false
	}
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
	_, err = os.ReadDir("Data")
	if os.IsNotExist(err) {
		err = os.Mkdir("Data", 0755)
		if err != nil {
			fmt.Println(err)
			return false
		}
		file, err := os.Create("Data/data.json")
		if err != nil {
			fmt.Println(err)
			return false
		}
		defer file.Close()
		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "    ")
		err = encoder.Encode(map[string]interface{}{})
		if err != nil {
			fmt.Println(err)
			return false
		}
		return true
	} else if err != nil {
		return false
	}

	file, err := os.ReadFile("config.json")
	if err != nil {
		fmt.Println(err)
		return false
	}
	data := config.JConfig{}
	err = json.Unmarshal(file, &data)
	if err != nil {
		fmt.Println("Failed to parse JSON:", err)
		return false
	}
	_, err = os.ReadFile(data.Data)
	if os.IsNotExist(err) {
		file, err := os.Create(data.Data)
		if err != nil {
			fmt.Println(err)
			return false
		}
		defer file.Close()
		msg, err := os.ReadFile("Data/data.json")
		if err != nil {
			fmt.Println(err)
			return false
		}
		if msg == nil {
			return false
		}
		file, err = os.Create(data.Data)
		if err != nil {
			fmt.Println(err)
			return false
		}
		defer file.Close()
		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "    ")
		err = encoder.Encode(msg)
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
