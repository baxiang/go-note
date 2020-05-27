package main

import "fmt"

type Monster struct {
	Name string
}

func NewMonster(name string) Monster {
	return Monster{Name: name}
}

type Fighter struct {
	Name string
}

func NewFighter(name string) Fighter {
	return Fighter{Name: name}
}

type Mission struct {
	Fighter Fighter
	Monster Monster
}

func NewMission(p Fighter, m Monster) Mission {
	return Mission{
		Fighter: p,
		Monster: m,
	}
}

func (m Mission) start() {
	fmt.Printf("%s defeates %s, world peace \n", m.Fighter.Name, m.Monster.Name)
}

func main() {
	monster := NewMonster("kitty")
	fighter := NewFighter("Ultraman")
	mission := NewMission(fighter, monster)
	mission.start()
}
