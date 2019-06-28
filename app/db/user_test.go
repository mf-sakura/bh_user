package db

import (
	"reflect"
	"testing"
)

func TestInsertUser(t *testing.T) {
	type args struct {
		user User
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "正常系(クエリチェック)",
			args: args{
				User{Name: "hoge"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := InsertUser(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("InsertUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGetUser(t *testing.T) {
	PrepareFixture("testdata/user")
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{1},
			want: &User{
				ID:   1,
				Name: "foo",
			},
		},
		{
			name:    "存在しないID",
			args:    args{100},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUser(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
