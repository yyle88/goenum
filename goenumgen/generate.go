// Package goenumgen: Go enum code generation engine with flexible naming modes
// Auto generates type-safe enum code with customizable naming patterns
// Supports multiple output formats and validation function generation
// Provides compile-time safety with runtime efficiency optimization
//
// goenumgen: Go 枚举代码生成引擎，支持灵活的命名模式
// 自动生成类型安全的枚举代码，支持可定制的命名模式
// 支持多种输出格式和验证函数生成
// 提供编译时安全和运行时效率优化
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

// Config represents the complete configuration for enum generation
// Contains all settings needed to generate type-safe enum code
// Supports generic types and flexible naming pattern configuration
// Enables customization of generated validation functions
//
// Config 代表枚举生成的完整配置
// 包含生成类型安全枚举代码所需的所有设置
// 支持泛型类型和灵活的命名模式配置
// 允许自定义生成的验证函数
type Config[T comparable] struct {
	Type       string           // Type name for generated enum // 生成枚举的类型名
	Name       string           // Base enum element name // 基本枚举元素名
	BasicValue T                // Basic string like "Status" in "StatusOK" "StatusNotFound" // 基本字符串如 "StatusOK" "StatusNotFound" 中的 "Status"
	DelimValue T                // Delimiter like "-" to get "status-ok" "status-wa" // 分隔符如 "-" 用于生成 "status-ok" "status-wa"
	Options    []*EnumOption[T] // Enum option definitions // 枚举选项定义
	NamingMode NamingModeEnum   // Naming pattern mode // 命名模式
	IsGenBasic bool             // Generate Basic() method // 生成 Basic() 方法
	IsGenValid bool             // Generate Valid() method // 生成 Valid() 方法
	IsGenCheck bool             // Generate Check() method // 生成 Check() 方法
}

// EnumOption represents a single enum value definition
// Contains the function name and corresponding value for generation
// Supports any comparable type as enum value
//
// EnumOption 代表单个枚举值定义
// 包含生成的函数名和对应的值
// 支持任何可比较类型作为枚举值
type EnumOption[T comparable] struct {
	Name        string // Generated enum function name // 生成的枚举函数名
	OptionValue T      // Enum code like 200/400/404/201 or "OK"/"WA" // 枚举代码如 200/400/404/201 或 "OK"/"WA"
}

// Generate creates enum code file with complete package structure
// Auto detects package name and imports required dependencies
// Formats generated code and writes to specified file path
// Ensures proper file permissions and error handling
//
// Generate 创建具有完整包结构的枚举代码文件
// 自动检测包名并导入所需依赖
// 格式化生成的代码并写入指定的文件路径
// 确保适当的文件权限和错误处理
func Generate[T comparable](config *Config[T], path string) {
	// Create print context for code generation
	// 创建代码生成的打印上下文
	ptx := printgo.NewPTX()
	pkgName := syntaxgo.GetPkgName(osmustexist.FILE(path))
	ptx.Println("package", pkgName)
	ptx.Println("import (")
	ptx.Println(utils.WithQuotes("slices"))
	ptx.Println(")")
	ptx.Println(GenerateCode(config))
	// Format and write generated code to file
	// 格式化并将生成的代码写入文件
	code := ptx.Bytes()
	code = rese.A1(formatgo.FormatBytes(code))
	must.Done(os.WriteFile(path, code, 0644))
}

// GenerateCode creates the actual enum implementation code
// Builds type definitions, constants, and validation methods
// Supports multiple naming modes and optional method generation
// Returns formatted Go code ready for compilation
//
// GenerateCode 创建实际的枚举实现代码
// 构建类型定义、常量和验证方法
// 支持多种命名模式和可选方法生成
// 返回格式化的 Go 代码，准备编译
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
