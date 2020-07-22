package solution

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
	"testing"
	"unicode"
)

func reorderLogFiles(logs []string) []string {
	if len(logs) <= 2 {
		return logs
	}
	var strLogs []string
	var numLogs []string
	for _, entry := range logs {
		index := strings.Index(entry, " ")
		if unicode.IsDigit(rune(entry[index+1])) {
			numLogs = append(numLogs, entry)
		} else {
			strLogs = append(strLogs, entry)
		}
	}
	sort.Slice(strLogs, func(i, j int) bool {
		iLog := strings.SplitN(strLogs[i], " ", 2)[1]
		jLog := strings.SplitN(strLogs[j], " ", 2)[1]
		if iLog == jLog {
			return strLogs[i] < strLogs[j]
		}
		return iLog < jLog
	})
	return append(strLogs, numLogs...)
}

func TestReorderLogFiles(t *testing.T) {
	testCases := []struct {
		Input  []string
		Output []string
	}{
		{
			Input:  []string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"},
			Output: []string{"let1 art can", "let3 art zero", "let2 own kit dig", "dig1 8 1 5 1", "dig2 3 6"},
		},
		{
			Input:  []string{"a1 9 2 3 1", "g1 act car", "zo4 4 7", "ab1 off key dog", "a8 act zoo"},
			Output: []string{"g1 act car", "a8 act zoo", "ab1 off key dog", "a1 9 2 3 1", "zo4 4 7"},
		},
		{
			Input:  []string{"8 fj dzz k", "5r 446 6 3", "63 gu psub", "5 ba iedrz", "6s 87979 5", "3r 587 01", "jc 3480612", "bb wsrd kp", "b aq cojj", "r5 6316 71"},
			Output: []string{"b aq cojj", "5 ba iedrz", "8 fj dzz k", "63 gu psub", "bb wsrd kp", "5r 446 6 3", "6s 87979 5", "3r 587 01", "jc 3480612", "r5 6316 71"},
		},
	}
	for _, tc := range testCases {
		got := reorderLogFiles(tc.Input)
		if !reflect.DeepEqual(got, tc.Output) {
			fmt.Printf("given : %#v\n", tc.Input)
			fmt.Printf("got   : %#v\n", got)
			fmt.Printf("wanted: %#v\n", tc.Output)
			t.Error()
		}
	}
}
