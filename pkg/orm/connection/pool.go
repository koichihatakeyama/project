package connection

import (
	"database/sql"
	"project/pkg/orm/config"
	"sync"
)

type ConnectionPool struct {
	db        *sql.DB
	config    *config.Config
	stmtCache *sync.Map
}

func NewConnectionPool(cfg *config.Config) (*ConnectionPool, error) {
	db, err := sql.Open(cfg.DriverName, cfg.DSN)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)

	return &ConnectionPool{
		db:        db,
		config:    cfg,
		stmtCache: &sync.Map{},
	}, nil
}

func (p *ConnectionPool) GetDB() *sql.DB {
	return p.db
}
