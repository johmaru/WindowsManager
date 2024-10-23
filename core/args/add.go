package args

import (
	"encoding/json"
	"fmt"
	"os"
	"wm/config"
)

func ApAdd(args3 string, args4 string) bool {

	file, err := os.ReadFile("config.json")
	if err != nil {
		fmt.Println("Failed to read config.json:", err)
		return false
	}
	fmt.Println("config.json content:", string(file))
	data := config.JConfig{}
	err = json.Unmarshal(file, &data)
	if err != nil {
		fmt.Println("Failed to parse config.json:", err)
		return false
	}
	fmt.Println("parsed config json", string(file))

	dataFilePath := data.Data
	file, err = os.ReadFile(dataFilePath)
	if err != nil {
		fmt.Println("Failed to read data file:", err)
		return false
	}
	fmt.Println("data.Data content:", string(file))

	var dynamicData map[string]interface{}
	err = json.Unmarshal(file, &dynamicData)
	if err != nil {
		fmt.Println("Failed to parse data file:", err)
		return false
	}
	if dynamicData == nil {
		dynamicData = make(map[string]interface{})
	}
	fmt.Println("Parsed data file:", dynamicData)

	newApp := map[string]interface{}{args3: args4}
	for key, value := range newApp {
		dynamicData[key] = value
	}
	fmt.Println("Updated data:", dynamicData)

	jsonBytes, err := json.MarshalIndent(dynamicData, "", "    ")
	if err != nil {
		fmt.Println("Failed to marshal updated data:", err)
		return false
	}
	fmt.Println("Marshalled updated data:", string(jsonBytes))
	err = os.WriteFile(dataFilePath, jsonBytes, 0644)
	if err != nil {
		fmt.Println("Failed to write to data file:", err)
		return false
	}
	fmt.Println("Successfully wrote to data file")

	file, err = os.ReadFile(dataFilePath)
	if err != nil {
		fmt.Println("Failed to read data file after writing:", err)
		return false
	}
	fmt.Println("data.Data content after writing:", string(file))

	return true
}
