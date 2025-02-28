package main

import (
	"os"
)

func readFile(filename string) string {
	dat, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(dat[:len(dat)-2])
}

func writeFile(buf *Buffer) {
	f, err := os.Create(buf.filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.WriteString(bufferContent(buf))
	if err != nil {
		panic(err)
	}
	err = f.Sync()
	if err != nil {
		panic(err)
	}
}
