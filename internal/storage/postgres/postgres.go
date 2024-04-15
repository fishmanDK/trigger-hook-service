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
	expiresAt := time.Now().Add(storage.Delay).UTC()

	_, err := p.db.Exec(query, bannerID, expiresAt)
	if err != nil {
		return fmt.Errorf("%s: %v", op, err)
	}
	return nil
}
func (p *Postgres) ScheduleDeletion(ctx context.Context, bannerID, tagID, featureID int64) error {
	const op = "postgres.SchedulePartialDeletion"

	query := `
		INSERT INTO deletion_requests (bannerID, tagID, featureID, expires_at) VALUES 
		($1, $2, $3, $4);
	`

	expiresAt := time.Now().Add(storage.Delay).UTC()
	if bannerID == 0 {
		_, err := p.db.Exec(query, 0, tagID, featureID, expiresAt)
		if err != nil {
			return fmt.Errorf("%s: %v", op, err)
		}
	} else {
		_, err := p.db.Exec(query, bannerID, 0, 0, expiresAt)
		if err != nil {
			return fmt.Errorf("%s: %v", op, err)
		}
	}

	return nil
}
