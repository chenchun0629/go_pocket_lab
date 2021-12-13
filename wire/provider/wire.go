//+build wireinject

package main

import "github.com/google/wire"

// ProviderSet
var monsterPlayerSet = wire.NewSet(NewMonster, NewPlayer)

// 结构构造器
var endingASet = wire.NewSet(monsterPlayerSet, wire.Struct(new(EndingA), "Player", "Monster"))
var endingBSet = wire.NewSet(monsterPlayerSet, wire.Struct(new(EndingB), "Player", "Monster"))

// 绑定值
var kitty = Monster{Name: "kitty"}

// 清理函数
func InitMission(name string) (Mission, func(), error) {
	wire.Build(NewMonster, NewPlayer, NewMission)
	return Mission{}, nil, nil
}

// 结构字段作为构造器
func InitPlayer() Player {
	wire.Build(NewMission_A, wire.FieldsOf(new(Mission), "Player"))
	return Player{}
}

func InitMonster() Monster {
	wire.Build(NewMission_A, wire.FieldsOf(new(Mission), "Monster"))
	return Monster{}
}

// 绑定值
func InitEndingA(name string) EndingA {
	wire.Build(NewPlayer, wire.Value(kitty), NewEndingA)
	return EndingA{}
}

func InitEndingB(name string) EndingB {
	wire.Build(NewPlayer, wire.Value(kitty), NewEndingB)
	return EndingB{}
}

// 结构构造器
//func InitEndingA(name string) EndingA {
//	wire.Build(endingASet)
//	return EndingA{}
//}
//
//func InitEndingB(name string) EndingB {
//	wire.Build(endingBSet)
//	return EndingB{}
//}

// ProviderSet
//func InitEndingA(name string) EndingA {
//	wire.Build(monsterPlayerSet, NewEndingA)
//	return EndingA{}
//}
//
//func InitEndingB(name string) EndingB {
//	wire.Build(monsterPlayerSet, NewEndingB)
//	return EndingB{}
//}
