package service

/*
	在应用分层中, service 层主要是针对业务逻辑的封装
*/

import (
	"blog/global"
	"blog/internal/dao"
	"context"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(global.DBEngine)
	return svc
}
