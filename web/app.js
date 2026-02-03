const container = document.getElementById("philosophers")
const log = document.getElementById("log")

const NUM_PHILOS = 5
const philosophers = {}

for (let i = 1; i <= NUM_PHILOS; i++) {
  const el = document.createElement("div")
  el.className = "philo thinking"
  el.innerText = `Philo ${i}`
  container.appendChild(el)

  philosophers[i] = el
}

function setState(id, state) {
  const el = philosophers[id]
  if (!el) return

  el.className = `philo ${state}`
}

const source = new EventSource("/events")

source.onmessage = (e) => {
  const event = JSON.parse(e.data)

  if (event.philo && event.state) {
    setState(event.philo, event.state)
  }

  log.textContent += JSON.stringify(event) + "\n"
  log.scrollTop = log.scrollHeight
}