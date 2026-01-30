# Dining philosophers problem - Solved In Go

A concurrent simulation of the classic [Dining Philosophers Problem](https://en.wikipedia.org/wiki/Dining_philosophers_problem) built in Go, demonstrating advanced concurrency patterns, deadlock prevention, and goroutine orchestration.

## 🎯 Overview

This project simulates multiple philosophers running concurrently as goroutines, competing for shared resources (forks) protected by mutexes. Each philosopher cycles through three states:

- **Thinking** - Philosophical contemplation (variable duration)
- **Eating** - Consuming spaghetti (requires 2 forks)
- **Sleeping** - Resting after a meal

The simulation must avoid deadlocks, prevent data races, and guarantee clean goroutine shutdown while detecting starvation scenarios.

## 🍝 The Problem

### Classic Dining Philosophers Problem

The [original problem](https://en.wikipedia.org/wiki/Dining_philosophers_problem), proposed by Edsger Dijkstra in 1965, explores resource contention and deadlock in concurrent systems. Philosophers sit at a round table with one fork between each pair, and they must acquire both neighboring forks to eat.

### This Implementation: A Survival Simulation

This project adds a **time-critical survival constraint** that transforms the problem from a theoretical exercise into a real-time scheduling challenge:

**Key Differences:**

| Aspect | Classic Problem | This Implementation |
|--------|----------------|---------------------|
| **Goal** | Avoid deadlock | Avoid deadlock + Keep everyone alive |
| **Time Constraint** | None | Philosophers die if they don't eat within `time_to_die` |
| **Termination** | Runs indefinitely | Stops on first death OR when meal quota is reached |
| **Monitoring** | No external observer | Dedicated monitor goroutine tracks health |
| **State Logging** | Not specified | Every state change must be logged with timestamps |
| **Death Detection** | N/A | Must detect and report death within 10ms |

**Why This Matters:**

The time constraint means **fairness is mandatory**, not optional. In the classic problem, a greedy scheduling algorithm might work. Here, if even one philosopher is starved, the entire simulation fails. This forces consideration of:

- Goroutine scheduling fairness
- Mutex acquisition order (resource hierarchy)
- Wake-up synchronization patterns
- Real-time monitoring without interference

**The Challenge:**

Given parameters like `5 philosophers, 800ms time_to_die, 200ms time_to_eat, 200ms time_to_sleep`, you have only **400ms of margin** (800ms - 400ms work). Any scheduling unfairness, mutex contention, or goroutine starvation will cause death.

The simulation becomes a test of **concurrent systems design under hard real-time constraints**.

### Visual Representation

```
         Philosopher 1
              🧠
         Fork 1   Fork 5
              |   |
    Philo 2 🧠---🍝---🧠 Philo 5
         |             |
      Fork 2        Fork 4
         |             |
    Philo 3 🧠-------🧠 Philo 4
              |
           Fork 3

Circular Table:
- 5 Philosophers → 5 Forks
- Each philosopher needs 2 forks to eat
- Fork N is shared between Philosopher N and N+1
- Philosopher 1 shares Fork 5 with Philosopher 5 (circular)
```

**Single Philosopher Edge Case:**
```
    Philosopher 1
         🧠
         |
      Fork 1

Problem: Only 1 fork available, but 2 needed to eat
Result: Impossible scenario → Must die within time_to_die
```

## 📋 Rules & Constraints

Based on the project specification, the implementation must follow these strict rules:

### Input Parameters

```bash
./philo <number_of_philosophers> <time_to_die> <time_to_eat> <time_to_sleep> [number_of_times_each_philosopher_must_eat]
```

- All times are in **milliseconds**
- All values must be **strictly positive**
- Philosophers are numbered 1 to N (not 0-indexed)
- Philosopher 1 sits next to philosopher N (circular table)

### Behavioral Rules

1. **No Communication** - Philosophers cannot talk to each other or know others' states
2. **No Global Variables** - All state must be properly encapsulated
3. **Death Detection** - A philosopher dies if they haven't started eating within `time_to_die` ms since their last meal (or simulation start)
4. **Fork Requirement** - Must acquire **both** left and right forks before eating
5. **State Logging** - Every state change must be logged with precise timestamps:
   ```
   timestamp_in_ms X has taken a fork
   timestamp_in_ms X is eating
   timestamp_in_ms X is sleeping
   timestamp_in_ms X is thinking
   timestamp_in_ms X died
   ```

### Termination Conditions

The simulation stops when **either**:
- A philosopher dies of starvation, OR
- All philosophers have eaten `number_of_times_each_philosopher_must_eat` meals (if specified)

### Correctness Requirements

- ✅ **No Data Races** - Must pass `go test -race`
- ✅ **No Deadlocks** - System must never freeze
- ✅ **Accurate Death Detection** - Death must be reported within 10ms of occurrence
- ✅ **No Message Overlap** - State logs must be synchronized (no garbled output)

## 🚀 Quick Start

### Build

```bash
make
```

### Run

```bash
./bin/philo <number_of_philosophers> <time_to_die> <time_to_eat> <time_to_sleep> [number_of_times_each_philosopher_must_eat]
```

**Parameters:**
- `number_of_philosophers` - Number of philosophers (and forks)
- `time_to_die` - Time (ms) a philosopher can survive without eating
- `time_to_eat` - Time (ms) it takes to eat
- `time_to_sleep` - Time (ms) spent sleeping
- `number_of_times_each_philosopher_must_eat` - *Optional* - Simulation ends when all philosophers reach this meal count

**Examples:**

```bash
# 5 philosophers, 800ms to live, 200ms eating, 200ms sleeping, 7 meals each
./bin/philo 5 800 200 200 7

# 3 philosophers, tight timing (tests starvation detection)
./bin/philo 3 610 200 200

# Single philosopher (impossible scenario - only 1 fork)
./bin/philo 1 800 200 200
```

### Test

```bash
# Run all tests
make test

# Run with race detector
go test -race -v

# Run specific test
go test -run TestPhiloStarvation/Standard_Success -v
```

## 🏗️ Architecture

### Core Components

```
philo/
├── main.go           # Entry point and argument parsing
├── program.go        # Program state and initialization
├── philosopher.go    # Philosopher struct and state
├── routine.go        # Philosopher lifecycle (think, eat, sleep)
├── monitor.go        # Starvation detection and termination
├── run.go            # Goroutine orchestration with WaitGroup
└── philo_test.go     # Table-driven integration tests
```

### Concurrency Model
```
┌───────────────────────────────────────────────────┐
│              Main Goroutine                       │
│  • Spawns philosophers & monitor                  │
│  • Waits on sync.WaitGroup                        │
└───────────────────────────────────────────────────┘
                        │
         ┌──────────────┼──────────────┐
         ▼              ▼              ▼
    ┌────────┐     ┌────────┐     ┌────────┐
    │ Philo 1│     │ Philo 2│ ... │ Philo N│
    │  (go)  │     │  (go)  │     │  (go)  │
    └────────┘     └────────┘     └────────┘
         │              │              │
         └──────────────┼──────────────┘
                        ▼
               ┌─────────────────┐
               │    Monitor (go) │
               │  • Check health │
               │  • call cancel()│
               └─────────────────┘
```
**Key Design Choices:**

1. **Mutexes over Channels** - Forks are modeled as `sync.Mutex` rather than channels because they represent exclusive ownership of stationary resources, not message passing. (See [Technical Deep Dive](#-technical-deep-dive) for analysis)

2. **Decoupled Monitor** - A separate goroutine monitors philosopher health without interfering with their routines, using `context.WithCancel()` for termination signaling

3. **WaitGroup + Context** - Combines `sync.WaitGroup` (tracks completion) with `context.Context` (signals cancellation) for clean shutdown

## 🔬 Technical Deep Dive

### Deadlock Prevention: Resource Hierarchy

The simulation implements **Dijkstra's Solution** - philosophers acquire forks in a strict order based on fork ID to break circular wait conditions:

```go
// Even-numbered philosophers: pick left fork first
if philo.id%2 == 0 {
    philo.forks.left.Lock()
    philo.forks.right.Lock()
} else {
    // Odd-numbered philosophers: pick right fork first
    philo.forks.right.Lock()
    philo.forks.left.Lock()
}
```

This eliminates the circular dependency chain at the source - at least one philosopher will always acquire both forks, preventing system-wide deadlock.

### Fairness Heuristic: Rotation Delay

For **odd-numbered philosopher counts** (e.g., 5 philosophers), perfect pairing is impossible. This creates a "desynchronization" problem where one philosopher is always "odd out."

**Solution:** Dynamic think time calculation for odd-numbered philosophers:

```go
thinkTime := (eatTime * 2) - sleepTime
```

**Why this works:**
- Without delay, the "odd philosopher" wakes early and can monopolize forks ("greedy thread")
- This formula calculates the exact "earliness" offset
- By delaying, we synchronize their fork requests with neighbors' completion times
- Result: Smooth rotational wave of resource acquisition instead of chaotic races

**Note:** This is an empirically-derived heuristic, not a formal proof. It works well for tested parameter ranges but isn't guaranteed for all edge cases.

### Race-Free Termination: Context + Mutex Interaction

**The Problem:** `sync.Mutex.Lock()` is **not context-aware** - it cannot be interrupted by cancellation.

**Symptom:** "Ghost meals" - philosophers record meals *after* the simulation has been cancelled:

```
200 1 died
--- Final Meal Statistics ---
Philosopher 1 finished 1/5 meals  // ← ghost!
```

**Root Cause:** When cancellation occurs while a philosopher waits for a fork, they don't stop immediately. They hang until they acquire the lock, then execute post-lock code.

**Solution:** Check context status **immediately after lock acquisition**:

```go
philo.forks.right.Lock()
if ctx.Err() != nil {
    philo.forks.right.Unlock()  // Release fork
    return                       // Exit without recording meal
}
```

This guarantees no state mutations occur after cancellation.

### Memory Footprint: Mutexes vs. Channels

**Why I chose `sync.Mutex` over channels:**

| Aspect | sync.Mutex | Channel |
|--------|-----------|---------|
| **Memory** | ~8 bytes | ~96 bytes + buffer |
| **Conceptual Model** | Exclusive lock on stationary resource | Message passing between goroutines |
| **Operation** | Atomic CPU instruction (CAS) | Lock internal mutex + queue manipulation |
| **Overhead** | Minimal | Heavier (channel uses mutex internally!) |

**Source:** `hchan` struct in [runtime/chan.go](https://go.dev/src/runtime/chan.go)

**Trade-off Accepted:** I lost the "blocking safety" of channels and had to implement manual deadlock prevention (resource hierarchy). But for this problem, mutexes are the correct semantic abstraction.

## 🧪 Test Coverage

The test suite uses **table-driven integration tests** covering 11 concurrency scenarios:

### Correctness Tests
- ✅ **Standard Success** - Normal operation with meal quotas
- ✅ **Immediate Stop on Meals** - Clean exit when quota reached
- ✅ **One Philosopher Fails** - Impossible scenario (1 fork, 2 needed)

### Starvation Detection
- ✅ **Instant Death** - Mathematically impossible timing
- ✅ **The Fairness Trap** - Odd-numbered philosophers competing fairly
- ✅ **Odd Number Math Fail** - `time_to_die < 2×time_to_eat` edge case

### Timing Edge Cases
- ✅ **Minimal Survival Slack** - 210ms buffer catches scheduler jitter
- ✅ **Heavy Eating Contention** - Forks almost always occupied (eat/sleep = 10:1)
- ✅ **Long Sleep Lazy Philo** - Wake-up synchronization (eat/sleep = 1:12)

### Scalability
- ✅ **Large Scale Survival** - 200 philosophers, 5 meals each
- ✅ **High Contention Stress** - 199 philosophers to maximize race surface area

**Run with race detector:**

```bash
go test -race -v
```

The race detector instruments every memory access to catch unsynchronized shared state mutations.

## 📊 Performance Analysis

### Scheduler Introspection

Use `GODEBUG` to observe Go's M:P:N scheduler in action:

```bash
GODEBUG=schedtrace=1000 ./bin/philo 5 800 200 200 2
```

**Sample Output:**
```
SCHED 56313ms: gomaxprocs=8 idleprocs=8 threads=7 spinningthreads=0 
needspinning=0 idlethreads=5 runqueue=0 [0 0 0 0 0 0 0 0]
```

**Key Observations:**

- **Work-Stealing Proof:** `schedticks=[3 11521 2 11121 14361 10340 0 0]`
  - Some cores handled 14,000+ ticks while others handled zero
  - Idle cores "steal" work from busy cores automatically
  - No manual core affinity needed - scheduler balances load dynamically

- **Thread Efficiency:** `threads=7` for `gomaxprocs=8`
  - Go created only 7 OS threads to handle 5 philosophers + 1 monitor + 1 main
  - Demonstrates M:P:N multiplexing (many goroutines on few threads)

### Goroutine Lifecycle

- **Initial Stack:** 2KB (defined as `_StackMin` in [runtime/stack.go](https://go.dev/src/runtime/stack.go))
- **OS Thread Comparison:** pthread = ~512KB fixed stack
- **Creation Cost:** User-space allocation (no syscall)

## 🐛 Known Gotchas

### 1. Mutex Deadlock with Tight Timing

**Scenario:** `time_to_die = 410ms`, `time_to_eat = 200ms`, `time_to_sleep = 200ms`

**Expected:** 10ms buffer should suffice  
**Reality:** Program fails due to scheduler jitter

**Explanation:** Non-realtime OS adds ~1-11ms of context switching overhead. The "10ms buffer" is consumed by the scheduler itself.

**Solution:** Use `time_to_die ≥ (time_to_eat + time_to_sleep) + 200ms` for reliable results.

### 2. Preemptive Scheduling (Go 1.14+)

**Old Go (≤1.13):** Cooperative scheduling - goroutines yield only at safe points (channels, mutex, sleep)  
**Modern Go (≥1.14):** Preemptive - runtime's `sysmon` thread checks every ~10ms for long-running goroutines

**Implication:** Tight loops (like infinite `think()`) no longer freeze the monitor goroutine. The scheduler will preempt them at function call boundaries.

## 🛠️ Error Handling

The project demonstrates Go's idiomatic error handling patterns:

### Three Error Classes

1. **Parse Errors** - Malformed input from `strconv`
   ```
   ./philo 5 800 abc 200
   → "argument 3: 'abc': contains non-numeric characters"
   ```

2. **Domain Errors** - Valid integers violating problem constraints
   ```
   ./philo 0 800 200 200
   → "argument 1: '0': value must be strictly positive"
   ```

3. **System Errors** - Runtime issues (integer overflow, etc.)
   ```
   ./philo 5 999999999999999999 200 200
   → strconv.ErrRange
   ```

### Error Wrapping

Uses `%w` to preserve error identity while adding context:

```go
if errors.Is(err, strconv.ErrSyntax) {
    return fmt.Errorf("argument %d: '%s': contains non-numeric characters", i+1, arg)
}
```

This allows upstream callers to still detect the root cause with `errors.Is()`.

## 📚 Learning Resources

- [Go Concurrency Patterns](https://go.dev/blog/pipelines)
- [Effective Go - Concurrency](https://go.dev/doc/effective_go#concurrency)
- [Go Memory Model](https://go.dev/ref/mem)
- [Dining Philosophers Problem](https://en.wikipedia.org/wiki/Dining_philosophers_problem)

## 🎓 What I Learned

This was my **first serious Go program**, built in 4 days coming from a C background. Key takeaways:

1. **"Share memory by communicating"** isn't always the answer - channels add overhead when mutexes better model the problem
2. **Context propagation** is elegant for cancellation but requires careful handling with blocking primitives
3. **Go's scheduler is sophisticated** - work-stealing and preemption happen automatically
4. **The race detector is non-negotiable** - tests can pass 100 times and still hide races

## 📝 License

MIT

## 🙏 Acknowledgments

Built as part of learning Go's concurrency model. Inspired by the classic synchronization problem and modern systems programming challenges.

---

**Author:** Hien Nguyen  
**GitHub:** [novth17](https://github.com/novth17)  
**Project:** [dining-philosophers-go](https://github.com/novth17/dining-philosophers-go)
