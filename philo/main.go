package main

import (
	"net/http"
	"os"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func startVisualizer(program *Program) {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer conn.Close()

		ticker := time.NewTicker(16 * time.Millisecond)
		defer ticker.Stop()

		for range ticker.C {
			data := make([]byte, program.numPhilos)
			for i := 0; i < program.numPhilos; i++ {
				data[i] = byte(atomic.LoadInt32(&program.philos[i].State))
			}
			if err := conn.WriteMessage(websocket.BinaryMessage, data); err != nil {
				return
			}
		}
	})
	http.ListenAndServe(":8080", nil)
}

func main() {
	argsWithoutExe := os.Args[1:]

	err := validateArgs(argsWithoutExe)
	if err != nil {
		fatal(err)
	}
	program, err := NewProgram(argsWithoutExe)
	if err != nil {
		fatal(err)
	}

	go startVisualizer(program)

	program.Run()
	program.printCPUInfo()
}