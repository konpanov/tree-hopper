package main

import (
	"github.com/gdamore/tcell/v2"
)

type Screen = tcell.Screen
type Style = tcell.Style

func drawWindow(screen Screen, buf *Buffer) {
	screen.Clear()
	drawCharachters(screen, buf)
	drawCursor(screen, buf)
	drawHighlighted(screen, buf)
	screen.Show()
}

func drawCursor(screen Screen, buf *Buffer) {
	cursor := *buf.cursor
	line := buf.lines[buf.cursor.line]
	offset := 0
	if len(line) != 0 {
		for i := 0; i < buf.cursor.char; i++ {
			offset += RuneWidth(rune(line[i])) - 1
		}
		if buf.mode == NormalMode {
			offset += RuneWidth(rune(line[cursor.char])) - 1
		}
	}
	screen.HideCursor()
	screen.SetCursorStyle(buf.cursor_style)
	screen.ShowCursor(cursor.char+offset, buf.cursor.line-buf.topLine)
}

func drawCharachters(s Screen, buf *Buffer) {
	_, h := s.Size()
	for row := 0; row < h && row < len(buf.lines); row++ {
		line := buf.lines[buf.topLine+row]
		drawLine(s, row, line, 0, len(line), buf.defStyle)
	}
}

func drawHighlighted(s Screen, buf *Buffer) {
	if buf.mode != VisualMode {
		return
	}
	vo := buf.visualOrigin
	cr := buf.cursor
	hls := vo
	hle := cr
	if vo.line > cr.line {
		hls, hle = hle, hls
	}

	if hls.line == hle.line {
		line := buf.lines[hls.line]
		row := hls.line - buf.topLine
		from := min(hls.char, hle.char)
		to := max(hls.char, hle.char)
		drawLine(s, row, line, from, to, buf.visStyle)
		return
	}

	line := buf.lines[hls.line]
	screenRow := hls.line - buf.topLine
	drawLine(s, screenRow, line, hls.char, len(line), buf.visStyle)

	for row := hls.line + 1; row < hle.line; row++ {
		line := buf.lines[row]
		screenRow := row - buf.topLine
		drawLine(s, screenRow, line, 0, len(line), buf.visStyle)
	}

	line = buf.lines[hle.line]
	screenRow = hle.line - buf.topLine
	drawLine(s, screenRow, line, 0, hle.char, buf.visStyle)
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
