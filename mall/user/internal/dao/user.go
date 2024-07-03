package dao

import (
	"context"
	"fmt"
	"user/database"
	"user/internal/model"
)

type userDao struct {
	*database.DBConn
}

func NewUerDao(conn *database.DBConn) *userDao {
	return &userDao{
		conn,
	}
}

func (d *userDao) Save(ctx context.Context, user *model.User) error {
	sql := fmt.Sprintf("insert into %s (name,gender) values(?,?)", user.TableName())
	result, err := d.Conn.ExecCtx(ctx, sql, user.Name, user.Gender)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.Id = id
	return nil
}
