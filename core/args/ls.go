package args

import (
	"fmt"
	"os"
	"text/tabwriter"
	Jlog "wm/core/log"
)

func Ls() bool {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		Jlog.Log(Jlog.Error, "Failed to get current working directory.", nil)
		return false
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		Jlog.Log(Jlog.Error, "Failed to read directory.", nil)
		return false
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)

	for _, entry := range entries {
		fmt.Fprint(w, entry.Name(), "\t")
	}
	fmt.Fprintln(w)
	w.Flush()

	return true
}
