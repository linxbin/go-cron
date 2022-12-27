package service

import (
	"context"
	"github.com/linxbin/cron-service/internal/cron"

	"github.com/linxbin/cron-service/global"
	"github.com/linxbin/cron-service/internal/dao"
)

type Service struct {
	ctx  context.Context
	dao  *dao.Dao
	cron *cron.Cron
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(global.DBEngine)
	return svc
}
