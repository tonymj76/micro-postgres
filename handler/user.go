package handler

import (
	"context"

	"github.com/tonymj76/micro-postgres/datastore"
	pbUser "github.com/tonymj76/micro-postgres/proto/user"
)

//Service _
type Service struct {
	R datastore.UserRepo
}

//AddUser _
func (s *Service) AddUser(ctx context.Context, req *pbUser.AddUserRequest, res *pbUser.User) (err error) {
	res, err = s.R.Create(ctx, req)
	return
}

//DelectUser _
func (s *Service) DelectUser(ctx context.Context, req *pbUser.DeleteUserRequest, res *pbUser.User) (err error) {
	res, err = s.R.Delete(ctx, req)
	return
}

//ListUsers _
func (s *Service) ListUsers(ctx context.Context, req *pbUser.ListUsersRequest, stream pbUser.UserService_ListUsersStream) error {
	return s.R.List(ctx, req, stream)
}
