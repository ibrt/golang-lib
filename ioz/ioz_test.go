package ioz_test

import (
	"io"
	"strings"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-lib/fixturez"
	"github.com/ibrt/golang-lib/ioz"
)

type Suite struct {
	// intentionally empty
}

func TestSuite(t *testing.T) {
	fixturez.RunSuite(t, &Suite{})
}

func (*Suite) TestCountingReader(g *WithT) {
	s := strings.Repeat("x", 1024)
	c := ioz.NewCountingReader(strings.NewReader(s))
	g.Expect(c.Count()).To(BeNumerically("==", 0))
	n, err := io.Copy(io.Discard, c)
	g.Expect(err).To(Succeed())
	g.Expect(n).To(BeNumerically("==", 1024))
	g.Expect(c.Count()).To(BeNumerically("==", 1024))
}
