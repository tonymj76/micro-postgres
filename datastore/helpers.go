package datastore

import (
	"database/sql"
	"errors"

	"github.com/Masterminds/squirrel"
	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"
	"github.com/golang/protobuf/ptypes/timestamp"
	pbError "github.com/micro/go-micro/v2/errors"
	"github.com/sirupsen/logrus"
	"github.com/tonymj76/micro-postgres/datastore/migrations"
	pbUser "github.com/tonymj76/micro-postgres/proto/user"
)

// version defines the current migration version. This ensures the app
// is always compatible with the version of the database.
const version = 1

func validateSchema(db *sql.DB) error {
	sourceInstance, err := bindata.WithInstance(bindata.Resource(migrations.AssetNames(), migrations.Asset))
	if err != nil {
		return err
	}

	targetInstance, err := postgres.WithInstance(db, new(postgres.Config))
	if err != nil {
		return err
	}
	m, err := migrate.NewWithInstance("go-bindata", sourceInstance, "postgres", targetInstance)
	if err != nil {
		return err
	}
	err = m.Migrate(version) // current version
	if err != nil && err != migrate.ErrNoChange {
		return err
	}
	return sourceInstance.Close()
}

func scanUser(row squirrel.RowScanner) (*pbUser.User, error) {
	user := pbUser.User{}
	user.CreatedAt = new(timestamp.Timestamp)
	err := row.Scan(
		&user.Id,
		(*roleWrapper)(&user.Role),
		(*timeWrapper)(user.CreatedAt),
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, pbError.BadRequest("404", "no user found", err.Error())
		}
		return nil, err
	}
	logrus.WithFields(logrus.Fields{
		"ID":      user.Id,
		"role":    user.GetRole(),
		"created": user.CreatedAt,
	}).Info("scaning into response")
	return &user, nil
}
