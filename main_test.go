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
			result int64
		}
		testcase struct {
			in  testin
			out testout
		}
	)

	cases := []testcase{
		{
			in:  testin{val: "0xFF"},
			out: testout{result: 0b11111111},
		},
		{
			in:  testin{val: "0x08"},
			out: testout{result: 0b1000},
		},
		{
			in:  testin{val: "0x04"},
			out: testout{result: 0b100},
		},
		{
			in:  testin{val: "0x02"},
			out: testout{result: 0b10},
		},
		{
			in:  testin{val: "0x01"},
			out: testout{result: 0b1},
		},
		{
			in:  testin{val: "0x0E"},
			out: testout{result: 0b1110},
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
