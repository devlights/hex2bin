package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/devlights/gomy/convert"
)

var (
	splitCount int
)

func init() {
	flag.IntVar(&splitCount, "s", 8, "ビットを分割表示する桁数")
}

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
	binStr, err := Convert(v)
	if err != nil {
		fmt.Printf("[Error] %s\n", err)
		return 2
	}

	if splitCount > 0 {
		var builder strings.Builder

		start := 0
		for {
			if len(binStr) < (start + splitCount) {
				builder.WriteString(binStr[start:])
				break
			}

			builder.WriteString(binStr[start : start+splitCount])
			builder.WriteString(" ")

			start += splitCount
		}

		binStr = builder.String()
	}

	fmt.Printf("%s\t-->\t%s\n", v, binStr)

	return 0
}

// Convert -- 指定された16進数文字列を2進数文字列に変換します.
func Convert(v string) (string, error) {
	return convert.Hex2Bin(v, "", -1)
}
