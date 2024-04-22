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
	_, err = f.WriteString(win.content)
	if err != nil {
		panic(err)
	}
	f.Sync()
}
