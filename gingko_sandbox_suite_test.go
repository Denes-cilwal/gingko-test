package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
- In Go, a dot-import allows you to access the exported identifiers of a package directly without needing to prefix them with the package name
- RegisterFailHandler(Fail) is the single line of glue code connecting Ginkgo to Gomega.
- If we were to avoid dot-imports this would read as gomega.RegisterFailHandler(ginkgo.Fail)
- what we're doing here is telling our matcher library (Gomega) which function to call (Ginkgo's Fail) in the event a failure is detected.
*/

func TestGingkoSandbox(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GingkoSandbox Suite")
}
