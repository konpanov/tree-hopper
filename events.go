package main

import (
	"github.com/gdamore/tcell/v2"
)

func handleEvents(ev tcell.Event, win *Window) {
	switch ev := ev.(type) {
	case *tcell.EventKey:
		handleQuitEvent(win, ev)
		handleInsertModeEvents(win, ev)
		handleNormalModeEvents(win, ev)
		handleVisualModeEvents(win, ev)
	}
}

func handleQuitEvent(window *Window, ev *tcell.EventKey) {
	window.quiting = ev.Key() == tcell.KeyCtrlC
}
