package main

import (
	"fmt"
	"strconv"
)

var pn = fmt.Println

func main() {
	// ParseFloat 将字符串转换为浮点数
	// s：要转换的字符串
	// bitSize：指定浮点类型（32:float32、64:float64）
	// 如果 s 是合法的格式，而且接近一个浮点值，
	// 则返回浮点数的四舍五入值（依据 IEEE754 的四舍五入标准）
	// 如果 s 不是合法的格式，则返回“语法错误”
	// 如果转换结果超出 bitSize 范围，则返回“超出范围”
	f, _ := strconv.ParseFloat("1.23", 32)
	pn(f)

	// ParseInt 将字符串转换为 int 类型
	// s：要转换的字符串
	// base：进位制（2 进制到 36 进制）
	// bitSize：指定整数类型（0:int、8:int8、16:int16、32:int32、64:int64）
	// 返回转换后的结果和转换时遇到的错误
	// 如果 base 为 0，则根据字符串的前缀判断进位制（0x:16，0:8，其它:10）

	i1, _ := strconv.ParseInt("2200000000", 0, 32)
	pn(i1)

	i2, _ := strconv.ParseInt("2200000000", 0, 64)
	pn(i2)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
