package main

import (
	"strings"
	"testing"

	"github.com/cirss/geist/util"
)

func TestBlazegraphCmd_import_two_triples(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	t.Run("import_nt", func(t *testing.T) {

		run("blazegraph destroy --dataset kb")
		run("blazegraph create --dataset kb")

		Main.InReader = strings.NewReader(`
			<http://tmcphill.net/data#x> <http://tmcphill.net/tags#tag> "seven" .
			<http://tmcphill.net/data#y> <http://tmcphill.net/tags#tag> "eight" .
		`)
		run("blazegraph import --format nt")

		outputBuffer.Reset()
		run("blazegraph export --format nt --sort=true")
		util.LineContentsEqual(t, outputBuffer.String(), `
			<http://tmcphill.net/data#x> <http://tmcphill.net/tags#tag> "seven" .
			<http://tmcphill.net/data#y> <http://tmcphill.net/tags#tag> "eight" .
		`)
	})

	t.Run("import_ttl", func(t *testing.T) {

		run("blazegraph destroy --dataset kb")
		run("blazegraph create --dataset kb")

		Main.InReader = strings.NewReader(`
			@prefix data: <http://tmcphill.net/data#> .
			@prefix tags: <http://tmcphill.net/tags#> .

			data:y tags:tag "eight" .
			data:x tags:tag "seven" .
		`)
		run("blazegraph import --format ttl")

		outputBuffer.Reset()
		run("blazegraph export --format nt --sort=true")
		util.LineContentsEqual(t, outputBuffer.String(), `
			<http://tmcphill.net/data#x> <http://tmcphill.net/tags#tag> "seven" .
			<http://tmcphill.net/data#y> <http://tmcphill.net/tags#tag> "eight" .
	`)
	})

	t.Run("import_jsonld", func(t *testing.T) {

		run("blazegraph destroy --dataset kb")
		run("blazegraph create --dataset kb")

		Main.InReader = strings.NewReader(`
			[
				{
					"@id": "http://tmcphill.net/data#x",
					"http://tmcphill.net/tags#tag": "seven"
				},
				{
					"@id": "http://tmcphill.net/data#y",
					"http://tmcphill.net/tags#tag": "eight"
				}
			]
		`)
		run("blazegraph import --format jsonld")

		outputBuffer.Reset()
		run("blazegraph export --format nt --sort=true")
		util.LineContentsEqual(t, outputBuffer.String(), `
			<http://tmcphill.net/data#x> <http://tmcphill.net/tags#tag> "seven"^^<http://www.w3.org/2001/XMLSchema#string> .
			<http://tmcphill.net/data#y> <http://tmcphill.net/tags#tag> "eight"^^<http://www.w3.org/2001/XMLSchema#string> .
`)
	})

	t.Run("import_ttl", func(t *testing.T) {

		run("blazegraph destroy --dataset kb")
		run("blazegraph create --dataset kb")

		Main.InReader = strings.NewReader(`
			@prefix data: <http://tmcphill.net/data#> .
			@prefix tags: <http://tmcphill.net/tags#> .

			data:y tags:tag "eight" .
			data:x tags:tag "seven" .
		`)
		run("blazegraph import --format ttl")

		outputBuffer.Reset()
		run("blazegraph export --format nt --sort=true")
		util.LineContentsEqual(t, outputBuffer.String(), `
			<http://tmcphill.net/data#x> <http://tmcphill.net/tags#tag> "seven" .
			<http://tmcphill.net/data#y> <http://tmcphill.net/tags#tag> "eight" .
	`)
	})

	t.Run("import_xml", func(t *testing.T) {

		run("blazegraph destroy --dataset kb")
		run("blazegraph create --dataset kb")

		Main.InReader = strings.NewReader(`
			<rdf:RDF xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#">

			<rdf:Description rdf:about="http://tmcphill.net/data#y">
				<tag xmlns="http://tmcphill.net/tags#">eight</tag>
			</rdf:Description>

			<rdf:Description rdf:about="http://tmcphill.net/data#x">
				<tag xmlns="http://tmcphill.net/tags#">seven</tag>
			</rdf:Description>

			</rdf:RDF>
		`)
		run("blazegraph import --format xml")

		outputBuffer.Reset()
		run("blazegraph export --format nt --sort=true")
		util.LineContentsEqual(t, outputBuffer.String(), `
			<http://tmcphill.net/data#x> <http://tmcphill.net/tags#tag> "seven" .
			<http://tmcphill.net/data#y> <http://tmcphill.net/tags#tag> "eight" .
	`)
	})
}
