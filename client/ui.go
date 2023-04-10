package main

import (
	"github.com/burntcarrot/pairpad/client/editor"
	"github.com/gorilla/websocket"
	"github.com/nsf/termbox-go"
)

// TUI is built using termbox-go.
// termbox allows us to set any content to individual cells, and hence, the basic building block of the editor is a "cell".

// UI creates a new editor view and runs the main loop.
func UI(conn *websocket.Conn) error {
	err := termbox.Init()
	if err != nil {
		return err
	}
	defer termbox.Close()

	e = editor.NewEditor()
	e.SetSize(termbox.Size())
	e.Draw()

	err = mainLoop(conn)
	if err != nil {
		return err
	}

	return nil
}

// mainLoop is the main update loop for the UI.
func mainLoop(conn *websocket.Conn) error {
	// termboxChan is used for sending and receiving termbox events.
	termboxChan := getTermboxChan()

	// msgChan is used for sending and receiving messages.
	msgChan := getMsgChan(conn)

	for {
		select {
		case termboxEvent := <-termboxChan:
			err := handleTermboxEvent(termboxEvent, conn)
			if err != nil {
				return err
			}
		case msg := <-msgChan:
			handleMsg(msg, conn)
		}
	}
}