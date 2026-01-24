HTTP Request
   │
   │ r.Context()
   ▼
goroutine ──────────────-┐
   │                     │
   │ processOrder(ctx)   │
   │                     │
   │ select {            │
   │   case <-ctx.Done() │◄── client disconnect / timeout
   │   case work done    │
   │ }                   │
   ▼                     │
wg.Done() ◄──────────────┘


SHUTDOWN

SIGINT (Ctrl+C)
   ↓
main receives signal
   ↓
stop accepting requests
   ↓
waitGroup.Wait()
   ↓
all goroutines exited
   ↓
process exits cleanly
