package datastore

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/micro/go-micro/v2/errors"
	"github.com/sirupsen/logrus"
	pbUser "github.com/tonymj76/micro-postgres/proto/user"
)

type pbUsers []*pbUser.User

//UserRepo communicates whith the database and squirrel
type UserRepo interface {
	Create(context.Context, *pbUser.AddUserRequest) (*pbUser.User, error)
	Delete(context.Context, *pbUser.DeleteUserRequest) (*pbUser.User, error)
	List(*pbUser.ListUsersRequest, pbUser.UserService_ListUsersStream) error
}

//Create _
func (c *Connection) Create(ctx context.Context, u *pbUser.AddUserRequest) (*pbUser.User, error) {
	c.Logger.Info("inserting data to db ", u)
	q := c.SB.Insert(
		"users",
	).SetMap(map[string]interface{}{
		"role": (roleWrapper)(u.GetRole()),
	}).Suffix(
		"RETURNING id, role, created_at",
	)
	c.Logger.Info("success create of user")
	return scanUser(q.QueryRowContext(ctx))
}

//Delete _
func (c *Connection) Delete(ctx context.Context, user *pbUser.DeleteUserRequest) (*pbUser.User, error) {
	dq, args, err := c.SB.Delete(
		"users",
	).Where(squirrel.Eq{"id": user.GetId()}).Suffix(
		"RETURNING id, role, created_at",
	).ToSql()
	if err != nil {
		return nil, err
	}
	return scanUser(c.DB.QueryRowContext(ctx, dq, args...))
}

//List _
func (c *Connection) List(req *pbUser.ListUsersRequest, stream pbUser.UserService_ListUsersStream) error {
	q := c.SB.Select(
		"id",
		"role",
		"created_at",
	).From(
		"users",
	).OrderBy(
		"created_at ASC",
	)
	if req.GetCreatedSince() != nil {
		q = q.Where(squirrel.Gt{
			"created_at": (*timeWrapper)(req.GetCreatedSince()),
		})
	}
	if req.GetOlderThan() != nil {
		q = q.Where(
			squirrel.Expr(
				"CURRENT_TIMESTAMP - created_at > ?", (*durationWrapper)(req.GetOlderThan()),
			),
		)
	}
	rows, err := q.QueryContext(stream.Context())
	if err != nil {
		return err
	}

	defer func() {
		cerr := rows.Close()
		if err == nil && cerr != nil {
			err = errors.InternalServerError("501", "service.user.List", err)
		}
	}()

	for rows.Next() {
		user, err := scanUser(rows)
		if err != nil {
			return errors.InternalServerError("501", "service.user.List", err)
		}
		err = stream.Send(&pbUser.User{
			Id:        user.Id,
			Role:      user.Role,
			CreatedAt: user.CreatedAt,
		})
		logrus.Info("GETTING STREAM ERRORS: ", err)
		if err != nil {
			return errors.InternalServerError("501", "service.user.List", err)
		}
		c.Logger.WithFields(logrus.Fields{
			"ID":      user.Id,
			"role":    user.GetRole(),
			"created": user.CreatedAt,
		}).Info("scaning into Stream")
	}
	return nil
}
