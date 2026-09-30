package parseport

import "testing"

func FuzzParsePort(f *testing.F) {
	for _, seed := range []string{"", "1", "80", "65535", "65536", "-1", "hello", " 443 "} {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, input string) {
		port, err := ParsePort(input)
		if err == nil && (port < 1 || port > 65535) {
			t.Fatalf("accepted out-of-range port %d from %q", port, input)
		}
	})
}
