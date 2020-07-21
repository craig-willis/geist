package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/tmcphillips/blazegraph-util/assert"
	"github.com/tmcphillips/blazegraph-util/blazegraph"
)

func TestBlazegraphClient_GetAllTriplesAsJSON_EmptyStore(t *testing.T) {
	bc := blazegraph.NewClient()
	bc.DeleteAll()
	triples, _ := bc.SelectAll()
	actual, _ := triples.JSONString()
	assert.JSONEquals(t, actual,
		`{
			"head" : {
				"vars" : [ "s", "p", "o" ]
			},
			"results" : {
				"bindings" : [ ]
			}
		}`)
}

func TestBlazegraphClient_InsertOneTriple(t *testing.T) {
	bc := blazegraph.NewClient()
	bc.DeleteAll()
	bc.PostData("application/x-turtle", []byte(`
	@prefix t: <http://tmcphill.net/tags#> .
	@prefix d: <http://tmcphill.net/data#> .
	d:y t:tag "seven" .
	`))

	resultSet, _ := bc.Select(
		`prefix t: <http://tmcphill.net/tags#>
		 SELECT ?s ?o
		 WHERE
		 { ?s t:tag ?o }
		 `)

	actual, _ := resultSet.JSONString()

	assert.JSONEquals(t,
		actual,
		`{
			"head" : { "vars" : [ "s", "o" ] },
			"results" : {
			  "bindings" : [ {
				"s" : { "type" : "uri",     "value" : "http://tmcphill.net/data#y" },
				"o" : { "type" : "literal", "value" : "seven" }
			  } ]
			}
		  }`)
}

func TestBlazegraphClient_InsertTwoTriples(t *testing.T) {
	bc := blazegraph.NewClient()
	bc.DeleteAll()
	bc.PostData("application/x-turtle", []byte(`
		@prefix t: <http://tmcphill.net/tags#> .
		@prefix d: <http://tmcphill.net/data#> .

		d:x t:tag "seven" .
		d:y t:tag "eight" .
	`))

	resultSet, _ := bc.Select(
		`prefix ab: <http://tmcphill.net/tags#>
		 SELECT ?s ?o
		 WHERE
		 { ?s ab:tag ?o }
		 `)

	actual, _ := resultSet.JSONString()

	assert.JSONEquals(t,
		actual,
		`{
			"head" : { "vars" : [ "s", "o" ] },
			"results" : {
			  "bindings" : [ {
				"s" : { "type" : "uri",     "value" : "http://tmcphill.net/data#x" },
				"o" : { "type" : "literal", "value" : "seven" }
			}, {
				"s" : { "type" : "uri",     "value" : "http://tmcphill.net/data#y" },
				"o" : { "type" : "literal", "value" : "eight" }
			  } ]
			}
		  }`)
}

func TestBlazegraphClient_InsertTwoTriples_Struct(t *testing.T) {
	bc := blazegraph.NewClient()
	bc.DeleteAll()
	bc.PostData("application/x-turtle", []byte(`
		@prefix t: <http://tmcphill.net/tags#> .
		@prefix d: <http://tmcphill.net/data#> .

		d:x t:tag "seven" .
		d:y t:tag "eight" .
	`))
	rs, _ := bc.Select(
		`prefix ab: <http://tmcphill.net/tags#>
		 SELECT ?s ?o
		 WHERE
		 { ?s ab:tag ?o }
		 `)

	assert.StringEquals(t, strings.Join(rs.Vars(), ", "), "s, o")

	assert.StringEquals(t, rs.Bindings()[0]["s"].Type, "uri")
	assert.StringEquals(t, rs.Bindings()[0]["s"].Value, "http://tmcphill.net/data#x")
	assert.StringEquals(t, rs.Bindings()[0]["o"].Type, "literal")
	assert.StringEquals(t, rs.Bindings()[0]["o"].Value, "seven")

	assert.StringEquals(t, rs.Bindings()[1]["s"].Type, "uri")
	assert.StringEquals(t, rs.Bindings()[1]["s"].Value, "http://tmcphill.net/data#y")
	assert.StringEquals(t, rs.Bindings()[1]["o"].Type, "literal")
	assert.StringEquals(t, rs.Bindings()[1]["o"].Value, "eight")
}

func ExampleBlazegraphClient_DumpAsNTriples() {
	bc := blazegraph.NewClient()
	bc.DeleteAll()
	bc.PostData("application/x-turtle", []byte(`
		@prefix t: <http://tmcphill.net/tags#> .
		@prefix d: <http://tmcphill.net/data#> .

		d:x t:tag "seven" .
		d:y t:tag "eight" .
	`))
	triples, _ := bc.ConstructAll("text/plain")
	fmt.Println(triples)
	// Output:
	// <http://tmcphill.net/data#y> <http://tmcphill.net/tags#tag> "eight" .
	// <http://tmcphill.net/tags#tag> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/1999/02/22-rdf-syntax-ns#Property> .
	// <http://tmcphill.net/tags#tag> <http://www.w3.org/2000/01/rdf-schema#subPropertyOf> <http://tmcphill.net/tags#tag> .
	// <http://tmcphill.net/data#x> <http://tmcphill.net/tags#tag> "seven" .
}
