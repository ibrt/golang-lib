package cfz_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-lib/cfz"
	"github.com/ibrt/golang-lib/fixturez"
)

type AttributesSuite struct {
	// intentionally empty
}

func TestAttributesSuite(t *testing.T) {
	fixturez.RunSuite(t, &AttributesSuite{})
}

func (*AttributesSuite) TestAttribute(g *WithT) {
	g.Expect(cfz.Attribute("CanonicalHostedZoneID").Name()).To(Equal("CanonicalHostedZoneID"))
	g.Expect(cfz.Attribute("CanonicalHostedZoneID").Logical()).To(Equal("CanonicalHostedZoneID"))
	g.Expect(cfz.Attribute("CanonicalHostedZoneID").Physical()).To(Equal("canonical-hosted-zone-id"))

	g.Expect(cfz.Attribute("Endpoint.Address").Name()).To(Equal("Endpoint.Address"))
	g.Expect(cfz.Attribute("Endpoint.Address").Logical()).To(Equal("EndpointAddress"))
	g.Expect(cfz.Attribute("Endpoint.Address").Physical()).To(Equal("endpoint-address"))
}
