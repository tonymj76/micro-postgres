package handler

import (
	"context"
	"testing"

	pbUser "github.com/tonymj76/micro-postgres/proto/user"
)

func TestService_AddUser(t *testing.T) {
	type args struct {
		ctx context.Context
		req *pbUser.AddUserRequest
		res *pbUser.User
	}
	tests := []struct {
		name    string
		s       *Service
		args    args
		wantErr bool
	}{ // NOt sure of this test
		name:    "add User",
		s:       &Service,
		args:    args{ctx: context.Background(), req: &pbUser.AddUserRequest{}, res: &pbUser.User{}},
		wantErr: true,
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.AddUser(tt.args.ctx, tt.args.req, tt.args.res); (err != nil) != tt.wantErr {
				t.Errorf("Service.AddUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
