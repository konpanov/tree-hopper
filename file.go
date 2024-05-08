package main

import (
	"os"
)

func readFile(filename string) string {
	dat, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(dat)
}

func writeFile(win *Window) {
	f, err := os.Create(win.filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.WriteString(windowContent(win))
	if err != nil {
		panic(err)
	}
	err = f.Sync()
	if err != nil {
		panic(err)
	}
}
