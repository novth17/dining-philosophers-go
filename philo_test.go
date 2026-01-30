package main 
import "testing"

func TestSimulation(t *testing.T) {
    cases := []struct {
        name       string
        philos     int
        dieTime    int
        eatTime    int
        expectDeath bool
    }{
        {"Standard 5", 5, 800, 200, false},
        {"Instant Starve", 5, 10, 200, true}, 
        {"Single Philo", 1, 800, 200, true}, 
    }
    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
        })
    }
}