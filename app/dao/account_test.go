package dao

import (
	"context"
	"testing"
	"yatter-backend-go/app/domain/object"
)

func TestCreate(t *testing.T) {
	tests := []struct {
		name    string
		ctx     context.Context
		account *object.Account
		wantErr bool
	}{
		{
			name: "example",
			ctx:  context.Background(),
			account: &object.Account{
				Username:     "Username",
				PasswordHash: "PasswordHash",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewTestDao(t).Account()
			if err := r.Create(tt.ctx, tt.account); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
			a, err := r.FindByUsername(tt.ctx, tt.account.Username)
			if err != nil {
				t.Fatal(err)
			}
			if a.PasswordHash != tt.account.PasswordHash {
				t.Errorf("unexpected PasswordHash got = %s, want %s", a.PasswordHash, tt.account.PasswordHash)
			}
		})
	}
}
