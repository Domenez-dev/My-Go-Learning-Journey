//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type Health struct {
	CurrentHealth int
	MaxHealth     int
}

type Energy struct {
	CurrentEnergy int
	MaxEnergy     int
}

type Player struct {
	Name   string
	Health Health
	Energy Energy
}

func (p *Player) Heal(value int) {
	new_health := p.Health.CurrentHealth + value
	if new_health > p.Health.MaxHealth {
		p.Health.CurrentHealth = p.Health.MaxHealth
	} else {
		p.Health.CurrentHealth = new_health
	}
}

func (p *Player) Rest(value int) {
	new_energy := p.Energy.CurrentEnergy + value
	if new_energy > p.Energy.MaxEnergy {
		p.Energy.CurrentEnergy = p.Energy.MaxEnergy
	} else {
		p.Energy.CurrentEnergy = new_energy
	}
}

var p = fmt.Printf

func main() {
	player := Player{"Zakaria", Health{50, 150}, Energy{100, 150}}

	p("old Health: %v\n", player.Health.CurrentHealth)
	player.Heal(50)
	p("new Health: %v\n", player.Health.CurrentHealth)

	p("old Energy: %v\n", player.Energy.CurrentEnergy)
	player.Rest(100)
	p("new Energy: %v\n", player.Energy.CurrentEnergy)
}
