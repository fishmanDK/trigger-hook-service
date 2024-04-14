package service

import (
	"context"
	"fmt"
	"log/slog"
	"time"
)

type Service struct {
	log      *slog.Logger
	deleter  Deleter
	tokenTTL time.Duration
}

type Deleter interface {
	ScheduleFullDeletion(ctx context.Context, bannerID int64) error
	ScheduleDeletion(ctx context.Context, bannerID, tagID, FeatureID int64) error
}

func NewService(
	log *slog.Logger,
	deleter Deleter,
	tokenTTL time.Duration,
) *Service {
	return &Service{
		deleter:  deleter,
		log:      log,
		tokenTTL: tokenTTL,
	}
}

//func (a *Service) ScheduleFullDeletion(ctx context.Context, bannerID int64) error {
//	const op = "service.ScheduleFullDeletion"
//
//	log := a.log.With(
//		slog.String("op", op),
//	)
//
//	err := a.deleter.ScheduleDeletion(ctx, bannerID, tagID, featureID)
//	if err != nil {
//		log.Error("failed to make a new application", err)
//		return fmt.Errorf("%s: %v", op, err)
//	}
//
//	return nil
//}

func (a *Service) ScheduleDeletion(ctx context.Context, bannerID, tagID, featureID int64) error {
	const op = "service.ScheduleDeletion"

	log := a.log.With(
		slog.String("op", op),
	)

	err := a.deleter.ScheduleDeletion(ctx, bannerID, tagID, featureID)
	if err != nil {
		log.Error("failed to make a new application", err)
		return fmt.Errorf("%s: %v", op, err)
	}

	return nil
}
