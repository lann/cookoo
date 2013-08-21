package cli

import (
	"github.com/masterminds/cookoo"
	"testing"
	"bytes"
	"strings"
)

func TestShowHelp(t *testing.T) {
	registry, router, context := cookoo.Cookoo();

	var out bytes.Buffer

	registry.Route("test", "Testing help.").Does(ShowHelp, "didShowHelp").
		Using("show").WithDefault(true).
		Using("writer").WithDefault(&out).
		Using("summary").WithDefault("This is a summary.")


		e := router.HandleRequest("test", context, false)

		if e != nil {
			t.Error("! Unexpected error.")
		}

		res := context.Get("didShowHelp").(bool)

		if !res {
			t.Error("! Expected help to be shown.")
		}

		msg := out.String()
		// fmt.Printf(msg)
		if !strings.Contains(msg, "SUMMARY\n") {
			t.Error("! Expected 'summary' as a header.")
		}
		if !strings.Contains(msg, "This is a summary.") {
			t.Error("! Expected 'This is a summary' to be in the output. Got ", msg)
		}
}