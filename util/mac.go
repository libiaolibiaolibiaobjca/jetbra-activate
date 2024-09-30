package util

import (
	"fmt"
	"math/rand"
	"net"
	"time"
)

func RandMac() string {
	// 初始化随机数生成器
	rand.Seed(time.Now().UnixNano())

	// 生成一个随机的MAC地址
	macAddr := make(net.HardwareAddr, 6)

	// 确保MAC地址的第一位不为0
	for macAddr[0] == 0 {
		rand.Read(macAddr)
	}

	// 将MAC地址转换为字符串形式
	macStr := macAddr.String()

	fmt.Println("Generated MAC address:", macStr)

	return macStr
}
