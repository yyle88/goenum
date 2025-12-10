package enums_test

import (
	"testing"

	"github.com/yyle88/goenum/goenumgen"
	"github.com/yyle88/runpath/runtestpath"
)

func TestGenerateGame(t *testing.T) {
	goenumgen.Generate(&goenumgen.Config[string]{
		Type:       "GameEnum",
		Name:       "Game",
		BasicValue: "게임",
		DelimValue: "-",
		Options: []*goenumgen.EnumOption[string]{
			{Name: "S시작", OptionValue: "시작"},
			{Name: "E종료", OptionValue: "종료"},
			{Name: "P일시정지", OptionValue: "일시정지"},
		},
		NamingMode: goenumgen.NamingMode.Suffix(),
	}, runtestpath.SrcPath(t))
}
