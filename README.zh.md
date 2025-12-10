[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/goenum/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/goenum/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/goenum)](https://pkg.go.dev/github.com/yyle88/goenum)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/goenum/main.svg)](https://coveralls.io/github/yyle88/goenum?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.25+-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/yyle88/goenum.svg)](https://github.com/yyle88/goenum/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/goenum)](https://goreportcard.com/report/github.com/yyle88/goenum)

# GOENUM

Go 枚举代码生成工具，让不同业务领域可以共享 OK、ERROR、PENDING 等常用枚举名称，通过命名空间隔离避免命名冲突。

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->

## 英文文档

[ENGLISH README](README.md)
<!-- TEMPLATE (ZH) END: LANGUAGE NAVIGATION -->

## 功能特性

🔒 **命名空间隔离** - 每个领域拥有独立的枚举空间，避免命名冲突
⚡ **枚举集合** - 生成的 `Enums()` 返回集合，支持 `Lookup`、`List`、`Get` 等方法
🎯 **简洁代码** - 直观语法匹配业务逻辑模式
✅ **编译保护** - 在构建时捕获枚举误用，而非运行时
🌍 **多种语言** - 支持使用多种语言字符生成枚举

## 安装

```bash
go get github.com/yyle88/goenum
```

## 使用方法

Go 缺乏真正的枚举命名空间。不同领域无法共享 `OK`、`ERROR`、`PENDING` 等通用值名称。

### 传统方式：需要冗长前缀

```go
type PackageStatus string
const (
    PackagePending   PackageStatus = "PENDING"
    PackageConfirmed PackageStatus = "CONFIRMED"
    PackageShipped   PackageStatus = "SHIPPED"
    PackageDelivered PackageStatus = "DELIVERED"
)

type PaymentStatus string
const (
    PaymentPending PaymentStatus = "PENDING"
    PaymentFailed  PaymentStatus = "FAILED"
    PaymentSuccess PaymentStatus = "SUCCESS"
    PaymentRefund  PaymentStatus = "REFUND"
)

// 冗长的 switch 语句，带有长前缀
func processPackage(status string) {
    switch PackageStatus(status) {
    case PackagePending:
        // 处理待处理
    case PackageConfirmed:
        // 处理已确认
    case PackageShipped:
        // 处理已发货
    case PackageDelivered:
        // 处理已交付
    }
}

func processPayment(status string) {
    switch PaymentStatus(status) {
    case PaymentPending:
        // 处理待支付
    case PaymentFailed:
        // 处理支付失败
    case PaymentSuccess:
        // 处理支付成功
    case PaymentRefund:
        // 处理退款
    }
}
```

### 使用 GOENUM：清晰的命名空间方法

```go
// 每个领域拥有自己的清晰命名空间
func processPackage(status string) {
    switch PackageStatusEnum(status) {
    case PackageStatus.Pending():
        // 处理待处理
    case PackageStatus.Confirmed():
        // 处理已确认
    case PackageStatus.Shipped():
        // 处理已发货
    case PackageStatus.Delivered():
        // 处理已交付
    }
}

func processPayment(status string) {
    switch PaymentStatusEnum(status) {
    case PaymentStatus.Pending():
        // 处理待支付
    case PaymentStatus.Failed():
        // 处理失败
    case PaymentStatus.Success():
        // 处理成功
    case PaymentStatus.Refund():
        // 处理退款
    }
}
```

### 使用 Enums() 验证枚举

每个生成的枚举类型都有 `Enums()` 方法，返回集合用于验证和查找：

```go
// 验证枚举值
if _, ok := PackageStatus.Enums().Lookup(status); !ok {
    return errors.New("无效的包裹状态")
}

// 获取所有有效枚举值
allStatuses := PackageStatus.Enums().List()
```

## 核心优势

🔒 **真正隔离** - `PackageStatus.Pending()` 和 `PaymentStatus.Pending()` 是完全不同的类型
⚡ **集合查询** - 通过 `Enums().Lookup()` 验证枚举值是否有效
🎯 **业务清晰** - 代码读起来像自然的业务语言
✅ **编译时安全** - 不可能混用不同领域的枚举

## 多语言支持
GOENUM 支持使用多种语言生成枚举：

```go
// 简体中文枚举示例
func processTask(status string) {
    taskStatus := TaskStatusEnum(status)
    switch taskStatus {
    case TaskStatus.C待处理():
        // handle pending task
    case TaskStatus.C已确认():
        // handle confirmed task
    case TaskStatus.C进行中():
        // handle active task
    case TaskStatus.C已完成():
        // handle completed task
    }
}
```

```go
// 繁体中文枚举示例
func processPermission(status string) {
    permStatus := PermissionStatusEnum(status)
    switch permStatus {
    case PermissionStatus.C開啟():
        // handle enabled permission
    case PermissionStatus.C關閉():
        // handle disabled permission
    }
}
```

```go
// 日文枚举示例
func processConnection(status string) {
    connStatus := ConnectionStatusEnum(status)
    switch connStatus {
    case ConnectionStatus.C接続():
        // handle connected
    case ConnectionStatus.C切断():
        // handle disconnected
    case ConnectionStatus.C待機():
        // handle waiting
    }
}
```

```go
// 韩语枚举示例
func processGame(status string) {
    gameStatus := GameStatusEnum(status)
    switch gameStatus {
    case GameStatus.C시작():
        // handle game start
    case GameStatus.C종료():
        // handle game end
    case GameStatus.C일시정지():
        // handle game pause
    }
}
```

---

**示例**: 查看 [examples](internal/examples)

---

## 关联项目

- **[enum](https://github.com/yyle88/enum)** - Go 枚举集合管理，支持元数据和基于 map 的查找
- **[goenum](https://github.com/yyle88/goenum)** - Go 枚举代码生成，支持命名空间隔离（本项目）
- **[protoenum](https://github.com/go-xlan/protoenum)** - Protocol Buffers 枚举代码生成，支持类型安全操作

---

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-25 03:52:28.131064 +0000 UTC -->

## 📄 许可证类型

MIT 许可证 - 详见 [LICENSE](LICENSE)。

---

## 💬 联系与反馈

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **问题报告？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **新颖思路？** 创建 issue 讨论
- 📖 **文档疑惑？** 报告问题，帮助我们完善文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，协助解决性能问题
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **反馈意见？** 欢迎提出建议和意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/repo-name.git`）
3. **导航**：进入克隆的项目（`cd repo-name`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：面向用户的更改需要更新文档
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Merge Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Merge Request 和报告问题来贡献此项目。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**祝你用这个包编程愉快！** 🎉🎉🎉

<!-- TEMPLATE (ZH) END: STANDARD PROJECT FOOTER -->

---

## GitHub 标星点赞

[![Stargazers](https://starchart.cc/yyle88/goenum.svg?variant=adaptive)](https://starchart.cc/yyle88/goenum)
