package snakecase_test

import (
	"github.com/azer/snakecase"
	"strings"
	"testing"
)

type SnakeTest struct {
	input  string
	output string
}

var tests = []SnakeTest{
	{"a", "a"},
	{"snake", "snake"},
	{"A", "a"},
	{"ID", "id"},
	{"MOTD", "motd"},
	{"Snake", "snake"},
	{"SnakeTest", "snake_test"},
	{"APIResponse", "api_response"},
	{"SnakeID", "snake_id"},
	{"SnakeIDGoogle", "snake_id_google"},
	{"LinuxMOTD", "linux_motd"},
	{"OMGWTFBBQ", "omgwtfbbq"},
	{"omg_wtf_bbq", "omg_wtf_bbq"},
}

func TestSnakeCase(t *testing.T) {
	for _, test := range tests {
		if snakecase.SnakeCase(test.input) != test.output {
			t.Errorf(`SnakeCase("%s"), wanted "%s", got \%s"`, test.input, test.output, snakecase.SnakeCase(test.input))
		}
	}
}

var benchmarks = []string{
	"a",
	"snake",
	"A",
	"Snake",
	"SnakeTest",
	"SnakeID",
	"SnakeIDGoogle",
	"LinuxMOTD",
	"OMGWTFBBQ",
	"omg_wtf_bbq",
}

func BenchmarkSnakeCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, input := range benchmarks {
			snakecase.SnakeCase(input)
		}
	}
}

func BenchmarkToLower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, input := range benchmarks {
			strings.ToLower(input)
		}
	}
}
