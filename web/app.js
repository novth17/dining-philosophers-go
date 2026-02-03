const container = document.getElementById("philosophers");
const log = document.getElementById("log");

const NUM_PHILOS = 5;
const philosophers = {};

// for (let i = 1; i <= NUM_PHILOS; i++) {
//   const el = document.createElement("div");
//   el.className = "philo thinking";
//   el.innerText = `Philo ${i}`;
//   container.appendChild(el);

//   philosophers[i] = el;
// }

const philos = {};
const philosDiv = document.getElementById("philosophers");

function getPhilo(id) {
  if (!philos[id]) {
    const div = document.createElement("div");
    div.className = "philo thinking";
    div.textContent = id;
    philos[id] = div;
    philosDiv.appendChild(div);
  }
  return philos[id];
}

function setState(id, state, time) {
  const p = getPhilo(id);

  p.className = `philo ${state}`;

  if (time !== undefined) {
    p.textContent = `Philo${id} ${time}ms`;
    p.title = `t=${time}ms`;
  } else {
    p.textContent = `${id} (no time)`;
  }

  console.log("setState", id, state, time);
}

const source = new EventSource("/events");

source.onmessage = (e) => {
  const ev = JSON.parse(e.data);

  if (ev.state && ev.philo !== undefined) {
    setState(ev.philo, ev.state, ev.time);
  }

  if (ev.event === "death") {
    setState(ev.philo, "dead", ev.time);
  }
};
