package main

import (
	"encoding/json"
	"log"
	"syscall/js"
	"time"
)

var message = make(chan []byte)

func main() {
	js.Global().Set("sendMessage", js.FuncOf(sendMessage))
	for m := range message {
		postToJs(m)
	}
}

func sendMessage(this js.Value, args []js.Value) any {
	time.Sleep(2 * time.Second)
	user := struct {
		Name    string `json:"name"`
		Surname string `json:"surname"`
	}{"George", "Spanos"}
	body, err := json.Marshal(user)
	if err != nil {
		log.Fatalln("failed to marshal")
	}
	message <- body
	return nil
}

func postToJs(message []byte) {
	js.Global().Call("postMessage", string(message))
}

// func sendError(err error) {
// 	message <- []byte(err.Error())
// }
