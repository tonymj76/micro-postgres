package datastore

import (
	"database/sql"
	"os"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/logrusadapter"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/sirupsen/logrus"
)

//Connection _
type Connection struct {
	Logger *logrus.Logger
	DB     *sql.DB
	SB     squirrel.StatementBuilderType
}

//NewConnection open a connection to db
func NewConnection(logger *logrus.Logger) (*Connection, error) {
	conn, err := pgx.ParseConfig(os.Getenv("CONNSTR"))
	if err != nil {
		return nil, err
	}
	conn.Logger = logrusadapter.NewLogger(logger)
	connStr := stdlib.RegisterConnConfig(conn)
	db, err := sql.Open("pgx", connStr)
	err = validateSchema(db)
	if err != nil {
		return nil, err
	}

	return &Connection{
		Logger: logger,
		DB:     db,
		SB:     squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).RunWith(db),
	}, err
}

//Close connection
func (c Connection) Close() error {
	return c.DB.Close()
}
