// Code generated using goenumgen. DO NOT EDIT.
// This file was auto generated via github.com/yyle88/goenum
// Generated from: game_test.go:11 -> enums_test.TestGenerateGame
// ========== GOENUM:DO-NOT-EDIT-SECTION:END ==========

package enums

import (
	"slices"
)

type GameEnum string

const Game = GameEnum("게임")

func (GameEnum) S시작() GameEnum {
	return "게임" + "-" + "시작"
}
func (GameEnum) E종료() GameEnum {
	return "게임" + "-" + "종료"
}
func (GameEnum) P일시정지() GameEnum {
	return "게임" + "-" + "일시정지"
}
func (GameEnum) Basic() GameEnum {
	return Game
}
func (GameEnum) Enums() []GameEnum {
	return []GameEnum{
		Game.S시작(),
		Game.E종료(),
		Game.P일시정지(),
	}
}
func (value GameEnum) Valid() bool {
	return slices.Contains(Game.Enums(), value)
}
func (value GameEnum) Check() bool {
	return value == Game || slices.Contains(Game.Enums(), value)
}
