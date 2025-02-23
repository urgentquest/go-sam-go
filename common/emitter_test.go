package common

import (
	"bytes"
	"testing"
)

var (
	from = RandPort()
	to   = RandPort()
)

func setupTestEmit() *SAMEmit {
	return &SAMEmit{
		I2PConfig: I2PConfig{
			SamHost:  "127.0.0.1",
			SamPort:  7656,
			SamMin:   "3.0",
			SamMax:   "3.1",
			Style:    "STREAM",
			TunName:  "testid",
			Fromport: from,
			Toport:   to,
			SigType:  "ED25519",
		},
	}
}

func TestSAMEmit_Hello(t *testing.T) {
	emit := setupTestEmit()
	want := "HELLO VERSION MIN=3.0 MAX=3.1 \n"

	t.Run("string output", func(t *testing.T) {
		if got := emit.Hello(); got != want {
			t.Errorf("Hello() = %v, want %v", got, want)
		}
	})

	t.Run("byte output", func(t *testing.T) {
		if got := emit.HelloBytes(); !bytes.Equal(got, []byte(want)) {
			t.Errorf("HelloBytes() = %v, want %v", got, []byte(want))
		}
	})
}

func TestSAMEmit_GenerateDestination(t *testing.T) {
	emit := setupTestEmit()
	want := "DEST GENERATE ED25519 \n"

	t.Run("string output", func(t *testing.T) {
		if got := emit.GenerateDestination(); got != want {
			t.Errorf("GenerateDestination() = %v, want %v", got, want)
		}
	})

	t.Run("byte output", func(t *testing.T) {
		if got := emit.GenerateDestinationBytes(); !bytes.Equal(got, []byte(want)) {
			t.Errorf("GenerateDestinationBytes() = %v, want %v", got, []byte(want))
		}
	})
}

func TestSAMEmit_Lookup(t *testing.T) {
	tests := []struct {
		name   string
		lookup string
		want   string
	}{
		{"basic lookup", "test.i2p", "NAMING LOOKUP NAME=test.i2p \n"},
		{"empty lookup", "", "NAMING LOOKUP NAME= \n"},
	}

	emit := setupTestEmit()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := emit.Lookup(tt.lookup); got != tt.want {
				t.Errorf("Lookup() = %v, want %v", got, tt.want)
			}
			if got := emit.LookupBytes(tt.lookup); !bytes.Equal(got, []byte(tt.want)) {
				t.Errorf("LookupBytes() = %v, want %v", got, []byte(tt.want))
			}
		})
	}
}

func TestSAMEmit_Connect(t *testing.T) {
	tests := []struct {
		name string
		dest string
		want string
	}{
		{
			"basic connect",
			"destination123",
			"STREAM CONNECT ID=testid " + from + " " + to + " DESTINATION=destination123 \n",
		},
		{
			"empty destination",
			"",
			"STREAM CONNECT ID=testid " + from + " " + to + " DESTINATION= \n",
		},
	}

	emit := setupTestEmit()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := emit.Connect(tt.dest); got != tt.want {
				t.Errorf("Connect() = %v, want %v", got, tt.want)
			}
			if got := emit.ConnectBytes(tt.dest); !bytes.Equal(got, []byte(tt.want)) {
				t.Errorf("ConnectBytes() = %v, want %v", got, []byte(tt.want))
			}
		})
	}
}

func TestSAMEmit_Accept(t *testing.T) {
	emit := setupTestEmit()
	want := "STREAM ACCEPT ID=testid " + from + " " + to + ""

	t.Run("string output", func(t *testing.T) {
		if got := emit.Accept(); got != want {
			t.Errorf("Accept() = %v, want %v", got, want)
		}
	})

	t.Run("byte output", func(t *testing.T) {
		if got := emit.AcceptBytes(); !bytes.Equal(got, []byte(want)) {
			t.Errorf("AcceptBytes() = %v, want %v", got, []byte(want))
		}
	})
}
