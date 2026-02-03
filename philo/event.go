package main

import (
	"time"
	"encoding/json"
	"net/http"
)

//contract 
//state → "thinking" | "eating" | "sleeping" | "dead"
//event → "death" | "simulation_end"
type Event struct {
	Time  int64  `json:"time"`
	Philo int    `json:"philo,omitempty"`
	State string `json:"state,omitempty"`
	Event string `json:"event,omitempty"`
}

func (program *Program) emit(ev Event) {
	select {
	case program.Events <- ev:
	default:
		// drop event if UI is slow
	}
}

func (philo *Philo) emitState(state string) {
	philo.prog.emit(Event{
		Time:  time.Since(philo.prog.startTime).Milliseconds(),
		Philo: philo.id,
		State: state,
	})
}

func (program *Program) eventsHandler(w http.ResponseWriter, r *http.Request) {
	// Tell the browser: this is SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")


	// Trigger the simulation to start when the first client connects
    program.once.Do(func() {
        close(program.StartSignal)
    })

	// Make sure streaming is supported
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	// Stop streaming when client disconnects
	ctx := r.Context()

	// Forward events forever
	for {
		select {
		case <-ctx.Done():
			// browser closed the tab
			return

		case ev := <-program.Events:
		w.Write([]byte("data: "))
		_ = json.NewEncoder(w).Encode(ev)
		w.Write([]byte("\n"))
		flusher.Flush()
		}
	}
}
