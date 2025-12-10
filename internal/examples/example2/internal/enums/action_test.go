package enums_test

import (
	"testing"

	"github.com/yyle88/goenum/goenumgen"
	"github.com/yyle88/runpath/runtestpath"
)

func TestGenerateAction(t *testing.T) {
	goenumgen.Generate(&goenumgen.Config[string]{
		Type:       "ActionEnum",
		Name:       "Action",
		BasicValue: "Action",
		DelimValue: "-",
		Options: []*goenumgen.EnumOption[string]{
			{
				Name:        "Start",
				OptionValue: "Start",
			},
			{
				Name:        "Close",
				OptionValue: "Close",
			},
			{
				Name:        "Pause",
				OptionValue: "Pause",
			},
			{
				Name:        "Reset",
				OptionValue: "Reset",
			},
		},
		NamingMode: goenumgen.NamingMode.Suffix(),
	}, runtestpath.SrcPath(t))
}
