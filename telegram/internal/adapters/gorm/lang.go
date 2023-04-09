package mygorm

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type LanguageRepository struct {
	db *gorm.DB
}

func NewLanguageRepository(db *gorm.DB) *LanguageRepository {
	return &LanguageRepository{db}
}

func (r *LanguageRepository) GetLang(ctx context.Context, id int) (string, error) {
	var u User
	err := r.db.WithContext(ctx).Where("telegram_id = ?", u.TelegramID).Find(&u).Error
	if err != nil {
		return "", errors.Wrap(err, "failed to get lang")
	}
	return u.LangCode, nil
}

func (r *LanguageRepository) SetLang(ctx context.Context, id int, lang string) error {
	var u User
	err := r.db.WithContext(ctx).Model(&u).
		Where("telegram_id = ?", u.TelegramID).
		Update("lang_code", lang).Error
	if err != nil {
		return errors.Wrap(err, "failed to set lang")
	}
	return nil
}
