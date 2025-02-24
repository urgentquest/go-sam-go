package stream

import (
	"testing"

	"github.com/go-i2p/go-sam-go/common"
	"github.com/go-i2p/i2pkeys"
)

func TestNewStreamSession_Integration(t *testing.T) {
	commonSam, err := common.NewSAM("127.0.0.1:7656")
	if err != nil {
		t.Fatalf("NewSAM() error = %v", err)
	}
	tests := []struct {
		name    string
		id      string
		options []string
		wantErr bool
	}{
		{
			name:    "basic session",
			id:      "test1",
			options: nil,
			wantErr: false,
		},
		{
			name:    "with options",
			id:      "test2",
			options: []string{"inbound.length=2", "outbound.length=2"},
			wantErr: false,
		},
		{
			name:    "invalid options",
			id:      "test3",
			options: []string{"invalid=true"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sam := &SAM{SAM: commonSam}
			keys, _ := i2pkeys.NewDestination()

			session, err := sam.NewStreamSession(tt.id, *keys, tt.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewStreamSession() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				session.Close()
			}
		})
	}
}
