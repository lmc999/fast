# Fast v1.0.0 - IPv4-only 版本

## 🚀 新功能

### 网络功能
- **IPv4-only 模式**：强制所有连接使用 IPv4，禁用 IPv6
- **接口绑定**：新增 `-i` 参数，支持指定网络接口进行测试

### 界面优化  
- **简洁输出**：移除加载动画，直接显示测试结果
- **服务器信息**：显示测试服务器的域名和 IP 地址
- **中文提示**：所有提示信息改为中文

## 📦 下载说明

| 文件名 | 平台 | 架构 |
|--------|------|------|
| fast_linux_amd64.tar.gz | Linux | x86_64 |
| fast_linux_arm64.tar.gz | Linux | ARM64 |
| fast_darwin_amd64.tar.gz | macOS | Intel |
| fast_darwin_arm64.tar.gz | macOS | Apple Silicon |
| fast_windows_amd64.exe.tar.gz | Windows | x86_64 |
| fast_windows_arm64.exe.tar.gz | Windows | ARM64 |

## 🔧 使用示例

```bash
# 基本使用
./fast

# 指定单位
./fast -m    # Mbps
./fast -k    # Kbps
./fast -g    # Gbps

# 指定网络接口
./fast -i eth0
./fast -i tun0
```

## 🛠️ 技术实现

- 自定义 HTTP Transport 强制 tcp4 连接
- 通过 net.InterfaceByName 绑定指定接口
- DNS 查询只获取 IPv4 地址

## 📝 更新日志

- 添加 IPv4-only 支持
- 添加网络接口绑定功能
- 优化输出格式
- 更新文档