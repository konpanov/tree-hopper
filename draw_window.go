package main

import "github.com/gdamore/tcell/v2"

func drawWindow(screen tcell.Screen, win *Window) {
	screen.Clear()
	screen.SetCursorStyle(win.cursor_style)
	col, row := 0, 0
	for i, r := range win.content {
		if i == win.cursor {
			screen.ShowCursor(col, row)
		}
		switch r {
		case '\r':
		case '\n':
			col = 0
			row++
		case '\t':
			_col := col
			for ; col < _col+8; col++ {
				SetContent(screen, col, row, r)
			}
		default:
			SetContent(screen, col, row, r)
			col++
		}
	}
	screen.Show()
}

func SetContent(screen tcell.Screen, col, row int, r rune) {
	screen.SetContent(col, row, r, nil, tcell.Style{})
}
