package cfz_test

import (
	"testing"

	cf "github.com/awslabs/goformation/v7/cloudformation"
	cftags "github.com/awslabs/goformation/v7/cloudformation/tags"
	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-lib/cfz"
	"github.com/ibrt/golang-lib/fixturez"
)

type TemplateUtilsSuite struct {
	// intentionally empty
}

func TestTemplateUtilsSuite(t *testing.T) {
	fixturez.RunSuite(t, &TemplateUtilsSuite{})
}

func (*TemplateUtilsSuite) TestTemplateUtils(g *WithT) {
	u := cfz.NewTemplateUtils("tn")

	g.Expect(u.GetStackName("test-stack")).To(Equal("tn-test-stack"))
	g.Expect(u.GetResourceName("test-resource")).To(Equal("tn-test-resource"))

	g.Expect(u.GenerateResourceTags("test-resource")).To(Equal([]cftags.Tag{
		{
			Key:   "Name",
			Value: "tn-test-resource",
		},
	}))

	tpl := cf.NewTemplate()
	u.AddOutputRef(tpl, "test-resource")
	u.AddOutputAtt(tpl, "test-resource", "TestAttribute")

	eTpl := cf.NewTemplate()
	eTpl.Outputs["TestResourceRef"] = cf.Output{Value: "eyAiUmVmIjogIlRlc3RSZXNvdXJjZSIgfQ=="}
	eTpl.Outputs["TestResourceAttTestAttribute"] = cf.Output{Value: "eyAiRm46OkdldEF0dCIgOiBbICJUZXN0UmVzb3VyY2UiLCAiVGVzdEF0dHJpYnV0ZSIgXSB9"}
	g.Expect(tpl).To(Equal(eTpl))
}
