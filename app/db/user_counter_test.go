package db

import (
	"reflect"
	"testing"
)

func TestGetUserCounter(t *testing.T) {
	PrepareFixture("testdata/user_counter")
	type args struct {
		userID int64
	}
	tests := []struct {
		name    string
		args    args
		want    *UserCounter
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{1},
			want: &UserCounter{
				UserID:           1,
				ReservationCount: 3,
			},
		},
		{
			name:    "異常系 存在しないuserID",
			args:    args{100},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUserCounter(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserCounter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserCounter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecrementUserReservationCount(t *testing.T) {
	PrepareFixture("testdata/user_counter")

	type args struct {
		userID int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{1},
		},
		{
			name: "正常系 IDが存在しない",
			args: args{100},
		},
		{
			name:    "異常系 マイナスになる",
			args:    args{2},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DecrementUserReservationCount(tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("DecrementUserReservationCount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIncrementUserReservationCount(t *testing.T) {
	PrepareFixture("testdata/user_counter")

	type args struct {
		userID int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{1},
		},
		{
			name: "正常系 カウンタが存在しない(INSERT)",
			args: args{3},
		},
		{
			name: "異常系 外部キー制約を満たさない",
			args: args{3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := IncrementUserReservationCount(tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("IncrementUserReservationCount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
