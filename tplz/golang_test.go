package tplz_test

import (
	"testing"
	ttpl "text/template"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-lib/fixturez"
	"github.com/ibrt/golang-lib/tplz"
)

type GolangSuite struct {
	// intentionally empty
}

func TestGolangSuite(t *testing.T) {
	fixturez.RunSuite(t, &GolangSuite{})
}

func (*GolangSuite) TestExecuteGolang(g *WithT) {
	okTpl, err := ttpl.New("").Parse("package main\nimport \"fmt\"\nfunc main() { fmt.Println(\"{{ . }}\") }")
	g.Expect(err).To(Succeed())

	errTpl, err := ttpl.New("").Parse(`{{ template "x" }}`)
	g.Expect(err).To(Succeed())

	goErrTpl, err := ttpl.New("").Parse("package main\nfuncmain() { fmt.Println(\"{{ . }}\") }")
	g.Expect(err).To(Succeed())

	g.Expect(tplz.ExecuteGolang(okTpl, "Hello World")).
		To(Equal([]byte("package main\n\nimport \"fmt\"\n\nfunc main() { fmt.Println(\"Hello World\") }\n")))

	g.Expect(tplz.ExecuteGolang(errTpl, nil)).Error().
		To(MatchError(`template: :1:12: executing "" at <{{template "x"}}>: template "x" not defined`))

	g.Expect(tplz.ExecuteGolang(goErrTpl, "Hello World")).Error().
		To(MatchError("2:1: expected declaration, found funcmain"))
}

func (*GolangSuite) TestMustExecuteGolang(g *WithT) {
	okTpl, err := ttpl.New("").Parse("package main\nimport \"fmt\"\nfunc main() { fmt.Println(\"{{ . }}\") }")
	g.Expect(err).To(Succeed())

	errTpl, err := ttpl.New("").Parse(`{{ template "x" }}`)
	g.Expect(err).To(Succeed())

	goErrTpl, err := ttpl.New("").Parse("package main\nfuncmain() { fmt.Println(\"{{ . }}\") }")
	g.Expect(err).To(Succeed())

	g.Expect(func() {
		g.Expect(tplz.MustExecuteGolang(okTpl, "Hello World")).
			To(Equal([]byte("package main\n\nimport \"fmt\"\n\nfunc main() { fmt.Println(\"Hello World\") }\n")))
	}).ToNot(Panic())

	g.Expect(func() {
		tplz.MustExecuteGolang(errTpl, nil)
	}).To(PanicWith(MatchError(`template: :1:12: executing "" at <{{template "x"}}>: template "x" not defined`)))

	g.Expect(func() {
		tplz.MustExecuteGolang(goErrTpl, "Hello World")
	}).To(PanicWith(MatchError("2:1: expected declaration, found funcmain")))
}

func (*GolangSuite) TestParseAndExecuteGolang(g *WithT) {
	g.Expect(tplz.ParseAndExecuteGolang("package main\nimport \"fmt\"\nfunc main() { fmt.Println(\"{{ . }}\") }", "Hello World")).
		To(Equal([]byte("package main\n\nimport \"fmt\"\n\nfunc main() { fmt.Println(\"Hello World\") }\n")))

	g.Expect(tplz.ParseAndExecuteGolang("{{ bad }}", "Hello World")).Error().
		To(MatchError(`template: :1: function "bad" not defined`))

	g.Expect(tplz.ParseAndExecuteGolang(`{{ template "x" }}`, nil)).Error().
		To(MatchError(`template: :1:12: executing "" at <{{template "x"}}>: template "x" not defined`))

	g.Expect(tplz.ParseAndExecuteGolang("package main\nfuncmain() { fmt.Println(\"{{ . }}\") }", "Hello World")).Error().
		To(MatchError("2:1: expected declaration, found funcmain"))
}

func (*GolangSuite) TestMustParseAndExecuteGolang(g *WithT) {
	g.Expect(func() {
		g.Expect(tplz.MustParseAndExecuteGolang("package main\nimport \"fmt\"\nfunc main() { fmt.Println(\"{{ . }}\") }", "Hello World")).
			To(Equal([]byte("package main\n\nimport \"fmt\"\n\nfunc main() { fmt.Println(\"Hello World\") }\n")))
	}).ToNot(Panic())

	g.Expect(func() {
		tplz.MustParseAndExecuteGolang("{{ bad }}", "Hello World")
	}).To(PanicWith(MatchError(`template: :1: function "bad" not defined`)))

	g.Expect(func() {
		tplz.MustParseAndExecuteGolang(`{{ template "x" }}`, nil)
	}).To(PanicWith(MatchError(`template: :1:12: executing "" at <{{template "x"}}>: template "x" not defined`)))

	g.Expect(func() {
		tplz.MustParseAndExecuteGolang("package main\nfuncmain() { fmt.Println(\"{{ . }}\") }", "Hello World")
	}).To(PanicWith(MatchError("2:1: expected declaration, found funcmain")))
}
