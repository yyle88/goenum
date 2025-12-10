package enums_test

import (
	"testing"

	"github.com/yyle88/goenum/goenumgen"
	"github.com/yyle88/runpath/runtestpath"
)

func TestGenerateConnection(t *testing.T) {
	goenumgen.Generate(&goenumgen.Config[string]{
		Type:       "ConnectionEnum",
		Name:       "Connection",
		BasicValue: "接続",
		DelimValue: "",
		Options: []*goenumgen.EnumOption[string]{
			{Name: "C接続", OptionValue: "接続"},
			{Name: "D切断", OptionValue: "切断"},
			{Name: "W待機", OptionValue: "待機"},
		},
		NamingMode: goenumgen.NamingMode.Suffix(),
	}, runtestpath.SrcPath(t))
}
