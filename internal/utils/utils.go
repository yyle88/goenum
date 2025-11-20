// Package utils: Code generation string handling utilities
// Provides string formatting and quoting functions for code generation
// Supports type conversion and quote manipulation for generated code
//
// utils: 代码生成的字符串处理工具
// 提供代码生成的字符串格式化和引用函数
// 支持生成代码的类型转换和引号处理
package utils

import (
	"fmt"
	"net/url"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/yyle88/rese"
)

// CodeString converts any value to its Go code string representation
// Handles int and string types with proper formatting
// Returns formatted string suitable for code generation
//
// CodeString 将任何值转换为其 Go 代码字符串表示
// 处理 int 和 string 类型，具有适当的格式
// 返回适合代码生成的格式化字符串
func CodeString(v any) string {
	switch a := v.(type) {
	case int:
		return strconv.Itoa(a)
	case string:
		return WithQuotes(a)
	default:
		return fmt.Sprint(a)
	}
}

// WithQuotes wraps a string with double quotes
// Used to generate quoted string literals in code
// Returns quoted string for code insertion
//
// WithQuotes 用双引号包装字符串
// 用于生成代码中的引用字符串字面量
// 返回用于插入代码的引用字符串
func WithQuotes(s string) string {
	return `"` + s + `"`
}

// TrimQuotes removes surrounding double quotes from a string
// Used to extract raw string content from quoted literals
// Returns unquoted string content
//
// TrimQuotes 从字符串中移除周围的双引号
// 用于从引用字面量中提取原始字符串内容
// 返回未引用的字符串内容
func TrimQuotes(s string) string {
	return strings.Trim(s, `"`)
}

// GetGenPosFuncMark gets position information to trace code generation
// Returns formatted string showing source file and function that triggered generation
// Uses runtime.Caller to walk up stack using specified frame count
// Supports URL-encoded file paths
//
// GetGenPosFuncMark 获取调用位置信息，用于追踪代码生成
// 返回格式化字符串，显示触发生成的源文件和函数
// 使用 runtime.Caller 按指定帧数向上遍历调用栈
// 支持URL编码的文件路径
func GetGenPosFuncMark(skip int) string {
	pc, file, line, ok := runtime.Caller(skip + 1)
	if !ok {
		return ""
	}
	filename := filepath.Base(file)

	// Decode URL-encoded filename
	// 解码URL编码的文件名
	filename = rese.C1(url.PathUnescape(filename))

	funcInfo := runtime.FuncForPC(pc)
	if funcInfo == nil {
		return filename
	}
	funcName := filepath.Base(funcInfo.Name())
	return fmt.Sprintf("%s:%d -> %s", filename, line, funcName)
}
