package sparql

import (
	"strings"
	"testing"

	"github.com/tmcphillips/blazegraph-util/assert"
)

var sr = SparqlResult{
	Head{[]string{"s", "o"}},
	Results{[]Binding{{
		"s": {"uri", "http://tmcphill.net/data#x"},
		"o": {"literal", "seven"},
	}, {
		"s": {"uri", "http://tmcphill.net/data#y"},
		"o": {"literal", "eight"},
	}}}}

func TestSparqlResult_Vars(t *testing.T) {
	assert.StringEquals(t, strings.Join(sr.Vars(), ", "), "s, o")
}

func TestSparqlResult_Bindings(t *testing.T) {

	assert.StringEquals(t, sr.Bindings()[0]["s"].Type, "uri")
	assert.StringEquals(t, sr.Bindings()[0]["s"].Value, "http://tmcphill.net/data#x")
	assert.StringEquals(t, sr.Bindings()[0]["o"].Type, "literal")
	assert.StringEquals(t, sr.Bindings()[0]["o"].Value, "seven")

	assert.StringEquals(t, sr.Bindings()[1]["s"].Type, "uri")
	assert.StringEquals(t, sr.Bindings()[1]["s"].Value, "http://tmcphill.net/data#y")
	assert.StringEquals(t, sr.Bindings()[1]["o"].Type, "literal")
	assert.StringEquals(t, sr.Bindings()[1]["o"].Value, "eight")
}

func TestSparqlResult_BoundValues(t *testing.T) {
	assert.StringEquals(t,
		strings.Join(sr.BoundValues(0), ", "),
		"http://tmcphill.net/data#x, seven")
	assert.StringEquals(t,
		strings.Join(sr.BoundValues(1), ", "),
		"http://tmcphill.net/data#y, eight")
}
