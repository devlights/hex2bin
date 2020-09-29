package main

import (
	"testing"
)

func TestConvert(t *testing.T) {
	type (
		testin struct {
			val string
		}
		testout struct {
			result string
		}
		testcase struct {
			name string
			in   testin
			out  testout
		}
	)

	cases := []testcase{
		{
			name: "convert 0xFF",
			in:   testin{val: "0xFF"},
			out:  testout{result: "11111111"},
		},
		{
			name: "convert 0x08",
			in:   testin{val: "0x08"},
			out:  testout{result: "00001000"},
		},
		{
			name: "convert 0x04",
			in:   testin{val: "0x04"},
			out:  testout{result: "00000100"},
		},
		{
			name: "convert 0x02",
			in:   testin{val: "0x02"},
			out:  testout{result: "00000010"},
		},
		{
			name: "convert 0x01",
			in:   testin{val: "0x01"},
			out:  testout{result: "00000001"},
		},
		{
			name: "convert 0x0E",
			in:   testin{val: "0x0E"},
			out:  testout{result: "00001110"},
		},
	}

	for _, c := range cases {
		b, err := Convert(c.in.val)
		if err != nil {
			t.Error(err)
		}

		if c.out.result != b {
			t.Errorf("[want] %v\t[got] %v\n", c.out.result, b)
		}
	}
}
