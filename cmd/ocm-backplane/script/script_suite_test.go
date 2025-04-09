package script

import (
	"io"
	"strings"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestScriptCmdSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ManagedJob Test Suite")
}

func MakeIoReader(s string) io.ReadCloser {
	r := io.NopCloser(strings.NewReader(s)) // r type is io.ReadCloser
	return r
}
