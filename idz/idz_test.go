package idz_test

import (
	"testing"

	"github.com/google/uuid"
	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-lib/fixturez"
	"github.com/ibrt/golang-lib/idz"
)

type Suite struct {
	// intentionally empty
}

func TestSuite(t *testing.T) {
	fixturez.RunSuite(t, &Suite{})
}

func (s *Suite) TestMustNewRandomID(g *WithT) {
	id := idz.MustNewRandomID()
	g.Expect(id).ToNot(BeEmpty())
	g.Expect(uuid.Parse(id)).Error().To(Succeed())
}

func (s *Suite) TestIsValidID(g *WithT) {
	g.Expect(idz.IsValidID(idz.MustNewRandomID())).To(BeTrue())
	g.Expect(idz.IsValidID(idz.MustNewRandomID() + "x")).To(BeFalse())
	g.Expect(idz.IsValidID("")).To(BeFalse())
}
