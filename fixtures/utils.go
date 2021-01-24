package fixtures

import (
	"encoding/json"

	"github.com/ibrt/golang-lib/errors"
)

// TestingT describes parts of a *testing.T struct.
type TestingT interface {
	Helper()
	Log(...interface{})
	Fail()
	FailNow()
}

// RequireNoError is like require.NoError, but properly formats attached error stack traces.
func RequireNoError(t TestingT, err error) {
	t.Helper()
	if err == nil {
		return
	}

	buf, _ := json.MarshalIndent(errors.ToResponse(err), "", "  ")
	t.Log(string(buf))
	t.FailNow()
}

// AssertNoError is like assert.NoError, but properly formats attached error stack traces.
func AssertNoError(t TestingT, err error) {
	t.Helper()
	if err == nil {
		return
	}

	buf, _ := json.MarshalIndent(errors.ToResponse(err), "", "  ")
	t.Log(string(buf))
	t.Fail()
}
