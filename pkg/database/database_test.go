package database

import "testing"

func TestConnectionStringToDSN(t *testing.T) {
	type args struct {
		connectionString string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   string
		wantErr bool
	}{
		{
			name: "Test with valid connection string and password",
			args: args{
				connectionString: "connectionname=test database=testdb user=testuser password=testpassword",
			},
			want:    "user=testuser password=testpassword database=testdb",
			want1:   "test",
			wantErr: false,
		},
		{
			name: "Test with valid connection string without password",
			args: args{
				connectionString: "connectionname=test database=testdb user=testuser",
			},
			want:    "user=testuser database=testdb",
			want1:   "test",
			wantErr: false,
		},
		{
			name: "Test with missing connectionname",
			args: args{
				connectionString: "database=testdb user=testuser password=testpassword",
			},
			want:    "",
			want1:   "",
			wantErr: true,
		},
		{
			name: "Test with missing database",
			args: args{
				connectionString: "connectionname=test user=testuser password=testpassword",
			},
			want:    "",
			want1:   "",
			wantErr: true,
		},
		{
			name: "Test with missing user",
			args: args{
				connectionString: "connectionname=test database=testdb password=testpassword",
			},
			want:    "",
			want1:   "",
			wantErr: true,
		},
		{
			name: "Test with empty connection string",
			args: args{
				connectionString: "",
			},
			want:    "",
			want1:   "",
			wantErr: true,
		},
		{
			name: "Test with = in the password",
			args: args{
				connectionString: "connectionname=test database=testdb user=testuser password=test=password",
			},
			want:    "user=testuser password=test=password database=testdb",
			want1:   "test",
			wantErr: false,
		},
		{
			name: "Test with invalid connection string",
			args: args{
				connectionString: "connectionname=test",
			},
			want:    "",
			want1:   "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := ConnectionStringToDSN(tt.args.connectionString)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConnectionStringToDSN() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConnectionStringToDSN() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ConnectionStringToDSN() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
