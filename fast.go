package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/ddo/go-fast"
)

func main() {
	var kb, mb, gb bool
	var iface string
	flag.BoolVar(&kb, "k", false, "Format output in Kbps")
	flag.BoolVar(&mb, "m", false, "Format output in Mbps")
	flag.BoolVar(&gb, "g", false, "Format output in Gbps")
	flag.StringVar(&iface, "i", "", "Network interface to use (e.g., tun0, eth0)")

	flag.Parse()

	if kb && (mb || gb) || (mb && kb) {
		fmt.Println("You may have at most one formating switch. Choose either -k, -m, or -g")
		os.Exit(-1)
	}

	// 创建自定义的 HTTP Transport，仅使用 IPv4
	ipv4Transport := &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			// 强制使用 tcp4 网络
			if network == "tcp" {
				network = "tcp4"
			}
			
			d := &net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: false, // 禁用双栈，仅使用 IPv4
			}
			
			// 如果指定了接口，设置本地地址绑定
			if iface != "" {
				intf, err := net.InterfaceByName(iface)
				if err != nil {
					return nil, fmt.Errorf("找不到接口 %s: %v", iface, err)
				}
				
				addrs, err := intf.Addrs()
				if err != nil {
					return nil, fmt.Errorf("获取接口 %s 地址失败: %v", iface, err)
				}
				
				// 查找 IPv4 地址
				var localIP net.IP
				for _, addr := range addrs {
					var ip net.IP
					switch v := addr.(type) {
					case *net.IPNet:
						ip = v.IP
					case *net.IPAddr:
						ip = v.IP
					default:
						// 尝试解析字符串格式的地址
						addrStr := addr.String()
						if idx := strings.Index(addrStr, "/"); idx != -1 {
							addrStr = addrStr[:idx]
						}
						ip = net.ParseIP(addrStr)
					}
					
					if ip != nil && ip.To4() != nil {
						localIP = ip
						break
					}
				}
				
				if localIP == nil {
					return nil, fmt.Errorf("接口 %s 没有 IPv4 地址", iface)
				}
				
				d.LocalAddr = &net.TCPAddr{IP: localIP}
				fmt.Printf("使用本地地址: %s\n", localIP.String())
			}
			
			return d.DialContext(ctx, network, addr)
		},
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	// 设置默认的 HTTP 客户端使用 IPv4-only transport
	http.DefaultTransport = ipv4Transport
	http.DefaultClient = &http.Client{Transport: ipv4Transport}

	fastCom := fast.New()

	// 显示使用的接口
	if iface != "" {
		fmt.Printf("使用网络接口: %s\n", iface)
	}

	// 初始化
	fmt.Println("正在连接测试服务器...")
	err := fastCom.Init()
	if err != nil {
		fmt.Printf("错误: %v\n", err)
		os.Exit(1)
	}

	// 获取测试服务器 URLs
	urls, err := fastCom.GetUrls()
	if err != nil {
		fmt.Printf("错误: %v\n", err)
		os.Exit(1)
	}

	// 解析并显示服务器信息
	if len(urls) > 0 {
		firstURL := urls[0]
		parsedURL, err := url.Parse(firstURL)
		if err == nil {
			host := parsedURL.Hostname()
			fmt.Printf("测试服务器: %s\n", host)
			
			// DNS 解析获取 IP 地址 (仅 IPv4)
			ips, err := net.LookupIP(host)
			if err == nil {
				var ipv4Addrs []string
				for _, ip := range ips {
					if ipv4 := ip.To4(); ipv4 != nil {
						ipv4Addrs = append(ipv4Addrs, ipv4.String())
					}
				}
				if len(ipv4Addrs) > 0 {
					fmt.Printf("服务器 IP: %s\n", strings.Join(ipv4Addrs, ", "))
				}
			}
		}
	}

	// 开始测试
	fmt.Println("正在测试下载速度...")
	
	// 测量速度
	KbpsChan := make(chan float64)
	
	var finalStatus string
	
	go func() {
		for Kbps := range KbpsChan {
			value, units := format(Kbps, kb, mb, gb)
			if kb || mb || gb {
				finalStatus = fmt.Sprintf("%s", value)
			} else {
				finalStatus = fmt.Sprintf("%s %s", value, units)
			}
		}
	}()

	err = fastCom.Measure(urls, KbpsChan)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
		os.Exit(1)
	}

	// 显示最终结果
	fmt.Printf("\n下载速度: %s\n", finalStatus)
}

func formatGbps(Kbps float64) (string, string, float64) {
	f := "%.2f"
	unit := "Gbps"
	value := Kbps / 1000000
	return f, unit, value
}

func formatMbps(Kbps float64) (string, string, float64) {
	f := "%.2f"
	unit := "Mbps"
	value := Kbps / 1000
	return f, unit, value
}

func formatKbps(Kbps float64) (string, string, float64) {
	f := "%.f"
	unit := "Kbps"
	value := Kbps
	return f, unit, value
}

func format(Kbps float64, kb bool, mb bool, gb bool) (string, string) {
	var value float64
	var unit string
	var f string

	if kb {
		f, unit, value = formatKbps(Kbps)
	} else if mb {
		f, unit, value = formatMbps(Kbps)
	} else if gb {
		f, unit, value = formatGbps(Kbps)
	} else if Kbps > 1000000 { // Gbps
		f, unit, value = formatGbps(Kbps)
	} else if Kbps > 1000 { // Mbps
		f, unit, value = formatMbps(Kbps)
	} else {
		f, unit, value = formatKbps(Kbps)
	}

	strValue := fmt.Sprintf(f, value)
	return strValue, unit
}