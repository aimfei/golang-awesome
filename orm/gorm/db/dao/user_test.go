package dao

import (
	"context"
	"golang-awesome/orm/gorm/db/model"
	"testing"
)

func TestCreateWallet(t *testing.T) {
	type args struct {
		ctx  context.Context
		user *model.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateWallet(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("CreateWallet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRecharge(t *testing.T) {
	type args struct {
		ctx       context.Context
		userId    uint32
		available uint64
		frozen    uint64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Recharge(tt.args.ctx, tt.args.userId, tt.args.available, tt.args.frozen); (err != nil) != tt.wantErr {
				t.Errorf("Recharge() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
