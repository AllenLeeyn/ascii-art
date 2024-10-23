package main

import (
	fileMgr "ascii-art/pkg/fileMgr"
	"bytes"
	"os"
	"testing"
)

func TestAsciiArtBasic(t *testing.T) {
	testCases := []struct {
		args     []string
		expected string
	}{
		{
			args:     []string{".", ""},
			expected: "",
		},
		{
			args:     []string{".", "\\n"},
			expected: "\n",
		},
		{
			args:     []string{".", "Hello\\n"},
			expected: "./test/basic/result02.txt",
		},
		{
			args:     []string{".", "hello"},
			expected: "./test/basic/result03.txt",
		},
		{
			args:     []string{".", "HeLlO"},
			expected: "./test/basic/result04.txt",
		},
		{
			args:     []string{".", "Hello There"},
			expected: "./test/basic/result05.txt",
		},
		{
			args:     []string{".", "1Hello 2There"},
			expected: "./test/basic/result06.txt",
		},
		{
			args:     []string{".", "{Hello There}"},
			expected: "./test/basic/result07.txt",
		},
		{
			args:     []string{".", "Hello\\nThere"},
			expected: "./test/basic/result08.txt",
		},
		{
			args:     []string{".", "Hello\\n\\nThere"},
			expected: "./test/basic/result09.txt",
		},
		{
			args:     []string{".", "grit:lab_2024\\n\\n/\\5c11 /\\r7\\n"},
			expected: "./test/basic/result10.txt",
		},
		{
			args:     []string{"."},
			expected: "./test/basic/resultBad00.txt",
		},
		{
			args:     []string{".", "Du äter smörgåsar"},
			expected: "./test/basic/resultBad01.txt",
		},
	}

	for i, tc := range testCases {
		result := getResults(tc.args)

		if i > 1 {
			tc.expected = fileMgr.ReadFile(tc.expected)
		}

		if result != tc.expected {
			t.Errorf("\033[31mTest case %d, %s \nfailed: expected\n%s\n got \n%s\033[39m", i+1, tc.args, tc.expected, result)

		} else {
			t.Logf("\033[32mTest case %d, %s \npassed: result\n%s\033[39m", i+1, tc.args, result)

		}
	}
}

func getResults(tcArgs []string) string {
	ogArgs := os.Args
	ogStdout := os.Stdout

	r, w, _ := os.Pipe()
	os.Args = tcArgs
	os.Stdout = w

	main()
	w.Close()

	os.Args = ogArgs
	os.Stdout = ogStdout

	var resultBuf bytes.Buffer
	resultBuf.ReadFrom(r)
	result := resultBuf.String()

	return result
}
