package errors_test

import (
	"net/http"
	"strings"
	"testing"

	"github.com/ibrt/golang-lib/errors"
	"github.com/stretchr/testify/require"
)

func TestToResponse(t *testing.T) {
	resp := errors.ToResponse(errors.Errorf("some error", errors.Prefix("prefix"), errors.ID("id"), errors.StatusCode(http.StatusUnauthorized)))
	require.Equal(t, errors.StatusUnauthorized, resp.StatusCode)
	require.Equal(t, errors.ID("id"), resp.ErrorID)
	require.Equal(t, "prefix: some error", resp.Message)
	require.NotEmpty(t, resp.StackTrace)
	require.True(t, strings.HasPrefix(resp.StackTrace[0], "errors_test.TestToResponse"))
}
