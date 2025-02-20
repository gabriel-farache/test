package cli

import (
	"testing"
)

func TestIOStreams_SetColorEnabled(t *testing.T) {
	tests := []struct {
		name                         string
		colorEnabled                 bool
		wantColorEnabled             bool
		wantProgressIndicatorEnabled bool
	}{
		{
			name:                         "Test with colorEnabled set to true",
			colorEnabled:                 true,
			wantColorEnabled:             true,
			wantProgressIndicatorEnabled: true,
		},
		{
			name:                         "Test with colorEnabled set to false",
			colorEnabled:                 false,
			wantColorEnabled:             false,
			wantProgressIndicatorEnabled: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			streams := &IOStreams{}
			streams.SetColorEnabled(tt.colorEnabled)
			if streams.colorEnabled != tt.wantColorEnabled {
				t.Errorf("SetColorEnabled() colorEnabled = %v, want %v", streams.colorEnabled, tt.wantColorEnabled)
			}
			if streams.progressIndicatorEnabled != tt.wantProgressIndicatorEnabled {
				t.Errorf("SetColorEnabled() progressIndicatorEnabled = %v, want %v", streams.progressIndicatorEnabled, tt.wantProgressIndicatorEnabled)
			}
		})
	}
}
