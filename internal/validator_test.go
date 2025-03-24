package internal

import "testing"

func TestTitleCaseKey(t *testing.T) {
	for input, want := range map[string]string{
		"foo-bar": "Foo-Bar",
		"cOnTeNt-TyPE": "Content-Type",
		"host": "Host",
		"host-": "Host-",
		"ha22-o3st": "Ha22-O3st",
	} {
		if got := AsTitle(input); got != want {
			t.Errorf("TitleCaseKey(%q) = %q, want %q", input, got, want)
		}
	}
}
