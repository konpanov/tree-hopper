package main

import (
	"github.com/gdamore/tcell/v2"
)

type Screen = tcell.Screen
type Style = tcell.Style

func drawWindow(screen Screen, win *Window) {
	screen.Clear()
	drawCharachters(screen, win)
	drawCursor(screen, win)
	drawHighlighted(screen, win)
	screen.Show()
}

func drawCursor(screen Screen, win *Window) {
	cursor := *win.cursor
	line := win.lines[win.cursor.line]
	offset := 0
	if len(line) != 0 {
		for i := 0; i < win.cursor.char; i++ {
			offset += RuneWidth(rune(line[i])) - 1
		}
		if win.mode == NormalMode {
			offset += RuneWidth(rune(line[cursor.char])) - 1
		}
	}
	screen.HideCursor()
	screen.SetCursorStyle(win.cursor_style)
	screen.ShowCursor(cursor.char+offset, win.cursor.line-win.topLine)
}

func drawCharachters(s Screen, win *Window) {
	_, h := s.Size()
	for row := 0; row < h && row < len(win.lines); row++ {
		line := win.lines[win.topLine+row]
		drawLine(s, row, line, 0, len(line), win.defStyle)
	}
}

func drawHighlighted(s Screen, win *Window) {
	if win.mode != VisualMode {
		return
	}
	vo := win.visualOrigin
	cr := win.cursor
	hls := vo
	hle := cr
	if vo.line > cr.line {
		hls, hle = hle, hls
	}

	if hls.line == hle.line {
		line := win.lines[hls.line]
		row := hls.line - win.topLine
		from := min(hls.char, hle.char)
		to := max(hls.char, hle.char)
		drawLine(s, row, line, from, to, win.visStyle)
		return
	}

	line := win.lines[hls.line]
	screenRow := hls.line - win.topLine
	drawLine(s, screenRow, line, hls.char, len(line), win.visStyle)

	for row := hls.line + 1; row < hle.line; row++ {
		line := win.lines[row]
		screenRow := row - win.topLine
		drawLine(s, screenRow, line, 0, len(line), win.visStyle)
	}

	line = win.lines[hle.line]
	screenRow = hle.line - win.topLine
	drawLine(s, screenRow, line, 0, hle.char, win.visStyle)
}

func drawLine(s Screen, row int, line string, from, to int, style Style) {
	w, _ := s.Size()
	offset := 0
	to = min(to, len(line))
	if len(line) == 0 {
		s.SetContent(0, row, ' ', nil, style)
	}
	for col := 0; col < from && col < len(line); col++ {
		offset += RuneWidth(rune(line[col])) - 1
	}
	for col := from; col < to && offset+col < w; col++ {
		r := rune(line[col])
		charWidth := RuneWidth(r)
		for i := 0; i < charWidth; i++ {
			x := offset + col + i
			s.SetContent(x, row, r, nil, style)
		}
		offset += charWidth - 1
	}
}

func RuneWidth(r rune) int {
	switch r {
	case '\t':
		return 8
	default:
		return 1
	}
}
