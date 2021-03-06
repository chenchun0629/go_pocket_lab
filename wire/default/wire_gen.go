// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

// Injectors from wire.go:

func InitMission(name string) (Mission, error) {
	player, err := NewPlayer(name)
	if err != nil {
		return Mission{}, err
	}
	monster := NewMonster()
	mission := NewMission(player, monster)
	return mission, nil
}
