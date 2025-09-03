[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/goenum/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/goenum/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/goenum)](https://pkg.go.dev/github.com/yyle88/goenum)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/goenum/main.svg)](https://coveralls.io/github/yyle88/goenum?branch=main)
![Supported Go Versions](https://img.shields.io/badge/Go-1.22%2C%201.23-lightgrey.svg)
[![GitHub Release](https://img.shields.io/github/release/yyle88/goenum.svg)](https://github.com/yyle88/goenum/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/goenum)](https://goreportcard.com/report/github.com/yyle88/goenum)

# goenum

Go 枚举生成和管理工具包，提供类型安全和灵活的命名模式。

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->
## 英文文档

[ENGLISH README](README.md)
<!-- TEMPLATE (ZH) END: LANGUAGE NAVIGATION -->

## 核心特性

🎯 **智能枚举生成**: 自动生成类型安全的枚举代码，支持可定制的命名模式  
⚡ **多种命名模式**: 支持前缀、后缀、中间和单一命名策略  
🔄 **类型安全**: 基于泛型约束的编译时枚举验证  
🌍 **灵活类型**: 支持任何可比较类型（int、string、自定义类型）  
📋 **验证函数**: 自动生成 Valid() 和 Check() 方法进行运行时验证

## 安装

```bash
go get github.com/yyle88/goenum
```

## 快速开始

### 1. 定义枚举配置

```go
package main

import (
    "github.com/yyle88/goenum/goenumgen"
)

func main() {
    // 配置枚举生成
    config := &goenumgen.Config[string]{
        Type:       "StatusEnum",
        Name:       "Status", 
        BasicValue: "Status",
        DelimValue: "-",
        NamingMode: goenumgen.NamingMode.Suffix(), // "Status-OK", "Status-Error"
        IsGenValid: true,
        IsGenCheck: true,
        Options: []*goenumgen.EnumOption[string]{
            {Name: "OK", OptionValue: "OK"},
            {Name: "Error", OptionValue: "Error"}, 
            {Name: "Pending", OptionValue: "Pending"},
        },
    }
    
    // 生成枚举代码
    goenumgen.Generate(config, "internal/enums/status.go")
}
```

### 2. 生成的枚举代码

上述配置生成以下代码：

```go
package enums

import "slices"

type StatusEnum string

const Status = StatusEnum("Status")

func (StatusEnum) OK() StatusEnum {
    return "Status" + "-" + "OK"
}

func (StatusEnum) Error() StatusEnum {
    return "Status" + "-" + "Error"
}

func (StatusEnum) Pending() StatusEnum {
    return "Status" + "-" + "Pending"
}

func (StatusEnum) Enums() []StatusEnum {
    return []StatusEnum{
        Status.OK(),
        Status.Error(), 
        Status.Pending(),
    }
}

func (value StatusEnum) Valid() bool {
    return slices.Contains(Status.Enums(), value)
}

func (value StatusEnum) Check() bool {
    return value == Status || slices.Contains(Status.Enums(), value)
}
```

### 3. 使用枚举

```go
package main

import (
    "fmt"
    "your-project/internal/enums"
    "github.com/yyle88/goenum"
)

func main() {
    // 创建枚举值
    status := enums.Status.OK()
    fmt.Println(status) // 输出: Status-OK
    
    // 验证枚举值
    if goenum.Valid(status) {
        fmt.Println("有效的枚举值")
    }
    
    // 支持基本值的检查
    if goenum.Check(enums.Status) {
        fmt.Println("基本值有效")
    }
    
    // 获取所有枚举值
    allStatuses := enums.Status.Enums()
    for _, s := range allStatuses {
        fmt.Printf("状态: %s, 有效: %t\n", s, s.Valid())
    }
}
```

## 命名模式

### 前缀模式
模式：`选项 + 分隔符 + 基本值`
```go
NamingMode: goenumgen.NamingMode.Prefix()
// 结果: "OK-Status", "Error-Status"
```

### 后缀模式  
模式：`基本值 + 分隔符 + 选项`
```go
NamingMode: goenumgen.NamingMode.Suffix()
// 结果: "Status-OK", "Status-Error"
```

### 中间模式
模式：`基本值 + 选项 + 分隔符`
```go
NamingMode: goenumgen.NamingMode.Middle()
// 结果: "StatusOK-", "StatusError-"
```

### 单一模式
模式：`选项`
```go
NamingMode: goenumgen.NamingMode.Single()
// 结果: "OK", "Error"
```

## 高级示例

### HTTP 状态码

```go
config := &goenumgen.Config[int]{
    Type:       "HTTPStatusEnum",
    Name:       "HTTPStatus",
    BasicValue: 0,
    DelimValue: 0, // 整数类型不使用分隔符
    NamingMode: goenumgen.NamingMode.Single(),
    IsGenValid: true,
    IsGenCheck: true,
    Options: []*goenumgen.EnumOption[int]{
        {Name: "OK", OptionValue: 200},
        {Name: "NotFound", OptionValue: 404},
        {Name: "InternalError", OptionValue: 500},
    },
}
```

### 数据库连接状态

```go
config := &goenumgen.Config[string]{
    Type:       "ConnStateEnum", 
    Name:       "ConnState",
    BasicValue: "conn",
    DelimValue: ".",
    NamingMode: goenumgen.NamingMode.Prefix(),
    IsGenBasic: true,
    IsGenValid: true,
    Options: []*goenumgen.EnumOption[string]{
        {Name: "Connected", OptionValue: "active"},
        {Name: "Disconnected", OptionValue: "inactive"},
        {Name: "Connecting", OptionValue: "pending"},
    },
}
// 生成: "active.conn", "inactive.conn", "pending.conn"
```

## 配置选项

| 字段 | 类型 | 描述 |
|------|------|------|
| `Type` | `string` | 生成的枚举类型名 |
| `Name` | `string` | 基础常量名 |  
| `BasicValue` | `T` | 枚举的基本值 |
| `DelimValue` | `T` | 复合名称的分隔符 |
| `Options` | `[]*EnumOption[T]` | 枚举选项定义 |
| `NamingMode` | `NamingModeEnum` | 命名模式策略 |
| `IsGenBasic` | `bool` | 生成 `Basic()` 方法 |
| `IsGenValid` | `bool` | 生成 `Valid()` 方法 | 
| `IsGenCheck` | `bool` | 生成 `Check()` 方法 |

## 验证函数

### `goenum.Valid()`
检查值是否存在于枚举集合中：
```go
if goenum.Valid(status) {
    // 值是定义的枚举选项之一
}
```

### `goenum.Check()`  
支持基本值回退的验证：
```go
if goenum.Check(status) {
    // 值是基本值或有效的枚举选项
}
```

## 项目结构

```
goenum/
├── goenum.go              # 主要验证函数
├── goenumgen/             # 代码生成包
│   ├── generate.go        # 生成引擎
│   └── naming_mode.go     # 命名模式定义
├── internal/
│   ├── constraint/        # 泛型类型约束
│   ├── utils/             # 工具函数
│   └── examples/          # 使用示例
│       ├── example1/      # 基本 int 枚举
│       ├── example2/      # 字符串枚举和验证
│       ├── example3/      # 开关模式枚举
│       └── example4/      # 复杂命名模式
└── README.md
```

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-08-28 08:33:43.829511 +0000 UTC -->

## 📄 许可证类型

MIT 许可证。详见 [LICENSE](LICENSE)。

---

## 🤝 项目贡献

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **发现问题？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **功能建议？** 创建 issue 讨论您的想法
- 📖 **文档疑惑？** 报告问题，帮助我们改进文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，帮助我们优化性能
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **意见反馈？** 欢迎所有建议和宝贵意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/repo-name.git`）
3. **导航**：进入克隆的项目（`cd repo-name`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：为面向用户的更改更新文档，并使用有意义的提交消息
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Pull Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Pull Request 和报告问题来为此项目做出贡献。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**使用这个包快乐编程！** 🎉

<!-- TEMPLATE (ZH) END: STANDARD PROJECT FOOTER -->

---

## GitHub 标星点赞

[![Stargazers](https://starchart.cc/yyle88/goenum.svg?variant=adaptive)](https://starchart.cc/yyle88/goenum)