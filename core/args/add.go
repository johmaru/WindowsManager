package args

import (
	"encoding/json"
	"fmt"
	"os"
	"wm/config"
	Jlog "wm/core/log"
)

func ApAdd(args3 string, args4 string) bool {

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

	dataFilePath := data.Data + "/data.json"
	file, err = os.ReadFile(dataFilePath)
	if err != nil {
		fmt.Println("Failed to read data file:", err)
		Jlog.Log(Jlog.Error, "Failed to read data file.", nil)
		return false
	}

	var dynamicData map[string]interface{}
	err = json.Unmarshal(file, &dynamicData)
	if err != nil {
		fmt.Println("Failed to parse data file:", err)
		Jlog.Log(Jlog.Error, "Failed to parse data file.", nil)
		return false
	}
	if dynamicData == nil {
		dynamicData = make(map[string]interface{})
	}

	newApp := map[string]interface{}{args3: args4}
	for key, value := range newApp {
		dynamicData[key] = value
	}

	jsonBytes, err := json.MarshalIndent(dynamicData, "", "    ")
	if err != nil {
		fmt.Println("Failed to marshal updated data:", err)
		Jlog.Log(Jlog.Error, "Failed to marshal updated data.", nil)
		return false
	}

	err = os.WriteFile(dataFilePath, jsonBytes, 0644)
	if err != nil {
		fmt.Println("Failed to write to data file:", err)
		Jlog.Log(Jlog.Error, "Failed to write to data file.", nil)
		return false
	}

	return true
}
