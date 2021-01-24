package errors_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ibrt/golang-lib/errors"
	"github.com/stretchr/testify/require"
)

func TestWrap(t *testing.T) {
	err := errors.Wrap(fmt.Errorf("test error"))
	require.NotNil(t, err)
	require.Equal(t, "test error", err.Error())
	require.True(t, strings.HasPrefix(errors.FormatStackTrace(errors.GetCallers(err))[0], "errors_test.TestWrap"))
	require.PanicsWithValue(t, "nil error", func() { _ = errors.Wrap(nil) })
	require.Equal(t, err, errors.Wrap(err))
	require.Equal(t, errors.ID(""), errors.GetID(err))
	require.Equal(t, errors.Metadata{}, errors.GetMetadata(err))
}

func TestMaybeWrap(t *testing.T) {
	err := errors.MaybeWrap(fmt.Errorf("test error"))
	require.NotNil(t, err)
	require.Equal(t, "test error", err.Error())
	require.True(t, strings.HasPrefix(errors.FormatStackTrace(errors.GetCallers(err))[0], "errors_test.TestMaybeWrap"))
	require.Nil(t, errors.MaybeWrap(nil))
}

func TestMustWrap(t *testing.T) {
	require.PanicsWithError(t, "test error", func() { errors.MustWrap(fmt.Errorf("test error")) })
	require.PanicsWithValue(t, "nil error", func() { errors.MustWrap(nil) })
}

func TestMaybeMustWrap(t *testing.T) {
	require.PanicsWithError(t, "test error", func() { errors.MaybeMustWrap(fmt.Errorf("test error")) })
	require.NotPanics(t, func() { errors.MaybeMustWrap(nil) })
}

func TestWrapRecover(t *testing.T) {
	err := errors.WrapRecover("test error")
	require.NotNil(t, err)
	require.Equal(t, "test error", err.Error())
	err = errors.WrapRecover(fmt.Errorf("test error"))
	require.NotNil(t, err)
	require.Equal(t, "test error", err.Error())
	require.NotNil(t, err)
	err = errors.WrapRecover(errors.Errorf("test error"))
	require.NotNil(t, err)
	require.Equal(t, "test error", err.Error())
	require.NotNil(t, err)
	require.PanicsWithValue(t, "nil recover", func() { _ = errors.WrapRecover(nil) })
}

func TestMaybeWrapRecover(t *testing.T) {
	err := errors.MaybeWrapRecover("test error")
	require.NotNil(t, err)
	require.Equal(t, "test error", err.Error())
	require.NotNil(t, err)
	err = errors.MaybeWrapRecover(fmt.Errorf("test error"))
	require.NotNil(t, err)
	require.Equal(t, "test error", err.Error())
	err = errors.MaybeWrapRecover(errors.Errorf("test error"))
	require.NotNil(t, err)
	require.Equal(t, "test error", err.Error())
	require.NotNil(t, err)
	require.Nil(t, errors.MaybeWrapRecover(nil))
}

func TestErrorf(t *testing.T) {
	err := errors.Errorf("test error")
	require.NotNil(t, t, err)
	require.Equal(t, "test error", err.Error())
	require.True(t, strings.HasPrefix(errors.FormatStackTrace(errors.GetCallers(err))[0], "errors_test.TestErrorf"))
	err = errors.Errorf("format %s", errors.Prefix("prefix"), errors.A("xxx"), errors.ID("id"), errors.M("k", "v"))
	require.NotNil(t, t, err)
	require.Equal(t, "prefix: format xxx", err.Error())
	require.Equal(t, errors.ID("id"), errors.GetID(err))
	require.Equal(t, errors.Metadata{"k": "v"}, errors.GetMetadata(err))
}

func TestMustErrorf(t *testing.T) {
	require.PanicsWithError(t, "test error", func() { errors.MustErrorf("test error") })
}

func TestAssert(t *testing.T) {
	require.NotPanics(t, func() { errors.Assertf(true, "test error") })
	require.PanicsWithError(t, "test error: value", func() { errors.Assertf(false, "test error: %v", errors.Args{"value"}) })
}

type testCloser struct {
	closed bool
}

// Close implements io.Closer.
func (c *testCloser) Close() error {
	c.closed = true
	return nil
}

func TestIgnoreClose(t *testing.T) {
	tc := &testCloser{}
	require.False(t, tc.closed)
	errors.IgnoreClose(tc)
	require.True(t, tc.closed)
}

func TestUnwrap(t *testing.T) {
	require.Nil(t, errors.Unwrap(nil))
	err := fmt.Errorf("test error")
	ret := errors.Unwrap(err)
	require.Equal(t, ret, err)
	ret = errors.Unwrap(errors.Wrap(err))
	require.Equal(t, ret, err)
}

func TestSafe(t *testing.T) {
	require.EqualError(t, errors.Safe(func() error { panic(errors.Errorf("test error")) })(), "test error")
	require.EqualError(t, errors.Safe(func() error { return errors.Errorf("test error") })(), "test error")
}

func TestID(t *testing.T) {
	require.Equal(t, errors.ID(""), errors.GetID(errors.Errorf("test error")))
	require.Equal(t, errors.ID(""), errors.GetID(fmt.Errorf("test error")))
	err := errors.Errorf("test error", errors.ID("id"))
	require.NotNil(t, err)
	require.Equal(t, errors.ID("id"), errors.GetID(err))
	id := errors.ID("id")
	require.Equal(t, "id", id.String())
	err = errors.Errorf("test error", id)
	require.NotNil(t, err)
	require.Equal(t, id, errors.GetID(err))
	require.True(t, id.In(err))
	require.False(t, id.In(errors.Errorf("test error")))
	require.False(t, id.In(fmt.Errorf("test error")))
}

func TestStatusCode(t *testing.T) {
	require.Equal(t, errors.StatusCode(0), errors.GetStatusCode(errors.Errorf("test error")))
	require.Equal(t, errors.StatusInternalServerError, errors.GetStatusCodeOr500(errors.Errorf("test error")))
	require.Equal(t, errors.StatusCode(0), errors.GetStatusCode(fmt.Errorf("test error")))
	require.Equal(t, errors.StatusInternalServerError, errors.GetStatusCodeOr500(fmt.Errorf("test error")))
	err := errors.Errorf("test error", errors.StatusCode(50))
	require.NotNil(t, err)
	require.Equal(t, errors.StatusCode(50), errors.GetStatusCode(err))
	err = errors.Errorf("test error", errors.StatusNotFound)
	require.NotNil(t, err)
	require.Equal(t, errors.StatusNotFound, errors.GetStatusCode(err))
	require.Equal(t, errors.StatusNotFound, errors.GetStatusCodeOr500(err))
	statusCode := errors.StatusNotFound
	require.Equal(t, 404, statusCode.Int())
	require.Equal(t, "Not Found", statusCode.String())
	err = errors.Errorf("test error", statusCode)
	require.NotNil(t, err)
	require.Equal(t, statusCode, errors.GetStatusCode(err))
	require.True(t, statusCode.In(err))
	require.False(t, statusCode.In(errors.Errorf("test error")))
	require.False(t, statusCode.In(fmt.Errorf("test error")))
}

func TestMetadata(t *testing.T) {
	require.Equal(t, errors.Metadata{}, errors.GetMetadata(errors.Errorf("test error")))
	require.Equal(t, errors.Metadata{}, errors.GetMetadata(fmt.Errorf("test error")))
	err := errors.Errorf("test error",
		errors.Metadata{"k1": "v1", "k2": 2},
		errors.M("k3", "v3"),
		errors.Metadata{"k4": "v4", "k5": "v5"},
		errors.Metadata{},
		errors.M("k5", "over"))
	require.Equal(t, errors.Metadata{
		"k1": "v1",
		"k2": 2,
		"k3": "v3",
		"k4": "v4",
		"k5": "over",
	}, errors.GetMetadata(err))
	require.Equal(t, "v1", errors.GetMetadata(err).Get("k1"))
	require.Equal(t, nil, errors.GetMetadata(err).Get("unknown"))
	require.Equal(t, "v1", errors.GetMetadata(err).GetString("k1"))
	require.Equal(t, "", errors.GetMetadata(err).GetString("k2"))
	require.Equal(t, "", errors.GetMetadata(err).GetString("unknown"))
	require.Nil(t, errors.Metadata(nil).Get("unknown"))
}

func TestPrefix(t *testing.T) {
	require.Equal(t, "p2 20: p1 10: test error",
		errors.Errorf("test error",
			errors.Prefix("p1 %v", 10),
			errors.Prefix("p2 %v", 20)).Error())
}

func TestCallers(t *testing.T) {
	require.True(t, strings.HasPrefix(errors.FormatStackTrace(errors.GetCallers(fmt.Errorf("test error ")))[0],
		"errors_test.TestCallers"))
	require.Empty(t, errors.GetCallers(errors.Errorf("test error", errors.Skip(1000))))
}
