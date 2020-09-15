package util

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/ucarion/jcs"
)

func canonicalJSON(jsonString string) (string, error) {
	var originalJSON interface{}
	json.Unmarshal([]byte(jsonString), &originalJSON)
	canonicalJSONString, err := jcs.Format(originalJSON)
	return canonicalJSONString, err
}

func StringEquals(t *testing.T, actual string, expected string) {
	if actual != expected {
		t.Log("assertStringEquals:\n\nexpected:\n" + expected + "\nactual:\n" + actual + "\n")
		t.FailNow()
	}
}

func TrimEachLine(s string) string {
	var sb strings.Builder
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		trimmedLine := Trim(line)
		sb.WriteString(trimmedLine)
		sb.WriteString("\n")
	}
	return sb.String()
}

func Trim(s string) string {
	return strings.Trim(s, " \n\r\t")
}

func LineContentsEqual(t *testing.T, actual string, expect string) {
	actualContent := strings.Split(Trim(actual), "\n")
	expectContent := strings.Split(Trim(expect), "\n")
	lineCount := len(actualContent)
	if lineCount != len(expectContent) {
		StringEquals(t, actual, expect)
	}
	for i := 0; i < lineCount; i++ {
		StringEquals(t, Trim(actualContent[i]), Trim(expectContent[i]))
	}
}

func JSONEquals(t *testing.T, actual string, expected string) {
	actualCanonical, _ := canonicalJSON(actual)
	expectedCanonical, _ := canonicalJSON(expected)
	if actualCanonical != expectedCanonical {
		t.Log("assertEquivalentJSON:\n" +
			"\nexpected: " + expectedCanonical +
			"\nactual:   " + actualCanonical + "\n")
		t.Fail()
	}
}
