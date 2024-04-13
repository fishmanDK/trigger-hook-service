package postgres

import (
	"context"
	"fmt"
	"github.com/fishmanDK/trigger_service/internal/config"
	"github.com/fishmanDK/trigger_service/internal/storage"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"time"
)

type BannerOperations interface {
	GetUserBanner() (struct{}, error)
}

type Postgres struct {
	db *sqlx.DB
}

func NewPostgres(cfg config.PostgresConfig) (*Postgres, error) {
	const op = "postgres.NewPostgres"
	db, err := sqlx.Open("postgres", cfg.String())
	if err != nil {
		return nil, fmt.Errorf("%s: %v", op, err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("%s: %v", op, err)
	}

	return &Postgres{
		db: db,
	}, nil
}

func (p *Postgres) ScheduleFullDeletion(ctx context.Context, bannerID int64) error {
	const op = "postgres.ScheduleFullDeletion"

	query := `
		INSERT INTO deletion_requests (bannerID, expires_at) VALUES 
		($1, $2);
	`

	expiresAt := time.Now().Add(storage.Delay)

	_, err := p.db.Exec(query, bannerID, expiresAt)
	if err != nil {
		return fmt.Errorf("%s: %v", op, err)
	}
	return nil
}
func (p *Postgres) SchedulePartialDeletion(ctx context.Context, tagID, featureID int64) error {
	const op = "postgres.SchedulePartialDeletion"

	query := `
		INSERT INTO deletion_requests (tagID, featureID, expires_at) VALUES 
		($1, $2, $3);
	`

	expiresAt := time.Now().Add(storage.Delay)

	_, err := p.db.Exec(query, tagID, featureID, expiresAt)
	if err != nil {
		return fmt.Errorf("%s: %v", op, err)
	}

	return nil
}
