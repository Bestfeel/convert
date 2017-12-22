package main

import (
	"strings"
	"fmt"
	"strconv"
	"flag"
	"os"
)

const (
	COLOR_RED     = uint8(iota + 91)
	COLOR_GREEN
	COLOR_YELLOW
	COLOR_BLUE
	COLOR_MAGENTA
)

func HexToBye(hex string) []byte {
	length := len(hex) / 2
	slice := make([]byte, length)
	rs := []rune(hex)

	for i := 0; i < length; i++ {
		s := string(rs[i*2: i*2+2])
		value, _ := strconv.ParseInt(s, 16, 10)
		slice[i] = byte(value & 0xFF)
	}
	return slice
}

var (
	hexStr = flag.String("s", "", "输入16进制的字符串")
)

const LOGO = `
                                                  _
                                          _______| |
                                         |_________|
                                          _________
                                         |  _______|   万物互联
                                         | |   ____
                                         | |  |__  |   机智云
                                         | |_____| |
                                         |_________|   Gizwits
                                          机智云只为硬件而生的云服务



-s string

    	输入16进制的字符串

example:

convert -s="00000003  15  00  0091  04  0000000000000000000000000000000000"

输出:

0000000315000091040000000000000000000000000000000000
[0 0 0 3 21 0 0 145 4 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
[0,0,0,3,21,0,0,145,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]

`

func FormatByte(b []byte) string {
	s := make([]string, len(b))
	for i := range b {
		s[i] = strconv.Itoa(int(b[i]))
	}
	return strings.Join(s, ",")
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Printf("\x1b[%d;1m%s\x1b[0m", COLOR_MAGENTA, LOGO+"\n")
		return
	}
	flag.Parse()
	trim := strings.Replace(*hexStr, " ", "", -1)
	str := strings.Replace(trim, "\n", "", -1)
	fmt.Println(str)
	bye := HexToBye(str)
	fmt.Println(bye)
	fmt.Printf("[%s]\n", FormatByte(bye))
}
