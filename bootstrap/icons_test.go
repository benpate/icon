package bootstrap

import (
	"errors"
	"strings"
	"testing"

	"github.com/benpate/derp"
	"github.com/benpate/derp/plugins"
	"github.com/stretchr/testify/require"
)

// failingWriter is an io.Writer whose every write fails with its err
type failingWriter struct {
	err error
}

// Write fails without writing anything
func (writer failingWriter) Write([]byte) (int, error) {
	return 0, writer.err
}

// captureReporter records every error that derp reports
type captureReporter struct {
	errors *[]error
}

// Report records the error
func (reporter captureReporter) Report(err error) {
	*reporter.errors = append(*reporter.errors, err)
}

// captureReports routes derp's reports into the returned slice until the test ends
func captureReports(t *testing.T) *[]error {
	t.Helper()

	reported := make([]error, 0)
	derp.SetPlugins(captureReporter{errors: &reported})
	t.Cleanup(func() { derp.SetPlugins(plugins.JSON{}) })

	return &reported
}

// TestProvider_Get verifies that common names translate to Bootstrap names, and others pass through
func TestProvider_Get(t *testing.T) {

	provider := Provider{}

	test := func(name string, expected string) {
		t.Helper()
		require.Equal(t, expected, provider.Get(name), "icon %q", name)
	}

	test("add", `<i class="bi bi-plus"></i>`)
	test("delete", `<i class="bi bi-trash"></i>`)
	test("shopping-cart", `<i class="bi bi-cart3"></i>`)
	test("star", `<i class="bi bi-star"></i>`)
	test("", `<i class="bi bi-"></i>`)
}

// TestProvider_Write verifies that Write writes the same markup that Get returns, and reports nothing
func TestProvider_Write(t *testing.T) {

	reported := captureReports(t)
	provider := Provider{}

	var builder strings.Builder
	provider.Write("edit", &builder)

	require.Equal(t, provider.Get("edit"), builder.String())
	require.Empty(t, *reported)
}

// TestProvider_Write_Failed verifies that a failed write is reported once, with its location and cause
func TestProvider_Write_Failed(t *testing.T) {

	reported := captureReports(t)
	writeError := errors.New("write failed")

	Provider{}.Write("edit", failingWriter{err: writeError})

	require.Len(t, *reported, 1)
	require.ErrorIs(t, (*reported)[0], writeError)
	require.Equal(t, "bootstrap.Provider.Write", derp.Location((*reported)[0]))
}
