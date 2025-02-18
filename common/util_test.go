package common

import (
	"testing"
)

func TestExtractDest_Cases(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "simple destination",
			input: "dest1234 STYLE=3",
			want:  "dest1234",
		},
		{
			name:  "empty input",
			input: "",
			want:  "",
		},
		{
			name:  "single word",
			input: "destination",
			want:  "destination",
		},
		{
			name:  "multiple spaces",
			input: "dest123   STYLE=3  KEY=value",
			want:  "dest123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExtractDest(tt.input)
			if got != tt.want {
				t.Errorf("ExtractDest(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestExtractPairString_Cases(t *testing.T) {
	tests := []struct {
		name  string
		input string
		key   string
		want  string
	}{
		{
			name:  "simple key-value",
			input: "KEY=value",
			key:   "KEY",
			want:  "value",
		},
		{
			name:  "no matching key",
			input: "OTHER=value",
			key:   "KEY",
			want:  "",
		},
		{
			name:  "multiple pairs",
			input: "FIRST=1 KEY=value LAST=3",
			key:   "KEY",
			want:  "value",
		},
		{
			name:  "empty input",
			input: "",
			key:   "KEY",
			want:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExtractPairString(tt.input, tt.key)
			if got != tt.want {
				t.Errorf("ExtractPairString(%q, %q) = %q, want %q",
					tt.input, tt.key, got, tt.want)
			}
		})
	}
}

func TestExtractPairInt_Cases(t *testing.T) {
	tests := []struct {
		name  string
		input string
		key   string
		want  int
	}{
		{
			name:  "valid integer",
			input: "NUM=123",
			key:   "NUM",
			want:  123,
		},
		{
			name:  "invalid integer",
			input: "NUM=abc",
			key:   "NUM",
			want:  0,
		},
		{
			name:  "no matching key",
			input: "OTHER=123",
			key:   "NUM",
			want:  0,
		},
		{
			name:  "empty input",
			input: "",
			key:   "NUM",
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExtractPairInt(tt.input, tt.key)
			if got != tt.want {
				t.Errorf("ExtractPairInt(%q, %q) = %d, want %d",
					tt.input, tt.key, got, tt.want)
			}
		})
	}
}

func TestSplitHostPort_Cases(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantHost string
		wantPort string
		wantErr  bool
	}{
		{
			name:     "valid hostport",
			input:    "localhost:1234",
			wantHost: "localhost",
			wantPort: "1234",
			wantErr:  false,
		},
		{
			name:     "missing port",
			input:    "localhost",
			wantHost: "localhost",
			wantPort: "0",
			wantErr:  false,
		},
		{
			name:     "empty input",
			input:    "",
			wantHost: "",
			wantPort: "0",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			host, port, err := SplitHostPort(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("SplitHostPort(%q) error = %v, wantErr %v",
					tt.input, err, tt.wantErr)
				return
			}
			if host != tt.wantHost {
				t.Errorf("SplitHostPort(%q) host = %q, want %q",
					tt.input, host, tt.wantHost)
			}
			if port != tt.wantPort {
				t.Errorf("SplitHostPort(%q) port = %q, want %q",
					tt.input, port, tt.wantPort)
			}
		})
	}
}
