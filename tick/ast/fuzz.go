// +build gofuzz
package ast

// Fuzzing definitions to detect edge/corner cases via https://github.com/dvyukov/go-fuzz.
//
// Currently, go-fuzz doesn't offer a way to automatically detect one of multiple Fuzz functions,
// so you will need to specify which function to use.
// First, build the instrumented package:
//
//     go-fuzz-build -func FuzzParse -o /tmp/FuzzParse.zip github.com/influxdata/kapacitor/tick/ast
//
// Then, run the tests:
//
//     go-fuzz -bin=/tmp/FuzzParse.zip -workdir=$GOPATH/src/github.com/influxdata/kapacitor/tick/ast/fuzz-data/Parse.

// The return value of the fuzz functions affects how the fuzzer finds new cases.
const (
	// "Correct" input that was correctly parsed.
	fuzzInteresting = 1

	// Not really interesting input.
	fuzzBoring = 0

	// Do not use this input, even if it results in new coverage.
	// (When would you use this?)
	fuzzIgnore = -1
)

func FuzzParse(data []byte) int {
	_, err := Parse(string(data))
	if err != nil {
		// There are practically infinite parser error conditions.
		// Don't ever add parsing errors to the corpus.
		return fuzzIgnore
	}

	return fuzzInteresting
}
