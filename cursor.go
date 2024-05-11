package main

//Moves cursor to the left by one position. Does not skip lines.
func cursorLeft(buf *Buffer) {
	buf.cursor.char = max(buf.cursor.char-1, 0)
}

//Moves cursor to the right by one position. Does not skip lines.
func cursorRight(buf *Buffer) {
	buf.cursor.char = min(buf.cursor.char+1, mostRight(buf))
}

//Moves cursor up one line to the end of the line
func cursorUp(buf *Buffer) {
	buf.cursor.line = max(buf.cursor.line-1, 0)
	buf.cursor.char = min(buf.cursor.char, mostRight(buf))
	buf.topLine = min(buf.topLine, buf.cursor.line)
}

//Moves cursor down one line to the start of the line
func cursorDown(buf *Buffer) {
	line_count := len(buf.lines)
	buf.cursor.line = min(buf.cursor.line+1, line_count-1)
	buf.cursor.char = min(buf.cursor.char, mostRight(buf))
	buf.topLine = max(buf.topLine, buf.cursor.line-buf.height+1)
}

func mostRight(buf *Buffer) int {
	if buf.mode == InsertMode {
		return len(buf.lines[buf.cursor.line])
	} else {
		return max(len(buf.lines[buf.cursor.line])-1, 0)
	}
}
