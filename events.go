package main

import (
	"github.com/gdamore/tcell/v2"
)

func handleEvents(ev tcell.Event, buf *Buffer) {
	switch ev := ev.(type) {
	case *tcell.EventKey:
		handleQuitEvent(buf, ev)
		handleInsertModeEvents(buf, ev)
		handleNormalModeEvents(buf, ev)
		handleVisualModeEvents(buf, ev)
	}
}

func handleQuitEvent(buf *Buffer, ev *tcell.EventKey) {
	buf.quiting = ev.Key() == tcell.KeyCtrlC
}
