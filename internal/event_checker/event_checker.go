package event_checker

import (
	"fmt"
	"github.com/fishmanDK/trigger_service/internal/clients/rabbitmq"
	"github.com/fishmanDK/trigger_service/internal/config"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

const (
	subjectDeleteBanner = "deleteRequest"
)

type event struct {
	id        int64 `db:"id"`
	bannerID  int64 `db:"bannerID"`
	tagID     int64 `db:"tagID"`
	featureID int64 `db:"featureID"`
}

type Checker struct {
	rabbitmq *rabbitmq.RabbitMQPublisher
	db       *sqlx.DB
}

func NewChecker(cfg config.PostgresConfig) (*Checker, error) {
	const op = "event_checker.NewChecker"
	db, err := sqlx.Open("postgres", cfg.String())
	if err != nil {
		return nil, fmt.Errorf("%s: %v", op, err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("%s: %v", op, err)
	}

	rabbitmq, err := rabbitmq.NewRabbitMQClient()
	if err != nil {
		return nil, fmt.Errorf("%s: %v", op, err)
	}

	return &Checker{
		rabbitmq: rabbitmq,
		db:       db,
	}, nil
}

func (c *Checker) Run(timeTicker time.Duration) error {
	const op = "event_checker.CheckExpiredRequests"

	ticker := time.NewTicker(timeTicker)
	defer ticker.Stop()

	selectQuery := `
			SELECT id, bannerID, tagID, featureID 
			FROM deletion_requests 
			WHERE expires_at < now() AT TIME ZONE 'UTC';;

	`

	deleteQuery := `
			DELETE FROM deletion_requests
			WHERE id = $1;
	`

	for range ticker.C {

		rows, err := c.db.Query(selectQuery)
		if err != nil {
			log.Println("Error querying database:", err)
			continue
		}
		defer rows.Close()

		for rows.Next() {
			var e event
			if err := rows.Scan(&e.id, &e.bannerID, &e.tagID, &e.featureID); err != nil {
				log.Println("Error scanning row:", err)
				continue
			}
			fmt.Println(e)
			var message rabbitmq.Message
			if e.bannerID != 0 {
				message.BannerID = e.bannerID
			} else {
				message.TagID = e.tagID
				message.FeatureID = e.featureID
			}

			err := c.rabbitmq.PublishMessage(message)
			if err != nil {
				log.Printf("Error publish message deletion_requests.id = %d: %v", e.id, err)
			}

			_, err = c.db.Exec(deleteQuery, e.id)
			if err != nil {
				log.Printf("failed delete reqest deletion_requests.id = %d: %v", e.id, err)
			}
		}
	}

	return nil
}
