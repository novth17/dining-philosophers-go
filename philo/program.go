package main

import (
	"sync"
	"time"
)

const (
	THINKING int32 = iota
	EATING
	SLEEPING
	DEAD
)

type Philo struct {
	id           int
	timeLastMeal time.Time
	mealCount    int
	leftFork     *sync.Mutex
	rightFork    *sync.Mutex
	prog         *Program
	State        int32 // Atomic state for visualizer
}

type Program struct {
	numPhilos     int
	timeDie       time.Duration
	timeEat       time.Duration
	timeSleep     time.Duration
	mealsRequired int
	startTime     time.Time
	stopSim       int32

	philos []Philo
	forks  []sync.Mutex

	logs  []string
	logMu sync.Mutex
	mealMu sync.Mutex
	wg    sync.WaitGroup
}

func NewProgram(args []string) (*Program, error) {
	var program Program
	err := initProgram(&program, args)
	if err != nil {
		return nil, err
	}
	return &program, nil
}