package tests

import (
	"testing"

	"github.com/cirss/geist"
	"github.com/cirss/geist/util"
)

func TestReportTemplate_constant_macro_function(t *testing.T) {

	rt := geist.NewTemplate(
		"main",
		`{{{
		{{ macro "M1" '''constant macro''' }}
		}}}

		{{ M1 }}
		`, nil, nil)

	rt.Parse()
	actual, _ := rt.Expand(nil)
	util.LineContentsEqual(t, actual,
		`
		constant macro
	`)
}

func TestReportTemplate_macro_function_with_one_parameter(t *testing.T) {

	rt := geist.NewTemplate(
		"main",
		`
		{{{
		{{ macro "M1" "P1" '''macro with one parameter: {{$P1}}''' }}
		}}}

		{{ M1 "AA" }}
		`, nil, nil)

	rt.Parse()
	actual, _ := rt.Expand(nil)
	util.LineContentsEqual(t, actual,
		`
		macro with one parameter: AA
	`)
}

func TestReportTemplate_macro_function_with_two_parameters(t *testing.T) {

	rt := geist.NewTemplate(
		"main",
		`
		{{{
		{{ macro "M1" "P1" "P2" '''macro with two parameters: {{$P1}}, {{$P2}}''' }}
		}}}

		{{ M1 "AA" "BB" }}
		`, nil, nil)

	rt.Parse()
	actual, _ := rt.Expand(nil)
	util.LineContentsEqual(t, actual,
		`
		macro with two parameters: AA, BB
	`)
}

func TestReportTemplate_macro_function_calling_macro_function(t *testing.T) {

	rt := geist.NewTemplate(
		"main",
		`
		{{{
		{{ macro "M1" '''macro one''' }}
		{{ macro "M2" '''macro two calls {{M1}}'''}}
		}}}

		{{ M2 }}
		`, nil, nil)

	rt.Parse()
	actual, _ := rt.Expand(nil)
	util.LineContentsEqual(t, actual,
		`
		macro two calls macro one
	`)
}
