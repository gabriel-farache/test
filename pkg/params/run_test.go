package params

import "testing"

func TestStringToBool(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "Test with 'true' in lowercase",
			s:    "true",
			want: true,
		},
		{
			name: "Test with 'TRUE' in uppercase",
			s:    "TRUE",
			want: true,
		},
		{
			name: "Test with 'True' in mixed case",
			s:    "True",
			want: true,
		},
		{
			name: "Test with 'yes' in lowercase",
			s:    "yes",
			want: true,
		},
		{
			name: "Test with 'YES' in uppercase",
			s:    "YES",
			want: true,
		},
		{
			name: "Test with 'Yes' in mixed case",
			s:    "Yes",
			want: true,
		},
		{
			name: "Test with '1'",
			s:    "1",
			want: true,
		},
		{
			name: "Test with 'false' in lowercase",
			s:    "false",
			want: false,
		},
		{
			name: "Test with 'no' in lowercase",
			s:    "no",
			want: false,
		},
		{
			name: "Test with '0'",
			s:    "0",
			want: false,
		},
		{
			name: "Test with empty string",
			s:    "",
			want: false,
		},
		{
			name: "Test with random string",
			s:    "random",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToBool(tt.s); got != tt.want {
				t.Errorf("StringToBool() = %v, want %v", got, tt.want)
			}
		})
	}
}
