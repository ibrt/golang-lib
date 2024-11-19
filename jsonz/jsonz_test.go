package jsonz_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-lib/fixturez"
	"github.com/ibrt/golang-lib/jsonz"
)

type Suite struct {
	// intentionally empty
}

func TestSuite(t *testing.T) {
	fixturez.RunSuite(t, &Suite{})
}

func (*Suite) TestMustMarshal(g *WithT) {
	g.Expect(func() {
		g.Expect(jsonz.MustMarshal(map[string]int{"a": 1})).To(Equal([]byte(`{"a":1}`)))
	}).ToNot(Panic())

	g.Expect(func() {
		jsonz.MustMarshal(func() {})
	}).To(PanicWith(MatchError("json: unsupported type: func()")))
}

func (*Suite) TestMustMarshalString(g *WithT) {
	g.Expect(func() {
		g.Expect(jsonz.MustMarshalString(map[string]int{"a": 1})).To(Equal(`{"a":1}`))
	}).ToNot(Panic())

	g.Expect(func() {
		jsonz.MustMarshalString(func() {})
	}).To(PanicWith(MatchError("json: unsupported type: func()")))
}

func (*Suite) TestMustMarshalIndent(g *WithT) {
	g.Expect(func() {
		g.Expect(jsonz.MustMarshalIndent(map[string]int{"a": 1}, "#", "  ")).To(Equal([]byte("{\n#  \"a\": 1\n#}")))
	}).ToNot(Panic())

	g.Expect(func() {
		jsonz.MustMarshalIndent(func() {}, "#", "  ")
	}).To(PanicWith(MatchError("json: unsupported type: func()")))
}

func (*Suite) TestMustMarshalIndentString(g *WithT) {
	g.Expect(func() {
		g.Expect(jsonz.MustMarshalIndentString(map[string]int{"a": 1}, "#", "  ")).To(Equal("{\n#  \"a\": 1\n#}"))
	}).ToNot(Panic())

	g.Expect(func() {
		jsonz.MustMarshalIndentString(func() {}, "#", "  ")

	}).To(PanicWith(MatchError("json: unsupported type: func()")))
}

func (*Suite) TestMustMarshalIndentDefault(g *WithT) {
	g.Expect(func() {
		g.Expect(jsonz.MustMarshalIndentDefault(map[string]int{"a": 1})).To(Equal([]byte("{\n  \"a\": 1\n}")))
	}).ToNot(Panic())

	g.Expect(func() {
		jsonz.MustMarshalIndentDefault(func() {})
	}).To(PanicWith(MatchError("json: unsupported type: func()")))
}

func (*Suite) TestMustMarshalIndentDefaultString(g *WithT) {
	g.Expect(func() {
		g.Expect(jsonz.MustMarshalIndentDefaultString(map[string]int{"a": 1})).To(Equal("{\n  \"a\": 1\n}"))
	}).ToNot(Panic())

	g.Expect(func() {
		jsonz.MustMarshalIndentDefaultString(func() {})
	}).To(PanicWith(MatchError("json: unsupported type: func()")))
}

func (*Suite) TestMustIndent(g *WithT) {
	g.Expect(func() {
		g.Expect(jsonz.MustIndent([]byte(`{"a":1}`), "#", "  ")).To(Equal([]byte("{\n#  \"a\": 1\n#}")))
	}).ToNot(Panic())

	g.Expect(func() {
		jsonz.MustIndent([]byte(`bad`), "#", "  ")
	}).To(PanicWith(MatchError("invalid character 'b' looking for beginning of value")))
}

func (*Suite) TestMustIndentString(g *WithT) {
	g.Expect(func() {
		g.Expect(jsonz.MustIndentString([]byte(`{"a":1}`), "#", "  ")).To(Equal("{\n#  \"a\": 1\n#}"))
	}).ToNot(Panic())

	g.Expect(func() {
		jsonz.MustIndentString([]byte(`bad`), "#", "  ")
	}).To(PanicWith(MatchError("invalid character 'b' looking for beginning of value")))
}

func (*Suite) TestMustIndentDefault(g *WithT) {
	g.Expect(func() {
		g.Expect(jsonz.MustIndentDefault([]byte(`{"a":1}`))).To(Equal([]byte("{\n  \"a\": 1\n}")))
	}).ToNot(Panic())

	g.Expect(func() {
		jsonz.MustIndentDefault([]byte(`bad`))
	}).To(PanicWith(MatchError("invalid character 'b' looking for beginning of value")))
}

func (*Suite) TestMustIndentDefaultString(g *WithT) {
	g.Expect(func() {
		g.Expect(jsonz.MustIndentDefaultString([]byte(`{"a":1}`))).To(Equal("{\n  \"a\": 1\n}"))
	}).ToNot(Panic())

	g.Expect(func() {
		jsonz.MustIndentDefaultString([]byte(`bad`))
	}).To(PanicWith(MatchError("invalid character 'b' looking for beginning of value")))
}

func (*Suite) TestUnmarshal(g *WithT) {
	type testStruct struct {
		K string `json:"k"`
	}

	g.Expect(jsonz.Unmarshal[testStruct]([]byte(`{"k": "v"}`))).
		To(Equal(&testStruct{K: "v"}))

	g.Expect(jsonz.Unmarshal[testStruct]([]byte(`bad`))).Error().
		To(MatchError("invalid character 'b' looking for beginning of value"))
}

func (*Suite) TestMustUnmarshal(g *WithT) {
	type testStruct struct {
		K string `json:"k"`
	}

	g.Expect(func() {
		g.Expect(jsonz.MustUnmarshal[testStruct]([]byte(`{"k": "v"}`))).
			To(Equal(&testStruct{K: "v"}))
	}).ToNot(Panic())

	g.Expect(func() {
		jsonz.MustUnmarshal[testStruct]([]byte(`bad`))
	}).To(PanicWith(MatchError("invalid character 'b' looking for beginning of value")))
}

func (*Suite) TestUnmarshalString(g *WithT) {
	type testStruct struct {
		K string `json:"k"`
	}

	g.Expect(jsonz.UnmarshalString[testStruct](`{"k": "v"}`)).
		To(Equal(&testStruct{K: "v"}))

	g.Expect(jsonz.UnmarshalString[testStruct](`bad`)).Error().
		To(MatchError("invalid character 'b' looking for beginning of value"))
}

func (*Suite) TestMustUnmarshalString(g *WithT) {
	type testStruct struct {
		K string `json:"k"`
	}

	g.Expect(func() {
		g.Expect(jsonz.MustUnmarshalString[testStruct](`{"k": "v"}`)).
			To(Equal(&testStruct{K: "v"}))
	}).ToNot(Panic())

	g.Expect(func() {
		jsonz.MustUnmarshalString[testStruct](`bad`)
	}).To(PanicWith(MatchError("invalid character 'b' looking for beginning of value")))
}
