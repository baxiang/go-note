package main

import "fmt"

type FightParam string
type MonsterParam string


type Monster struct {
	Name string
}

func NewMonster(name MonsterParam) Monster {
	return Monster{Name: string(name)}
}

type Fighter struct {
	Name string
}

func NewFighter(name FightParam) Fighter {
	return Fighter{Name: string(name)}
}

type Mission struct {
	Fighter Fighter
	Monster Monster
}

func NewMission(f Fighter, m Monster) Mission {
	return Mission{
		Fighter: f,
		Monster: m,
	}
}

func (m Mission) start() {
	fmt.Printf("%s defeates %s, world peace \n", m.Fighter.Name, m.Monster.Name)
}

func main() {
	mission := InitMission(FightParam("kitty"),MonsterParam("Ultraman"))
	mission.start()
}