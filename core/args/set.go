package args

import (
	"encoding/json"
	"fmt"
	"os"
	"wm/config"
)

func Set(path string) bool {
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
	data.Data = path

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
}
