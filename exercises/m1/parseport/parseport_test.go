package parseport

import "testing"

func TestParsePort(t *testing.T) {
	tests := []struct {
		in      string
		want    int
		wantErr bool
	}{
		{"1", 1, false},
		{"65535", 65535, false},
		{" 8080 ", 8080, false},
		{"0", 0, true},
		{"65536", 0, true},
		{"-1", 0, true},
		{"hello", 0, true},
		{"", 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got, err := ParsePort(tt.in)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ParsePort(%q) error = %v", tt.in, err)
			}
			if !tt.wantErr && got != tt.want {
				t.Fatalf("ParsePort(%q) = %d, want %d", tt.in, got, tt.want)
			}
		})
	}
}
