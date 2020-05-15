package handler

import (
	"context"

	"github.com/micro/go-micro/v2/errors"
	"github.com/sirupsen/logrus"
	"github.com/tonymj76/micro-postgres/datastore"
	pbUser "github.com/tonymj76/micro-postgres/proto/user"
)

//Service _
type Service struct {
	R datastore.UserRepo
}

//AddUser _
func (s *Service) AddUser(ctx context.Context, req *pbUser.AddUserRequest, res *pbUser.User) error {
	db, err := s.R.Create(ctx, req)
	if err != nil {
		return errors.InternalServerError("501", "service.user.AddUser", err.Error())
	}
	res.CreatedAt = db.CreatedAt
	res.Id = db.Id
	res.Role = db.Role
	logrus.Info("response feed back ", res)
	return nil
}

//DeleteUser _
func (s *Service) DeleteUser(ctx context.Context, req *pbUser.DeleteUserRequest, res *pbUser.User) error {
	db, err := s.R.Delete(ctx, req)
	if err != nil {
		return errors.InternalServerError("501", "service.user.DeleteUser", err.Error())
	}
	res.Id = db.Id
	res.Role = db.Role
	return nil
}

//ListUsers _
func (s *Service) ListUsers(ctx context.Context, req *pbUser.ListUsersRequest, stream pbUser.UserService_ListUsersStream) error {
	err := s.R.List(req, stream)
	if err != nil {
		return errors.InternalServerError("501", "service.user.ListUsers", err.Error())
	}
	return nil
}
