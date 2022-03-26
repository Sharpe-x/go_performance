package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*描述
请解析IP地址和对应的掩码，进行分类识别。要求按照A/B/C/D/E类地址归类，不合法的地址和掩码单独归类。

所有的IP地址划分为 A,B,C,D,E五类

A类地址从1.0.0.0到126.255.255.255;

B类地址从128.0.0.0到191.255.255.255;

C类地址从192.0.0.0到223.255.255.255;

D类地址从224.0.0.0到239.255.255.255；

E类地址从240.0.0.0到255.255.255.255


私网IP范围是：

从10.0.0.0到10.255.255.255

从172.16.0.0到172.31.255.255

从192.168.0.0到192.168.255.255


子网掩码为二进制下前面是连续的1，然后全是0。（例如：255.255.255.32就是一个非法的掩码）
（注意二进制下全是1或者全是0均为非法子网掩码）

注意：
1. 类似于【0.*.*.*】和【127.*.*.*】的IP地址不属于上述输入的任意一类，也不属于不合法ip地址，计数时请忽略
2. 私有IP地址和A,B,C,D,E类地址是不冲突的*/

func main() {
	sc := bufio.NewScanner(os.Stdin)
	ip, mask := [4]uint8{}, [4]uint8{}
	var a, b, c, d, e, i, p int
	for sc.Scan() {
		str := sc.Text()
		ipStrs := strings.Split(str, "~")
		isIpStr, isMaskStr := IsIpStr(ipStrs[0], &ip), IsIpStr(ipStrs[1], &mask)
		// 忽略本地回环地址
		if isIpStr && isMaskStr && (ip[0] == 0 || ip[0] == 127) {
			continue
		}

		if !isIpStr || !isMaskStr || !IsMask(&mask) {
			i++
			continue
		}

		ip0 := ip[0]
		switch {
		case ip0 >= 1 && ip0 <= 126:
			a++
		case ip0 >= 128 && ip0 <= 191:
			b++
		case ip0 >= 192 && ip0 <= 223:
			c++
		case ip0 >= 224 && ip0 <= 239:
			d++
		case ip0 >= 240:
			e++
		}

		if ip[0] == 10 || (ip[0] == 172 && ip[1] >= 16 && ip[1] <= 31) || (ip[0] == 192 && ip[1] == 168) {
			p++
		}
	}

	fmt.Println(a, b, c, d, e, i, p)

}

func IsMask(m *[4]uint8) bool {
	var mask uint32
	mask = mask | uint32(m[0])<<24
	mask = mask | uint32(m[1])<<16
	mask = mask | uint32(m[2])<<8
	mask = mask | uint32(m[3])
	return mask != 0 && mask != 0xFFFFFFFF && mask == ((mask^0xFFFFFFFF)+1)|mask
}

func IsIpStr(ipStr string, result *[4]uint8) bool {
	strs := strings.Split(ipStr, ".")
	if len(strs) != 4 {
		return false
	}

	for i := 0; i < 4; i++ {
		temp, err := strconv.Atoi(strs[i])
		if err != nil || temp < 0 || temp > 255 {
			return false
		}
		result[i] = uint8(temp)
	}

	return true
}
