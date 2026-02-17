package main

import (
	"testing"
	"time"
)

func TestPhiloStarvation(t *testing.T) {
	tests := []struct {
		name          string
		args          []string
		shouldSurvive bool
	}{
		{
            name:          "Standard Success",
            args:          []string{"5", "800", "200", "200", "7"},
            shouldSurvive: true,
        },
        {
            name:          "Instant Death",
            args:          []string{"5", "5", "200", "200"},
            shouldSurvive: false,
        },
        {
            name:          "One Philosopher Fails",
            args:          []string{"1", "800", "200", "200"},
            shouldSurvive: false,
        },
        {
            name:          "Large Scale Survival",
            args:          []string{"200", "800", "200", "200", "5"},
            shouldSurvive: true,
        },
		{
			name:          "Potential Deadlock",
			args:          []string{"100", "800", "200", "200"},
			shouldSurvive: true,
		},
        {
            name:          "Minimal Survival Slack",
            // 410ms life, 200ms eat, 200ms sleep. 
            // Only 10ms of "thinking" time allowed. Catches slow mutex logic.
            args:          []string{"5", "610", "200", "200"},
            shouldSurvive: true, 
        },
        {
            name:          "The Fairness Trap (Odd Number)",
            // 3 philos. If 1 and 2 keep eating, 3 will starve. 
            // This tests if fork-picking logic allows everyone a turn.
            args:          []string{"3", "610", "200", "200"},
            shouldSurvive: true,
        },
        {
            name:          "Immediate Stop on Meals",
            // Everyone only eats once. Program should exit immediately.
            args:          []string{"5", "800", "200", "200", "1"},
            shouldSurvive: true,
        },
        {
            name:          "High Contention Stress",
            // 199 is a lot of goroutines fighting for forks. 
            // Best case for catching Data Races with go test -race.
            args:          []string{"199", "610", "200", "200"},
            shouldSurvive: true,
        },
		{
			// Scenario: Eat is much longer than sleep.
			// Forks are almost always occupied.
			name:          "Heavy Eating Contention",
			args: []string{"3", "610", "300", "30"},
			shouldSurvive: false,
		},
		{
			// Scenario: Sleep is much longer than eat.
			// Tests if philos wake up and sync-clash for forks.
			name:          "Long Sleep Lazy Philo",
			args:          []string{"5", "800", "50", "600"},
			shouldSurvive: true,
		},
		{
			// Scenario: The "3x" Rule test. 
			// If die < 2*eat, an odd number will likely fail.
			name:          "Odd Number Math Fail",
			args:          []string{"3", "390", "200", "50"},
			shouldSurvive: false,
		},

		{
			name:          "Stress test very close number",
			args:          []string{"6", "320", "100", "200", "50"},
			shouldSurvive: true,
		},
		{
            name:          "Large Number Of Philo Stress",
			// a very brave move 
            args:          []string{"10000", "800", "200", "200", "1"},
            shouldSurvive: true,
        },
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			program, err := NewProgram(testcase.args)
			if err != nil {
				t.Fatalf("init failed: %v", err)
			}

			done := make(chan struct{})

			go func() {
				program.Run()
				close(done)
			}()

			// ---- CASE 1: Meals required → wait for completion ----
			if program.mealsRequired != -1 {
				select {
				case <-done:
					// OK, simulation finished naturally
				case <-time.After(60 * time.Second):
					program.Stop()
					t.Fatal("simulation did not finish in time")
				}

				if !testcase.shouldSurvive {
					t.Fatal("expected death, but meals were completed")
				}

				for _, p := range program.philos {
					if p.mealCount < program.mealsRequired {
						t.Fatalf(
							"philo %d only ate %d/%d meals",
							p.id,
							p.mealCount,
							program.mealsRequired,
						)
					}
				}

				return
			}

			// ---- CASE 2: No meals → expect death - program.mealsRequired == -1 ----
			//The simulation should only stop if someone dies.
			survivalWindow := time.Duration(program.timeDie * 3) * time.Millisecond

			if survivalWindow > 5 * time.Second {
    			survivalWindow = 5 * time.Second
			}

			select {
			case <-done:
				if testcase.shouldSurvive {
					t.Fatal("simulation ended early but should survive")
				}
			case <-time.After(survivalWindow):
				program.Stop()
				if !testcase.shouldSurvive {
					t.Fatal("expected death, but simulation kept running")
				}
			}

		})
	}
}