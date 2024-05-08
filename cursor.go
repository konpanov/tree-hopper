package main

//Moves cursor to the left by one position. Does not skip lines.
func cursorLeft(win *Window) {
	win.cursor.char = max(win.cursor.char-1, 0)
}

//Moves cursor to the right by one position. Does not skip lines.
func cursorRight(win *Window) {
	win.cursor.char = min(win.cursor.char+1, mostRight(win))
}

//Moves cursor up one line to the end of the line
func cursorUp(win *Window) {
	win.cursor.line = max(win.cursor.line-1, 0)
	win.cursor.char = min(win.cursor.char, mostRight(win))
}

//Moves cursor down one line to the start of the line
func cursorDown(win *Window) {
	line_count := len(win.lines)
	win.cursor.line = min(win.cursor.line+1, line_count-1)
	win.cursor.char = min(win.cursor.char, mostRight(win))
}

func mostRight(win *Window) int {
	if win.mode == InsertMode {
		return len(win.lines[win.cursor.line])
	} else {
		return max(len(win.lines[win.cursor.line])-1, 0)
	}
}
