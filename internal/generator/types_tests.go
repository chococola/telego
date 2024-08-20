package main

import (
	"fmt"
	"regexp"
	"strings"
)

const funcPattern = `
func \(\w* \*(\w*)\) (\w*)\(\) string {
	return (\w*)
}
`

var funcRegexp = regexp.MustCompile(funcPattern)

func generateTypesTests(types string) {
	typesTestsFile := openFile(generatedTypesTestsFilename)
	defer func() { _ = typesTestsFile.Close() }()

	data := strings.Builder{}

	data.WriteString(`package telego

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestTypesInterfaces(t *testing.T) {
`)

	funcs := funcRegexp.FindAllStringSubmatch(types, -1)

	logInfof("Func count: %d", len(funcs))

	for _, f := range funcs {
		funcType := f[1]
		funcName := f[2]
		funcReturn := f[3]

		data.WriteString(fmt.Sprintf("\tassert.Implements(t, (*INTERFACE)(nil), &%s{})\n", funcType))
		data.WriteString(fmt.Sprintf("\tassert.Equal(t, %s, (&%s{}).%s())\n\n", funcReturn, funcType, funcName))
	}

	data.WriteString("}\n")

	_, err := typesTestsFile.WriteString(data.String())
	exitOnErr(err)

	formatFile(typesTestsFile.Name())
}
