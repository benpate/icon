package bootstrap

import (
	"errors"
	"strings"
	"testing"

	"github.com/benpate/derp"
	"github.com/benpate/derp/plugins"
	"github.com/benpate/icon"
	"github.com/stretchr/testify/require"
)

// Provider must keep satisfying the interface that consumers depend on
var _ icon.Provider = Provider{}

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

// TestProvider_Get verifies that every generic name translates to its Bootstrap name
func TestProvider_Get(t *testing.T) {

	provider := Provider{}

	test := func(name string, bootstrapName string) {
		t.Helper()
		require.Equal(t, `<i class="bi bi-`+bootstrapName+`"></i>`, provider.Get(name), "icon %q", name)
	}

	test("add",             "plus")
	test("cancel",          "x-lg")
	test("check-circle",    "check-circle-fill")
	test("chevron-left",    "chevron-left")
	test("chevron-right",   "chevron-right")
	test("delete",          "trash")
	test("edit",            "pencil-square")
	test("email",           "envelope")
	test("file",            "file-earmark")
	test("flag",            "flag-fill")
	test("folder",          "folder")
	test("grip-vertical",   "grip-vertical")
	test("grip-horizontal", "grip-horizontal")
	test("home",            "house")
	test("info",            "info-circle")
	test("link",            "link-45deg")
	test("loading",         "arrow-clockwise")
	test("login",           "box-arrow-in-right")
	test("save",            "check-lg")
	test("settings",        "gear-fill")
	test("server",          "hdd-stack")
	test("reply",           "reply-fill")
	test("unlink",          "link-45deg")
	test("upload",          "cloud-arrow-up")
	test("user",            "person-circle")
	test("users",           "people")
	test("article",         "card-text")
	test("forward",         "forward")
	test("inbox",           "inbox")
	test("message",         "chat-left-text")
	test("picture",         "image")
	test("pictures",        "images")
	test("shopping-cart",   "cart3")
}

// TestProvider_Get_Passthrough verifies that a name with no translation is used as a Bootstrap name
func TestProvider_Get_Passthrough(t *testing.T) {

	provider := Provider{}

	require.Equal(t, `<i class="bi bi-star"></i>`, provider.Get("star"))
	require.Equal(t, `<i class="bi bi-plus"></i>`, provider.Get("plus"))
	require.Equal(t, `<i class="bi bi-"></i>`, provider.Get(""))

	// Names are matched exactly, so a different case is not translated
	require.Equal(t, `<i class="bi bi-ADD"></i>`, provider.Get("ADD"))
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

// FuzzProvider_Get verifies that any name renders as one <i> element, and that Write matches Get
func FuzzProvider_Get(f *testing.F) {

	f.Add("add")
	f.Add("star")
	f.Add("")
	f.Add(`"><script>`)
	f.Add("\xff\xfe")

	f.Fuzz(func(t *testing.T, name string) {

		provider := Provider{}
		markup := provider.Get(name)

		require.True(t, strings.HasPrefix(markup, `<i class="bi bi-`), "markup %q", markup)
		require.True(t, strings.HasSuffix(markup, `"></i>`), "markup %q", markup)

		var builder strings.Builder
		provider.Write(name, &builder)
		require.Equal(t, markup, builder.String())
	})
}
