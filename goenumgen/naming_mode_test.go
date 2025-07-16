package goenumgen_test

import (
	"testing"

	"github.com/yyle88/goenum/goenumgen"
	"github.com/yyle88/runpath/runtestpath"
)

func TestGenerateAction(t *testing.T) {
	goenumgen.Generate(&goenumgen.Config[string]{
		Type:       "NamingModeEnum",
		Name:       "NamingMode",
		BasicValue: "Naming",
		DelimValue: "-",
		Options: []*goenumgen.EnumOption[string]{
			{
				Name:        "Prefix",
				OptionValue: "Prefix",
			},
			{
				Name:        "Suffix",
				OptionValue: "Suffix",
			},
			{
				Name:        "Middle",
				OptionValue: "Middle",
			},
			{
				Name:        "Single",
				OptionValue: "Single",
			},
		},
		NamingMode: goenumgen.NamingMode.Suffix(),
		IsGenBasic: false,
		IsGenValid: false,
		IsGenCheck: false,
	}, runtestpath.SrcPath(t))
}
