// ./philo numPhilos timeDie timeEat timeSleep [mealsRequired]
package main

import (
	"sync"
	"time"
)
type Philo struct {
	id int
	prog *Program
	timeLastMeal time.Time
	mealCount int //need mutex or atomic
}

type Program struct {
	
	numPhilos int
	timeDie time.Duration
	timeEat time.Duration
	timeSleep time.Duration
	mealsRequired int 

	startTime time.Time
	stopSim bool

	// 	slice header
	//  ├─ data ─────► [ Philo ][ Philo ][ Philo ]
	//  ├─ len = 3
	//  └─ cap = 3
	philos []Philo //slice of struct Philo
	forks []sync.Mutex //slice of struct of fork mutexes

	logMutex sync.Mutex
	mealMutex sync.Mutex
	wg         sync.WaitGroup
	
	//ctx context.Context // this is an interface, a shared signal says "program must stop now"
}

//this function assigns each args to the equivalent in the struct
func NewProgram(args []string) (*Program, error) {

	var program Program
	
	err := initProgram(&program, args)
	if (err != nil) {
		return nil, err
	}
	return &program, nil
}