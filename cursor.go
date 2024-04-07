package main

//Moves cursor to the left by one position. Does not skip lines.
func cursorLeft(win *Window) {
	if !isLineStart(win.content, win.cursor) {
		win.cursor--
	}
}

//Moves cursor to the right by one position. Does not skip lines.
func cursorRight(win *Window) {
	if !isLineEnd(win.content, win.cursor) {
		win.cursor++
	}
}

//Moves cursor up one line to the end of the line
func cursorUp(win *Window) {
	for !isLineStart(win.content, win.cursor) {
		win.cursor--
	}
	if win.cursor != 0 {
		win.cursor--
		if win.content[win.cursor] == '\n' {
			win.cursor--
		}
		if win.content[win.cursor] == '\r' {
			win.cursor--
		}
	}
	if win.content[win.cursor] == '\n' {
		win.cursor++
	}
}

//Moves cursor down one line to the start of the line
func cursorDown(win *Window) {
	for !isLineEnd(win.content, win.cursor) {
		win.cursor++
	}
	if win.cursor != len(win.content)-1 {
		win.cursor++
		if win.content[win.cursor] == '\r' {
			win.cursor++
		}
		if win.content[win.cursor] == '\n' {
			win.cursor++
		}
	}
}

func isLineStart(content string, pos int) bool {
	return pos == 0 || isNewLine(content[pos-1])
}
func isLineEnd(content string, pos int) bool {
	return pos == len(content)-1 || isNewLine(content[pos+1])
}

func isNewLine(r byte) bool {
	return r == '\r' || r == '\n'
}
