package args

import (
	"encoding/json"
	"fmt"
	"os"
	"wm/config"
)

func Check() bool {
	file, err := os.ReadFile("config.json")
	if err != nil {
		fmt.Println(err)
		return false
	}
	data := config.JConfig{}
	json.Unmarshal(file, &data)
	if err != nil {
		fmt.Println("Failed to parse JSON:", err)
		return false
	}
	fmt.Println(data.Data)
	return true
}

func PathCheck(key string) bool {
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

	filepath := data.Data + "/data.json"

	file, err = os.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
		return false
	}

	var dynamicData map[string]interface{}
	err = json.Unmarshal(file, &dynamicData)
	if err != nil {
		fmt.Println("Failed to parse data file:", err)
		return false
	}
	if dynamicData == nil {
		fmt.Print("No data found")
		return false
	}

	if value, exist := dynamicData[key]; exist {
		fmt.Println(value)
	} else {
		fmt.Println("Key not found")
		return false
	}

	return true
}
