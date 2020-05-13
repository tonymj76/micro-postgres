package datastore

import (
	"database/sql"
	"log"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/logrusadapter"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/micro/go-micro/v2/config"
	"github.com/sirupsen/logrus"
)

//Connection _
type Connection struct {
	Logger *logrus.Logger
	DB     *sql.DB
	SB     squirrel.StatementBuilderType
}

func loadConfig() string {
	err := config.LoadFile("config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	src := config.Map()
	conf := src["database"].(map[string]interface{})["source"].(string)
	return conf
}

//NewConnection open a connection to db
func NewConnection(logger *logrus.Logger) (*Connection, error) {
	conn, err := pgx.ParseConfig(loadConfig())
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
