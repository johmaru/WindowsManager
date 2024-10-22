package args

import (
	"fmt"
)

func HelpList() bool {
	_, err := fmt.Print(`Not exist a Options to help command` + "\n")
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func SetList() bool {
	_, err := fmt.Print(`Example : wm set <FirstOption> <Path>
		                 
FirstOption:

	-	data - Set the data file path
								
Path:		
							
	-   path to your choiced file location` + "\n")
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func CheckList() bool {
	_, err := fmt.Print(`Example : wm check <FirstOption>
	
FirstOption:
						 
	-	data - Check the data file path` + "\n")
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func LsList() bool {
	_, err := fmt.Print(`Not exist a Options to ls command` + "\n")
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func List() bool {
	_, err := fmt.Print(`Example : wm list <options> <command>
	                 
Options:


	 -   options - List the options of the command

Command:

	-   set - List the set command options

	-   check - List the check command options

	-   ls - List the ls command options

	-   list - List the list command options

	-   help - List the help command options

	`)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
