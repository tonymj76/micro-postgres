package handler

import (
	"context"

	"github.com/micro/go-micro/v2/errors"
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
	if err != nil {
		return errors.InternalServerError("501", "service.user.AddUser", err.Error())
	}
	return nil
}

//DeleteUser _
func (s *Service) DeleteUser(ctx context.Context, req *pbUser.DeleteUserRequest, res *pbUser.User) (err error) {
	res, err = s.R.Delete(ctx, req)
	if err != nil {
		return errors.InternalServerError("501", "service.user.DeleteUser", err.Error())
	}
	return nil
}

//ListUsers _
func (s *Service) ListUsers(ctx context.Context, req *pbUser.ListUsersRequest, stream pbUser.UserService_ListUsersStream) error {
	err := s.R.List(ctx, req, stream)
	if err != nil {
		return errors.InternalServerError("501", "service.user.ListUsers", err.Error())
	}
	return nil
}
