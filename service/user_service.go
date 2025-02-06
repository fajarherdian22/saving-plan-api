package service

import (
	"context"
	"database/sql"

	"github.com/fajarherdian22/saving-plan-api/exception"
	"github.com/fajarherdian22/saving-plan-api/repository"
	"github.com/go-sql-driver/mysql"
)

type UserServiceImpl struct {
	q *repository.Queries
}

func NewUserService(q *repository.Queries) *UserServiceImpl {
	return &UserServiceImpl{
		q: q,
	}
}

func (service *UserServiceImpl) CreateUser(ctx context.Context, arg repository.CreateUserParams) (repository.User, error) {
	var resp repository.User
	err := service.q.CreateUser(ctx, arg)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			switch mysqlErr.Number {
			case 1062:
				return resp, exception.NewForbiddenError("Duplicate entry error")
			case 1451:
				return resp, exception.NewBadRequestError(mysqlErr.Message)
			}
		}
		return resp, exception.NewNotFoundError(err.Error())
	}

	payload, err := service.q.GetUser(ctx, arg.Email)
	if err != nil {
		return payload, exception.NewNotFoundError(err.Error())
	}

	return payload, nil
}

func (service *UserServiceImpl) GetUser(ctx context.Context, arg string) (repository.User, error) {
	var resp repository.User
	payload, err := service.q.GetUser(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			return resp, exception.NewBadRequestError(err.Error())
		}
		return resp, exception.NewInternalError(err.Error())
	}
	return payload, nil
}
