package sparql

type SparqlResult struct {
	Head    Head
	Results Results
}

type Head struct {
	Vars []string
}

type Results struct {
	Bindings []Binding
}

type Binding map[string]struct {
	Type  string
	Value string
}

func (sr *SparqlResult) Vars() []string {
	return sr.Head.Vars
}

func (sr *SparqlResult) Bindings() []Binding {
	return sr.Results.Bindings
}

func (sr *SparqlResult) BoundValues(bindingIndex int) []string {
	variables := sr.Vars()
	values := make([]string, len(variables))
	binding := sr.Bindings()[bindingIndex]
	for columnIndex, varName := range variables {
		values[columnIndex] = binding[varName].Value
	}
	return values
}
