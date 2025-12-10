// Package goenumgen: Go enum code generation engine with flexible naming modes
// Auto generates type safe enum code with customizable naming patterns
// Returns enum.Enums collection for validation and lookup operations
// Provides compile-time safe checks with runtime performance optimization
//
// goenumgen: Go 枚举代码生成引擎，支持灵活的命名模式
// 自动生成类型安全的枚举代码，支持可定制的命名模式
// 返回 enum.Enums 集合用于验证和查找操作
// 提供编译时安全检查和运行时性能优化
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
// Contains all settings needed to generate type safe enum code
// Supports generic types and flexible naming pattern configuration
//
// Config 代表枚举生成的完整配置
// 包含生成类型安全枚举代码所需的所有设置
// 支持泛型类型和灵活的命名模式配置
type Config[T comparable] struct {
	Type       string           // Type name for generated enum // 生成枚举的类型名
	Name       string           // Base enum element name // 基本枚举元素名
	BasicValue T                // Basic string like "Status" in "StatusOK" "StatusNotFound" // 基本字符串如 "StatusOK" "StatusNotFound" 中的 "Status"
	DelimValue T                // Delimiter like "-" to get "status-ok" "status-wa" // 分隔符如 "-" 用于生成 "status-ok" "status-wa"
	Options    []*EnumOption[T] // Enum option definitions // 枚举选项定义
	NamingMode NamingModeEnum   // Naming pattern mode // 命名模式
}

// EnumOption represents a single enum value definition
// Contains the function name and corresponding value for code generation
// Supports comparable types as enum value
//
// EnumOption 代表单个枚举值定义
// 包含生成的函数名和对应的值
// 支持可比较类型作为枚举值
type EnumOption[T comparable] struct {
	Name        string // Generated enum function name // 生成的枚举函数名
	OptionValue T      // Enum code like 200/400/404/201 or "OK"/"WA" // 枚举代码如 200/400/404/201 或 "OK"/"WA"
}

// Generate creates enum code file with complete package structure
// Auto detects package name and imports required dependencies
// Formats generated code and writes to specified file path
// Ensures file permissions and handles potential issues
//
// Generate 创建具有完整包结构的枚举代码文件
// 自动检测包名并导入所需依赖
// 格式化生成的代码并写入指定的文件路径
// 确保文件权限并处理潜在问题
func Generate[T comparable](config *Config[T], path string) {
	// Create print context for code generation
	// 创建代码生成的打印上下文
	ptx := printgo.NewPTX()

	// Add generation header comments
	// 添加生成头部注释
	ptx.Println("// Code generated using goenumgen. DO NOT EDIT.")
	ptx.Println("// This file was auto generated via github.com/yyle88/goenum")
	if genPos := utils.GetGenPosFuncMark(1); genPos != "" {
		ptx.Println("// Generated from:", genPos)
	}
	ptx.Println("// ========== GOENUM:DO-NOT-EDIT-SECTION:END ==========")
	ptx.Println()

	pkgName := syntaxgo.GetPkgName(osmustexist.FILE(path))
	ptx.Println("package", pkgName)
	ptx.Println("import (")
	ptx.Println(utils.WithQuotes("github.com/yyle88/goenum"))
	ptx.Println(")")
	ptx.Println(GenerateCode(config))
	// Format and write generated code to file
	// 格式化并将生成的代码写入文件
	code := ptx.Bytes()
	code = rese.A1(formatgo.FormatBytes(code))
	must.Done(os.WriteFile(path, code, 0644))
}

// GenerateCode creates the actual enum implementation code
// Builds type definitions, constants, and enum collection method
// Supports multiple naming modes for value generation
// Returns formatted Go code for compilation
//
// GenerateCode 创建实际的枚举实现代码
// 构建类型定义、常量和枚举集合方法
// 支持多种命名模式用于值生成
// 返回格式化的 Go 代码用于编译
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

	// Generate Enums() method returning *enum.Enums collection
	// 生成 Enums() 方法，返回 *enum.Enums 集合
	ptx.Println("func(", config.Type, ")", "Enums()", "*goenum.Enums[", config.Type, "]", "{")
	ptx.Println("return goenum.NewEnums(")
	for _, op := range config.Options {
		ptx.Println("goenum.NewEnum(", config.Name, ".", op.Name, "()", "),")
	}
	ptx.Println(")")
	ptx.Println("}")

	return ptx.String()
}
