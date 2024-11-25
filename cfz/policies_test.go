package cfz_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-lib/cfz"
	"github.com/ibrt/golang-lib/fixturez"
)

type PoliciesSuite struct {
	// intentionally empty
}

func TestPoliciesSuite(t *testing.T) {
	fixturez.RunSuite(t, &PoliciesSuite{})
}

func (*PoliciesSuite) TestAssumeRolePolicyDocument(g *WithT) {
	g.Expect(cfz.NewAssumeRolePolicyDocument("test-service")).To(Equal(map[string]any{
		"Version": "2012-10-17",
		"Statement": []map[string]any{
			{
				"Effect": "Allow",
				"Principal": map[string]any{
					"Service": "test-service",
				},
				"Action": "sts:AssumeRole",
			},
		},
	}))
}

func (*PoliciesSuite) TestPolicyDocument(g *WithT) {
	g.Expect(
		cfz.NewPolicyDocument(
			cfz.NewPolicyStatement().
				AddActions("test-action").
				AddResources("test-resource"),
		)).
		To(Equal(map[string]any{
			"Version": "2012-10-17",
			"Statement": []any{
				map[string]any{
					"Effect":   "Allow",
					"Action":   []string{"test-action"},
					"Resource": []string{"test-resource"},
				},
			},
		}))
}

func (*PoliciesSuite) TestPolicyStatement(g *WithT) {
	g.Expect(
		cfz.NewPolicyStatement().
			AddActions("test-action").
			AddResources("test-resource").
			Build()).
		To(Equal(map[string]any{
			"Effect":   "Allow",
			"Action":   []string{"test-action"},
			"Resource": []string{"test-resource"},
		}))

	g.Expect(
		cfz.NewPolicyStatement().
			SetCurrentRootAccountPrincipal().
			AddActions("test-action").
			AddResources("test-resource").
			Build()).
		To(Equal(map[string]any{
			"Effect":   "Allow",
			"Action":   []string{"test-action"},
			"Resource": []string{"test-resource"},
			"Principal": map[string]any{
				"AWS": "eyAiRm46OlN1YiIgOiAiYXJuOmF3czppYW06OiR7QVdTOjpBY2NvdW50SWR9OnJvb3QiIH0=",
			},
		}))

	g.Expect(
		cfz.NewPolicyStatement().
			SetWildcardPrincipal().
			AddActions("test-action").
			AddResources("test-resource").
			Build()).
		To(Equal(map[string]any{
			"Effect":    "Allow",
			"Action":    []string{"test-action"},
			"Resource":  []string{"test-resource"},
			"Principal": "*",
		}))

	g.Expect(
		cfz.NewPolicyStatement().
			SetServicePrincipal("test-service").
			AddActions("test-action").
			AddResources("test-resource").
			Build()).
		To(Equal(map[string]any{
			"Effect":   "Allow",
			"Action":   []string{"test-action"},
			"Resource": []string{"test-resource"},
			"Principal": map[string]any{
				"Service": "test-service",
			},
		}))

	g.Expect(
		cfz.NewPolicyStatement().
			SetAnyRootAccountPrincipal().
			AddActions("test-action").
			AddResources("test-resource").
			Build()).
		To(Equal(map[string]any{
			"Effect":   "Allow",
			"Action":   []string{"test-action"},
			"Resource": []string{"test-resource"},
			"Principal": map[string]any{
				"AWS": "*",
			},
		}))
}
