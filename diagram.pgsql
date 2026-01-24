┌──────────┐
│  Client  │
└────┬─────┘
     │  HTTP POST /orders
     ▼
┌──────────────────────────┐
│     HTTP Handler         │
│  (handler.go)            │
│                          │
│ - parse request          │
│ - validate input         │
│ - create Order           │
│ - start goroutine        │
└────┬─────────────────────┘
     │ passes context.Context
     ▼
┌──────────────────────────┐
│  Order Service           │
│  (service.go)            │
│                          │
│ - business logic         │
│ - process order          │
│ - handle errors          │
│ - respect ctx.Done()     │
└────┬─────────────────────┘
     │
     │ updates
     ▼
┌──────────────────────────┐
│ In-Memory Store          │
│ (store.go)               │
│                          │
│ map[orderID]*Order       │
│ + sync.Mutex             │
└──────────────────────────┘


main()
 └── starts HTTP server
     └── handler starts goroutine
         └── goroutine tracked by WaitGroup