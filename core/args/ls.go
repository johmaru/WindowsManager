package args

import (
	"fmt"
	"os"
)

func Ls() bool {
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
	return true
}
