package cfz_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-lib/cfz"
	"github.com/ibrt/golang-lib/fixturez"
)

type ResourcesSuite struct {
	// intentionally empty
}

func TestResourcesSuite(t *testing.T) {
	fixturez.RunSuite(t, &ResourcesSuite{})
}

func (*ResourcesSuite) TestResource(g *WithT) {
	res := cfz.Resource("test-resource")
	g.Expect(res.Ref()).To(Equal("eyAiUmVmIjogIlRlc3RSZXNvdXJjZSIgfQ=="))
	g.Expect(res.GetAtt("Arn")).To(Equal("eyAiRm46OkdldEF0dCIgOiBbICJUZXN0UmVzb3VyY2UiLCAiQXJuIiBdIH0="))
	g.Expect(res.Logical()).To(Equal("TestResource"))
	g.Expect(res.Physical()).To(Equal("test-resource"))
}
