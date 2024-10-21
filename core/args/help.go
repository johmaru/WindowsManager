package args

import (
	"fmt"
	"os"
)

type HelpError int

const (
	SetError HelpError = iota
	CheckError
	NormalError
)

func Help() bool {
	_, err := fmt.Print(`Usage: ` + os.Args[0] + ` [command]

Commands:
  help          - Show help
  set <options> - Set the file path
  check <options> - Check the file path
`)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func ErrorHelp(err HelpError) bool {
	switch err {
	case SetError:
		_, err := fmt.Print("Usage: " + os.Args[0] + " set <options> <path>\n")
		if err != nil {
			fmt.Println(err)
			return false
		}
		return true

	case CheckError:
		_, err := fmt.Print("Usage: " + os.Args[0] + " check <options>\n")
		if err != nil {
			fmt.Println(err)
			return false
		}
		return true

	case NormalError:
		_, err := fmt.Print("Usage: " + os.Args[0] + " help\n")
		if err != nil {
			fmt.Println(err)
			return false
		}
		return true
	default:
		fmt.Println("Unknown error")
		return false
	}
	return true
}
