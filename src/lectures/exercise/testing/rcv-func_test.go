//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//
//--Requirements:
//* Write unit tests that ensure:
//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests

package main

import "testing"

var test_units = []Player{
	{"Zakaria", Health{50, 150}, Energy{100, 150}},
}

func TestHeal(t *testing.T) {
	p := test_units[0]
	heal_units := []int{50, 100, -50, -100, 1000, -1000}
	for _, h := range heal_units {
		p.Heal(h)
		if p.Health.CurrentHealth > p.Health.MaxHealth {
			t.Errorf("Player overhealed %v, test case: %v", p.Health.CurrentHealth, h)
		}
		if p.Health.CurrentHealth < 0 {
			t.Errorf("Player health below 0 %v, test case: %v", p.Health.CurrentHealth, h)
		}
	}
}

func TestRest(t *testing.T) {
	p := test_units[0]
	heal_units := []int{50, 100, -50, -100, 1000, -1000}
	for _, h := range heal_units {
		p.Rest(h)
		if p.Energy.CurrentEnergy > p.Energy.MaxEnergy {
			t.Errorf("Player overdose %v, test case: %v", p.Energy.CurrentEnergy, h)
		}
		if p.Energy.CurrentEnergy < 0 {
			t.Errorf("Energy insufficiant %v, test case: %v", p.Energy.CurrentEnergy, h)
		}
	}
}
