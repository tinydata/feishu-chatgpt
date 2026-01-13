# Webpage Display Checker / 网页展示检查工具

检查网页是否能正常展示的命令行工具。

A command-line tool to check if a webpage displays normally.

## 功能特点 / Features

- ✅ 自动尝试 HTTPS 和 HTTP 连接 / Automatically tries both HTTPS and HTTP
- ✅ 支持重定向跟踪 / Supports redirect tracking  
- ✅ 验证页面内容可读性 / Verifies page content readability
- ✅ 返回详细状态信息 / Returns detailed status information
- ✅ 清晰的成功/失败指示 / Clear success/failure indication

## 构建 / Build

```bash
cd code
go build -o check-webpage ./cmd/check-domain/main.go
```

## 使用方法 / Usage

```bash
# 检查域名 / Check a domain
./check-webpage winbests.net

# 检查完整URL / Check a full URL
./check-webpage https://winbests.net
./check-webpage http://example.com

# 检查API端点 / Check an API endpoint
./check-webpage api.openai.com
```

## 输出示例 / Example Output

### 网页在线 / Webpage Online
```
Checking webpage: google.com
=====================================
Status: ✓ ONLINE (HTTP 200)
Details: Webpage displays normally via HTTPS (Status: 200 OK)
```

### 网页离线 / Webpage Offline
```
Checking webpage: winbests.net
=====================================
Status: ✗ OFFLINE/UNREACHABLE
Details: Cannot connect to webpage: dial tcp: lookup winbests.net: server misbehaving
```

## 退出码 / Exit Codes

- `0`: 网页正常展示 / Webpage displays normally
- `1`: 网页无法访问或展示异常 / Webpage is unreachable or displays abnormally

## 在代码中使用 / Use in Code

```go
import "start-feishubot/utils"

// 检查域名
isOnline, message, statusCode := utils.CheckWebpageDisplay("example.com")
if isOnline {
    fmt.Printf("网页正常，状态码: %d\n", statusCode)
} else {
    fmt.Printf("网页异常: %s\n", message)
}

// 检查完整URL
isOnline, message, statusCode := utils.CheckURLDisplay("https://example.com/api")
if isOnline {
    fmt.Printf("URL可访问，状态码: %d\n", statusCode)
}
```

## winbests.net 检查结果 / winbests.net Check Result

根据最新检查，**winbests.net 无法正常展示**。

According to the latest check, **winbests.net cannot display normally**.

详细报告请查看: [DOMAIN_CHECK_REPORT.md](../../DOMAIN_CHECK_REPORT.md)

For detailed report, see: [DOMAIN_CHECK_REPORT.md](../../DOMAIN_CHECK_REPORT.md)
