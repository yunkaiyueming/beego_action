package helpers

import (
	"fmt"
	"net"
)

func GetHost() {
	domain := "www.baidu.com"
	ip := "124.193.0.2"
	p := fmt.Println
	p(net.LookupHost(domain))  //获得域名对应的所有Ip地址
	p(net.LookupIP(domain))    //LookupIP函数查询主机的ipv4和ipv6地址序列。
	p(net.LookupCNAME(domain)) //查找dns规范名
	p(net.LookupAddr(ip))
	p(net.LookupPort("tcp", "telnet")) //查看主机tcp telnet的服务的端口
	p(net.ResolveIPAddr("ip", domain)) //将域名解析Ip地址
	p(net.ParseIP(ip))                 //转换ip字符串为ip对象
}
