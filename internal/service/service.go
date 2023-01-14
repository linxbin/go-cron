package service

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/linxbin/cron-service/internal/cron"

	"github.com/linxbin/cron-service/global"
	"github.com/linxbin/cron-service/internal/dao"
)

type Service struct {
	ctx    context.Context
	dao    *dao.Dao
	cron   *cron.Cron
	engine *gorm.DB
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx, engine: global.DBEngine}
	svc.dao = dao.New(global.DBEngine)
	return svc
}
