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
