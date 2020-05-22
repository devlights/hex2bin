package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/devlights/gomy/convert"
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

	fmt.Printf("%s\t-->\t%s\n", v, b)

	return 0
}

// Convert -- 指定された16進数文字列を2進数文字列に変換します.
func Convert(v string) (string, error) {
	return convert.Hex2Bin(v, "0b", 0)
}
