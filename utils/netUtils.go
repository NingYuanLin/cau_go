package utils

import (
	"net"
)

func GetLocalhostIps() ([]string, error) {
	// https://blog.csdn.net/qq_41035588/article/details/121407894
	// TODO:实验性功能，在多网卡等情况下是否还有效？
	var ips []string
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		return ips, err
	}
	for _, address := range addresses {
		// 检查ip地址判断是否回环地址
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP.String())
				//fmt.Println(ipNet.IP.String())
			}
		}
	}
	return ips, nil
}
