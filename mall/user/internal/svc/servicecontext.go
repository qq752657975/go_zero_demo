package svc

import (
	"user/database"
	"user/internal/config"
	"user/internal/dao"
	"user/internal/repo"
)

// ServiceContext logic 就是普通的service层，依赖的资源池
type ServiceContext struct {
	Config   config.Config
	UserRepo repo.UserRepo
}

func NewServiceContext(c config.Config) *ServiceContext {
	connect := database.Connect(c.Mysql.DataSource)
	return &ServiceContext{
		Config:   c,
		UserRepo: dao.NewUerDao(connect),
	}
}
