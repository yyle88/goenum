package goenumgen

import (
	"os"
	"reflect"

	"github.com/yyle88/formatgo"
	"github.com/yyle88/goenum/internal/utils"
	"github.com/yyle88/must"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/printgo"
	"github.com/yyle88/rese"
	"github.com/yyle88/syntaxgo"
)

type Config[T comparable] struct {
	Type       string // Type name. type xxx string. "xxx" is type name.
	Name       string // Elem name.
	BasicValue T      // Such as "Status" as basic string in "StatusOK" "StatusBadRequest" "StatusNotFound" "StatusCreated"
	DelimValue T      // You can get "status-ok" "status-wa", this "-" means delim
	Options    []*EnumOption[T]
	NamingMode NamingModeEnum // You can get "status-ok" "ok-status" "ok" with this mode
	IsGenBasic bool
	IsGenValid bool
	IsGenCheck bool
}

type EnumOption[T comparable] struct {
	Name        string // Enum func name
	OptionValue T      // Enum code text. such as 200 / 400 / 404 / 201 | "OK" / "WA"
}

func Generate[T comparable](config *Config[T], path string) {
	ptx := printgo.NewPTX()
	pkgName := syntaxgo.GetPkgName(osmustexist.FILE(path))
	ptx.Println("package", pkgName)
	ptx.Println("import (")
	ptx.Println(utils.WithQuotes("slices"))
	ptx.Println(")")
	ptx.Println(GenerateCode(config))
	code := ptx.Bytes()
	code = rese.A1(formatgo.FormatBytes(code))
	must.Done(os.WriteFile(path, code, 0644))
}

func GenerateCode[T comparable](config *Config[T]) string {
	ptx := printgo.NewPTX()
	ptx.Println("type", config.Type, reflect.TypeOf(config.BasicValue).String())
	ptx.Println("const", config.Name, "=", config.Type, "(", utils.CodeString(config.BasicValue), ")")
	for _, op := range config.Options {
		ptx.Println("func(", config.Type, ")", op.Name, "()", config.Type, "{")
		switch config.NamingMode {
		case NamingMode.Prefix():
			ptx.Println("return", utils.CodeString(op.OptionValue), "+", utils.CodeString(config.DelimValue), "+", utils.CodeString(config.BasicValue))
		case NamingMode.Suffix():
			ptx.Println("return", utils.CodeString(config.BasicValue), "+", utils.CodeString(config.DelimValue), "+", utils.CodeString(op.OptionValue))
		case NamingMode.Middle():
			ptx.Println("return", utils.CodeString(config.BasicValue), "+", utils.CodeString(op.OptionValue), "+", utils.CodeString(config.DelimValue))
		case NamingMode.Single():
			ptx.Println("return", utils.CodeString(op.OptionValue))
		case NamingMode: // use this value as default is also right
			ptx.Println("return", utils.CodeString(op.OptionValue))
		default:
			ptx.Println("return", utils.CodeString(op.OptionValue))
		}
		ptx.Println("}")
	}

	if config.IsGenBasic {
		ptx.Println("func (", config.Type, ") Basic() ", config.Type, " {")
		ptx.Println("return ", config.Name)
		ptx.Println("}")
	}

	ptx.Println("func(", config.Type, ")", "Enums()", "[]", config.Type, "{")
	ptx.Println("return []", config.Type, "{")
	for _, op := range config.Options {
		ptx.Println(config.Name, ".", op.Name, "()", ",")
	}
	ptx.Println("}")
	ptx.Println("}")

	if config.IsGenValid {
		ptx.Println("func (", "value", config.Type, ") Valid() bool {")
		ptx.Println("return slices.Contains(", config.Name, ".Enums(), value)")
		ptx.Println("}")
	}

	if config.IsGenCheck {
		ptx.Println("func (", "value", config.Type, ") Check() bool {")
		ptx.Println("return value == ", config.Name, " || slices.Contains(", config.Name, ".Enums(), value)")
		ptx.Println("}")
	}

	return ptx.String()
}
