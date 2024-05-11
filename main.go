package main

import (
	"log"
	"os"
	"strings"

	"github.com/gdamore/tcell/v2"
)

type Mode int8

const (
	NormalMode Mode = iota
	InsertMode Mode = iota
	VisualMode Mode = iota
)

type CursorPos struct {
	line int
	char int
}

type Buffer struct {
	filename     string
	content      string
	lines        []string
	defStyle     tcell.Style
	visStyle     tcell.Style
	cursor_style tcell.CursorStyle
	cursor       *CursorPos
	visualOrigin *CursorPos
	quiting      bool
	mode         Mode
	height       int
	width        int
	newLineChar  string
	topLine      int
}

func main() {
	filename := "main.go"
	if len(os.Args) >= 2 {
		filename = os.Args[1]
	}
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	w, h := screen.Size()
	buf := createBufferFromFile(filename, w, h)
	defer quit(screen, buf)
	for !buf.quiting {
		drawWindow(screen, buf)
		handleEvents(screen.PollEvent(), buf)
	}
}

func createBufferFromString(content string, width, height int) *Buffer {
	newLineChar := "\r\n"
	lines := strings.Split(content, newLineChar)
	visColor := tcell.ColorGray
	buf := &Buffer{
		filename:     "",
		content:      content,
		lines:        lines,
		cursor:       &CursorPos{0, 0},
		visualOrigin: &CursorPos{0, 0},
		cursor_style: tcell.CursorStyleSteadyBlock,
		quiting:      false,
		mode:         NormalMode,
		width:        width,
		height:       height,
		newLineChar:  newLineChar,
		topLine:      0,
		defStyle:     tcell.StyleDefault,
		visStyle:     tcell.StyleDefault.Background(visColor),
	}
	return buf
}

func createBufferFromFile(filename string, width, height int) *Buffer {
	content := readFile(filename)
	buf := createBufferFromString(content, width, height)
	buf.filename = filename
	return buf
}

func bufferContent(buf *Buffer) string {
	return strings.Join(buf.lines, buf.newLineChar)
}

func quit(screen tcell.Screen, buf *Buffer) {
	maybePanic := recover()
	screen.Fini()
	if maybePanic != nil {
		panic(maybePanic)
	}
}
