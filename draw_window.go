package main

import "github.com/gdamore/tcell/v2"

func drawWindow(screen tcell.Screen, win *Window) {
	screen.Clear()
	drawCharachters(screen, win)
	drawCursor(screen, win)
	screen.Show()
}

func drawCursor(screen tcell.Screen, win *Window) {
	cursor := *win.cursor
	line := win.lines[win.cursor.line]
	offset := 0
	if len(line) != 0 {
		for i := 0; i < win.cursor.char; i++ {
			if line[i] == '\t' {
				offset += 7
			}
		}
		if win.mode == NormalMode && line[cursor.char] == '\t' {
			offset += 7
		}
	}
	screen.HideCursor()
	screen.SetCursorStyle(win.cursor_style)
	screen.ShowCursor(cursor.char+offset, win.cursor.line)
}

func drawCharachters(screen tcell.Screen, win *Window) {
	for r, line := range win.lines {
		offset := 0
		for c, char := range line {
			switch char {
			case '\t':
				offset += 7
				SetContent(screen, c+offset, r, ' ')
			default:
				SetContent(screen, c+offset, r, char)
			}
		}
	}
}

func SetContent(screen tcell.Screen, col, row int, r rune) {
	screen.SetContent(col, row, r, nil, tcell.Style{})
}
