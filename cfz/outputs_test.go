package cfz_test

import (
	"testing"

	awscft "github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-lib/cfz"
	"github.com/ibrt/golang-lib/fixturez"
	"github.com/ibrt/golang-lib/memz"
)

type OutputsSuite struct {
	// intentionally empty
}

func TestOutputsSuite(t *testing.T) {
	fixturez.RunSuite(t, &OutputsSuite{})
}

func (*OutputsSuite) TestOutputs(g *WithT) {
	o := cfz.NewOutputs(&awscft.Stack{
		Outputs: []awscft.Output{
			{
				OutputKey:   memz.Ptr("TestResourceRef"),
				OutputValue: memz.Ptr("tov-1"),
			},
			{
				OutputKey:   memz.Ptr("TestResourceAttTestAttribute"),
				OutputValue: memz.Ptr("tov-2"),
			},
		},
	})

	g.Expect(o.Ref("test-resource")).To(Equal("tov-1"))
	g.Expect(o.Att("test-resource", "TestAttribute")).To(Equal("tov-2"))

	g.Expect(func() {
		o.Ref("unknown-res")
	}).To(PanicWith(MatchError("no such export: ref for unknown-res")))

	g.Expect(func() {
		o.Att("unknown-res", "UnknownAtt")
	}).To(PanicWith(MatchError("no such export: att UnknownAtt for unknown-res")))
}
