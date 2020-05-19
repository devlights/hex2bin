package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	os.Exit(run())
}

func run() int {
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("no arguments\tusage:: hex2bin hex-value(e.g: 0xFF)")
		return 1
	}

	v := args[0]
	b, err := Convert(v)
	if err != nil {
		fmt.Printf("[Error] %s\n", err)
		return 2
	}

	fmt.Printf("%s\t-->\t0b%b\n", v, b)

	return 0
}

// Convert -- 指定された16進数文字列を2進数に変換します.
func Convert(v string) (int64, error) {
	if len(v) == 0 {
		return 0, nil
	}

	if strings.HasPrefix(v, "0x") {
		v = strings.Replace(v, "0x", "", 1)
	}

	return strconv.ParseInt(v, 16, 32)
}
