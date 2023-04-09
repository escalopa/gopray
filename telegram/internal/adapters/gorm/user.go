package mygorm

import (
	"context"

	"github.com/escalopa/gopray/pkg/core"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) StoreUser(ctx context.Context, params core.CreateUserParams) error {
	// Create user
	u := User{
		TelegramID: int64(params.TelegramID),
		FirstName:  params.FirstName,
		LastName:   params.LastName,
		Username:   params.Username,
		LangCode:   params.LangCode,
	}
	err := r.db.WithContext(ctx).Create(&u).Error
	if err != nil {
		return errors.Wrap(err, "failed to store user")
	}
	return nil
}

func (r *UserRepository) GetUser(ctx context.Context, id int) (core.User, error) {
	u := User{TelegramID: int64(id)}
	err := r.db.WithContext(ctx).Where("telegram_id = ?", id).First(&u).Error
	if err != nil {
		return core.User{}, errors.Wrap(err, "failed to get user")
	}
	return fromUser(u), nil
}

func fromUser(u User) core.User {
	return core.User{
		TelegramID: u.TelegramID,
		IsAdmin:    u.IsAdmin,
	}
}
