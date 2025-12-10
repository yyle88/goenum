// Code generated using goenumgen. DO NOT EDIT.
// This file was auto generated via github.com/yyle88/goenum
// Generated from: game_test.go:11 -> enums_test.TestGenerateGame
// ========== GOENUM:DO-NOT-EDIT-SECTION:END ==========

package enums

import (
	"github.com/yyle88/goenum"
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
func (GameEnum) Enums() *goenum.Enums[GameEnum] {
	return goenum.NewEnums(
		goenum.NewEnum(Game.S시작()),
		goenum.NewEnum(Game.E종료()),
		goenum.NewEnum(Game.P일시정지()),
	)
}
