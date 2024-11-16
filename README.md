# 高并发 DNS 压力测试工具

**这是一个使用 Go 编写的高并发 DNS 压力测试工具。**  
它能够通过模拟大量 DNS 请求，借助公共 DNS 服务器放大请求，从而测试自建权威 DNS 的性能和稳定性。

---

## 特性

1. **高并发**：利用 Goroutine 并发发送 DNS 查询，最大化测试压力。
2. **公共 DNS 放大**：通过多个公共 DNS 服务器转发请求，快速施加负载。
3. **随机子域名生成**：动态生成随机的多级子域名，避免缓存影响，确保请求多样性。
4. **支持自定义查询类型**：支持 A、MX、NS 等多种 DNS 查询类型。
5. **易于扩展**：可自定义增加 DNS 服务器列表和测试参数。

---

## 使用方法

### 1. 安装

确保您已安装 Go（1.16 或更高版本）。

```bash
go build -o dns_stress_test main.go
```

### 2. 运行

```bash
./dns_stress_test -domain="example.com" -type="A"
```

- **`-domain`**：测试的主域名（默认值为 `example.com`）。
- **`-type`**：DNS 查询类型（默认值为 `A`，支持 `A`、`MX`、`NS` 等）。

### 3. 示例

测试域名 `test.com` 的 A 记录：

```bash
./dns_stress_test -domain="test.com" -type="A"
```

测试域名 `example.com` 的 MX 记录：

```bash
./dns_stress_test -domain="example.com" -type="MX"
```

---

## 配置

### 添加更多 DNS 服务器

编辑代码中的 `dnsServers` 列表，添加您需要的公共 DNS 服务器。例如：

```go
dnsServers := []string{
    "8.8.8.8:53",
    "8.8.4.4:53",
    "1.1.1.1:53",
    "1.0.0.1:53",
    "9.9.9.9:53",
    "208.67.222.222:53", // OpenDNS
    "208.67.220.220:53", // OpenDNS
}
```

### 调整发送频率

可通过在 `sendDNSQueries` 函数内调整 `time.Sleep` 控制请求频率，例如：

```go
time.Sleep(10 * time.Millisecond) // 每 10 毫秒发送一次请求
```

---

## 工作原理

1. **公共 DNS 放大**  
   工具通过公共 DNS 服务器代理请求，将负载传递到自建的权威 DNS，从而测试其性能。

2. **随机子域名生成**  
   每次请求都会生成一个随机多级子域名（例如 `abc.def.ghi.example.com`），确保测试多样性，绕过缓存。

3. **高并发请求**  
   使用 Goroutine 并发向多个公共 DNS 服务器发送请求，模拟高负载环境。

---

## 注意事项

1. **测试合法性**  
   在使用该工具进行压力测试前，请确保对目标 DNS 服务器拥有合法测试权限。

2. **对网络的影响**  
   该工具可能产生大量网络流量，请避免对生产环境造成不必要的影响。

3. **公共 DNS 限制**  
   部分公共 DNS 服务器可能对过多请求采取限流策略，请合理配置负载分布。

---

## 许可证

本项目基于 [Apache 2.0 许可证](https://www.apache.org/licenses/LICENSE-2.0) 开源，您可以自由使用、修改和分发，但需保留原始版权声明及许可证信息。

---

## 联系

如有问题或建议，请提交 Issue 或 Pull Request。
