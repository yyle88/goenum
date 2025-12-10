[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/goenum/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/goenum/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/goenum)](https://pkg.go.dev/github.com/yyle88/goenum)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/goenum/main.svg)](https://coveralls.io/github/yyle88/goenum?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.25+-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/yyle88/goenum.svg)](https://github.com/yyle88/goenum/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/goenum)](https://goreportcard.com/report/github.com/yyle88/goenum)

# GOENUM

Go enum code generation that enables different business domains to share common enum names like OK, ERROR, PENDING without naming conflicts through namespace isolation.

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->

## CHINESE README

[中文说明](README.zh.md)
<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

## Features

🔒 **Namespace Isolation** - Each domain has its own enum space, preventing naming conflicts
⚡ **Enum Collection** - Generated `Enums()` returns collection with `Lookup`, `List`, `Get` methods
🎯 **Clean Code** - Intuitive syntax that matches business logic patterns
✅ **Compile Protection** - Catch enum misuse at build time, not runtime
🌍 **Multi-Language** - Generate enums using various language characters

## Installation

```bash
go get github.com/yyle88/goenum
```

## Usage

Go lacks true enum namespaces. Different domains can't share common value names like `OK`, `ERROR`, `PENDING`.

### Before: Verbose Prefixes Required

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
```

```go
// Verbose switch statements with long prefixes
func processPackage(status string) {
    switch PackageStatus(status) {
    case PackagePending:
        // handle pending
    case PackageConfirmed:
        // handle confirmed
    case PackageShipped:
        // handle shipped
    case PackageDelivered:
        // handle delivered
    }
}
```

```go
func processPayment(status string) {
    switch PaymentStatus(status) {
    case PaymentPending:
        // handle pending
    case PaymentFailed:
        // handle failed
    case PaymentSuccess:
        // handle success
    case PaymentRefund:
        // handle refund
    }
}
```

### With GOENUM: Clean Namespace Methods

```go
// Each domain gets its own clean namespace
func processPackage(status string) {
    switch PackageStatusEnum(status) {
    case PackageStatus.Pending():
        // handle pending
    case PackageStatus.Confirmed():
        // handle confirmed
    case PackageStatus.Shipped():
        // handle shipped
    case PackageStatus.Delivered():
        // handle delivered
    }
}
```

```go
func processPayment(status string) {
    switch PaymentStatusEnum(status) {
    case PaymentStatus.Pending():
        // handle pending
    case PaymentStatus.Failed():
        // handle failed
    case PaymentStatus.Success():
        // handle success
    case PaymentStatus.Refund():
        // handle refund
    }
}
```

### Enum Validation with Enums()

Each generated enum type has an `Enums()` method returning a collection that supports validation and lookup:

```go
// Validate enum value
if _, ok := PackageStatus.Enums().Lookup(status); !ok {
    return errors.New("invalid package status")
}

// Get all valid enum values
allStatuses := PackageStatus.Enums().List()
```

## Multi-Language Support

GOENUM supports enum generation in multiple languages:

```go
// Chinese enum usage example
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
// Traditional Chinese enum example
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
// Japanese enum example
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
// Korean enum example
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

**Examples**: See [examples](internal/examples)

---

## Related Projects

- **[enum](https://github.com/yyle88/enum)** - Go enum collection management with metadata support and map-based lookup
- **[goenum](https://github.com/yyle88/goenum)** - Go enum code generation with namespace isolation (this package)
- **[protoenum](https://github.com/go-xlan/protoenum)** - Protocol Buffers enum code generation with type-safe operations

---

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-25 03:52:28.131064 +0000 UTC -->

## 📄 License

MIT License - see [LICENSE](LICENSE).

---

## 💬 Contact & Feedback

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Mistake reports?** Open an issue on GitHub with reproduction steps
- 💡 **Fresh ideas?** Create an issue to discuss
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share the use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize through reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo to get new releases and features
- 🌟 **Success stories?** Share how this package improved the workflow
- 💬 **Feedback?** We welcome suggestions and comments

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage UI).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement the changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation to support client-facing changes
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a merge request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project via submitting merge requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Have Fun Coding with this package!** 🎉🎉🎉

<!-- TEMPLATE (EN) END: STANDARD PROJECT FOOTER -->

---

## GitHub Stars

[![Stargazers](https://starchart.cc/yyle88/goenum.svg?variant=adaptive)](https://starchart.cc/yyle88/goenum)
