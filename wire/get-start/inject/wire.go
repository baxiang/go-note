//+build wireinject

package main

import "github.com/google/wire"

func InitMission(f FightParam,m MonsterParam)Mission{
    wire.Build(NewFighter,NewMonster,NewMission)
	return Mission{}
}