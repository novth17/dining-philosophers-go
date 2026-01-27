package main

import "sync"

func  (philo *Philo) routine() {
	for {
		philo.eat()
		//philo.sleep()
		//philo.think()

	}
}

//eat
func (philo *Philo) eat() {

	program := philo.prog

	left := philo.id - 1
	right := philo.id % program.numPhilos

	var mtxFirst, mtxSecond *sync.Mutex //pointer to a sync.Mutex

	if (philo.id % 2 == 0) {
		mtxFirst = &program.forks[left]
		mtxSecond = &program.forks[right]
	} else {
		mtxFirst = &program.forks[right]
		mtxSecond = &program.forks[left]
	}

	mtxFirst.Lock()
	program.printMtx(philo.id, " has taken a fork")
	
	
	mtxSecond.Lock()
	program.printMtx(philo.id, " has taken a fork")
	
	mtxFirst.Unlock()
	mtxSecond.Unlock()
}

//think


//sleep

//take fork