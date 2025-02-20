package cli

import (
	"os"
	"reflect"
	"testing"
)

func TestIs256ColorSupported(t *testing.T) {
	tests := []struct {
		name      string
		term      string
		colorterm string
		want      bool
	}{
		{
			name:      "Test with TERM and COLORTERM unset",
			term:      "",
			colorterm: "",
			want:      false,
		},
		{
			name:      "Test with TERM set to '256color'",
			term:      "256color",
			colorterm: "",
			want:      true,
		},
		{
			name:      "Test with TERM set to '24bit'",
			term:      "24bit",
			colorterm: "",
			want:      true,
		},
		{
			name:      "Test with TERM set to 'truecolor'",
			term:      "truecolor",
			colorterm: "",
			want:      true,
		},
		{
			name:      "Test with COLORTERM set to '256color'",
			term:      "",
			colorterm: "256color",
			want:      true,
		},
		{
			name:      "Test with COLORTERM set to '24bit'",
			term:      "",
			colorterm: "24bit",
			want:      true,
		},
		{
			name:      "Test with COLORTERM set to 'truecolor'",
			term:      "",
			colorterm: "truecolor",
			want:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("TERM", tt.term)
			os.Setenv("COLORTERM", tt.colorterm)
			if got := Is256ColorSupported(); got != tt.want {
				t.Errorf("Is256ColorSupported() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewColorScheme(t *testing.T) {
	tests := []struct {
		name         string
		enabled      bool
		is256enabled bool
		want         *ColorScheme
	}{
		{
			name:         "Test with enabled and is256enabled set to false",
			enabled:      false,
			is256enabled: false,
			want:         &ColorScheme{enabled: false, is256enabled: false},
		},
		{
			name:         "Test with enabled set to true and is256enabled set to false",
			enabled:      true,
			is256enabled: false,
			want:         &ColorScheme{enabled: true, is256enabled: false},
		},
		{
			name:         "Test with enabled set to false and is256enabled set to true",
			enabled:      false,
			is256enabled: true,
			want:         &ColorScheme{enabled: false, is256enabled: true},
		},
		{
			name:         "Test with enabled and is256enabled set to true",
			enabled:      true,
			is256enabled: true,
			want:         &ColorScheme{enabled: true, is256enabled: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewColorScheme(tt.enabled, tt.is256enabled); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewColorScheme() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColorScheme_ColorStatus(t *testing.T) {
	type args struct {
		status string
	}
	tests := []struct {
		name string
		c    *ColorScheme
		args args
		want string
	}{
		{
			name: "Test with status 'succeeded' ",
			c: &ColorScheme{
				enabled:      false,
				is256enabled: false,
			},
			args: args{
				status: "succeeded",
			},
			want: "succeeded",
		},
		{
			name: "Test with status 'failed' ",
			c: &ColorScheme{
				enabled:      false,
				is256enabled: false,
			},
			args: args{
				status: "failed",
			},
			want: "failed",
		},
		{
			name: "Test with status 'pipelineruntimeout' ",
			c: &ColorScheme{
				enabled:      false,
				is256enabled: false,
			},
			args: args{
				status: "pipelineruntimeout",
			},
			want: "Timeout",
		},
		{
			name: "Test with status 'norun' ",
			c: &ColorScheme{
				enabled:      false,
				is256enabled: false,
			},
			args: args{
				status: "norun",
			},
			want: "norun",
		},
		{
			name: "Test with status 'running' ",
			c: &ColorScheme{
				enabled:      false,
				is256enabled: false,
			},
			args: args{
				status: "running",
			},
			want: "running",
		},
		{
			name: "Test with status 'unknown' ",
			c: &ColorScheme{
				enabled:      false,
				is256enabled: false,
			},
			args: args{
				status: "unknown",
			},
			want: "unknown",
		},
		{
			name: "Test with status 'succeeded' (enabled)",
			c: &ColorScheme{
				enabled:      true,
				is256enabled: false,
			},
			args: args{
				status: "succeeded",
			},
			want: "\x1b[0;32msucceeded\x1b[0m",
		},
		{
			name: "Test with status 'failed' (enabled)",
			c: &ColorScheme{
				enabled:      true,
				is256enabled: false,
			},
			args: args{
				status: "failed",
			},
			want: "\x1b[0;31mfailed\x1b[0m",
		},
		{
			name: "Test with status 'pipelineruntimeout' (enabled)",
			c: &ColorScheme{
				enabled:      true,
				is256enabled: false,
			},
			args: args{
				status: "pipelineruntimeout",
			},
			want: "\x1b[0;33mTimeout\x1b[0m",
		},
		{
			name: "Test with status 'norun' (enabled)",
			c: &ColorScheme{
				enabled:      true,
				is256enabled: false,
			},
			args: args{
				status: "norun",
			},
			want: "\x1b[0;38;5;246mnorun\x1b[0m",
		},
		{
			name: "Test with status 'running' (enabled)",
			c: &ColorScheme{
				enabled:      true,
				is256enabled: false,
			},
			args: args{
				status: "running",
			},
			want: "\x1b[0;34mrunning\x1b[0m",
		},
		{
			name: "Test with status 'unknown' (enabled)",
			c: &ColorScheme{
				enabled:      true,
				is256enabled: false,
			},
			args: args{
				status: "unknown",
			},
			want: "unknown",
		},
		{
			name: "Test with status 'succeeded' (enabled/is256enabled)",
			c: &ColorScheme{
				enabled:      true,
				is256enabled: true,
			},
			args: args{
				status: "succeeded",
			},
			want: "\x1b[0;32msucceeded\x1b[0m",
		},
		{
			name: "Test with status 'failed' (enabled/is256enabled)",
			c: &ColorScheme{
				enabled:      true,
				is256enabled: true,
			},
			args: args{
				status: "failed",
			},
			want: "\x1b[0;31mfailed\x1b[0m",
		},
		{
			name: "Test with status 'pipelineruntimeout' (enabled/is256enabled)",
			c: &ColorScheme{
				enabled:      true,
				is256enabled: true,
			},
			args: args{
				status: "pipelineruntimeout",
			},
			want: "\x1b[0;33mTimeout\x1b[0m",
		},
		{
			name: "Test with status 'norun' (enabled/is256enabled)",
			c: &ColorScheme{
				enabled:      true,
				is256enabled: true,
			},
			args: args{
				status: "norun",
			},
			want: "\x1b[0;38;5;246mnorun\x1b[0m",
		},
		{
			name: "Test with status 'running' (enabled/is256enabled)",
			c: &ColorScheme{
				enabled:      true,
				is256enabled: true,
			},
			args: args{
				status: "running",
			},
			want: "\x1b[0;34mrunning\x1b[0m",
		},
		{
			name: "Test with status 'unknown' (enabled/is256enabled)",
			c: &ColorScheme{
				enabled:      true,
				is256enabled: true,
			},
			args: args{
				status: "unknown",
			},
			want: "unknown",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.ColorStatus(tt.args.status); got != tt.want {
				t.Errorf("ColorScheme.ColorStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColorScheme_ColorFromString(t *testing.T) {
	tests := []struct {
		name string
		c    *ColorScheme
		s    string
		want string
	}{
		{
			name: "Test with color 'bold'",
			c:    &ColorScheme{enabled: true, is256enabled: false},
			s:    "bold",
			want: "\x1b[0;1;39mbold\x1b[0m",
		},
		{
			name: "Test with color 'red'",
			c:    &ColorScheme{enabled: true, is256enabled: false},
			s:    "red",
			want: "\x1b[0;31mred\x1b[0m",
		},
		{
			name: "Test with color 'yellow'",
			c:    &ColorScheme{enabled: true, is256enabled: false},
			s:    "yellow",
			want: "\x1b[0;33myellow\x1b[0m",
		},
		{
			name: "Test with color 'green'",
			c:    &ColorScheme{enabled: true, is256enabled: false},
			s:    "green",
			want: "\x1b[0;32mgreen\x1b[0m",
		},
		{
			name: "Test with color 'gray'",
			c:    &ColorScheme{enabled: true, is256enabled: false},
			s:    "gray",
			want: "\x1b[0;7;30mgray\x1b[0m",
		},
		{
			name: "Test with color 'magenta'",
			c:    &ColorScheme{enabled: true, is256enabled: false},
			s:    "magenta",
			want: "\x1b[0;35mmagenta\x1b[0m",
		},
		{
			name: "Test with color 'cyan'",
			c:    &ColorScheme{enabled: true, is256enabled: false},
			s:    "cyan",
			want: "\x1b[0;36mcyan\x1b[0m",
		},
		{
			name: "Test with color 'blue'",
			c:    &ColorScheme{enabled: true, is256enabled: false},
			s:    "blue",
			want: "\x1b[0;34mblue\x1b[0m",
		},
		{
			name: "Test with color 'unknown'",
			c:    &ColorScheme{enabled: true, is256enabled: false},
			s:    "unknown",
			want: "unknown",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := tt.c.ColorFromString(tt.s)
			if got := fn(tt.s); got != tt.want {
				t.Errorf("ColorFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}
