package sDeps

import (
	"context"

	"go.keploy.io/server/pkg/models"
	"go.uber.org/zap"
)

func NewSDepsService(c models.InfraDepsDB, log *zap.Logger) *SDeps {
	return &SDeps{
		sdb: c,
		log: log,
	}
}

type SDeps struct {
	sdb models.InfraDepsDB
	log *zap.Logger
}

func (s *SDeps) Insert(ctx context.Context, doc models.InfraDeps) error {
	if count, err := s.sdb.CountDocs(ctx, doc.AppID, doc.TestName); err == nil && count > 0 {
		return s.sdb.UpdateArr(ctx, doc.AppID, doc.TestName, doc)
	}
	return s.sdb.Insert(ctx, doc)
}

func (s *SDeps) Get(ctx context.Context, app string, testName string) ([]models.InfraDeps, error) {
	return s.sdb.Get(ctx, app, testName)
}
