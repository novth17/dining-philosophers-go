# Concurrent Order Processor (Go Learning Lab)

A high-concurrency simulation of an order-processing system built to explore the Go CSP (Communicating Sequential Processes) model. This project serves as a practical deep-dive into Go's concurrency primitives, memory safety, and the differences between Go and low-level languages like C/C++.

## 🚀 Learning Objectives
- Master Goroutine orchestration and lifecycle management.
- Implement robust cancellation patterns using the Context API.
- Detect and resolve non-deterministic Data Races.

Compare Go's user-space scheduling with OS-level threading (pthreads).


## 🛠 Features (In-Progress)
[ ] Asynchronous Scooping: Concurrent task execution using sync.WaitGroup.

[ ] The Melting Timer: Deadline propagation and task cancellation via context.WithTimeout.

[ ] Thread-Safe Accounting: Shared state protection using sync.Mutex.

[ ] HTTP Interface: RESTful entry points for order ingestion.

[ ] Persistent Inventory: Transitioning from in-memory maps to a persistent store.

## 🔬 Technical Deep-Dive
### Concurrency & Synchronization
Unlike the pthread_create and pthread_join patterns used in C, this project utilizes Go’s lightweight Goroutines. I utilized a sync.WaitGroup to ensure the main process (the "Shop Owner") waits for all background workers to finish before exiting.

### Cancellation Logic
To prevent "Goroutine leaks" (similar to memory leaks in C), I implemented the Context API. Every worker listens to the ctx.Done() signal within a select block. If a timeout is reached or a client disconnects, the worker immediately ceases operations and cleans up resources.

### The Race Detector
A key milestone was identifying a data race in the revenue counter. By running the Go Race Detector, I identified unsynchronized memory access where multiple goroutines were performing non-atomic increments.

## 📜 Reflections for C/C++ Developers
Coming from a background in C and C++, the most striking difference is the Go Scheduler (GMP Model). While I am used to managing the cost of OS thread context switches, Go’s ability to multiplex thousands of goroutines onto a handful of threads at a 2KB stack cost is a significant paradigm shift in how I approach scalability.
