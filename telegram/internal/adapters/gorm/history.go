package mygorm

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type HistoryRepository struct {
	db *gorm.DB
}

func NewHistoryRepository(db *gorm.DB) *HistoryRepository {
	return &HistoryRepository{db}
}

func (r *HistoryRepository) GetPrayerMessageID(ctx context.Context, userID int) (int, error) {
	var u User
	err := r.db.WithContext(ctx).Where("telegram_id = ?", u.TelegramID).First(&u).Error
	if err != nil {
		return 0, errors.Wrap(err, "failed to get last message id")
	}
	return int(u.LastMessageID), nil
}

func (r *HistoryRepository) StorePrayerMessageID(ctx context.Context, userID int, messageID int) error {
	u := User{TelegramID: int64(userID), LastMessageID: int64(messageID)}
	err := r.db.WithContext(ctx).Model(&u).
		Where("telegram_id = ?", u.TelegramID).
		Update("last_message_id", u.LastMessageID).Error
	if err != nil {
		return errors.Wrap(err, "failed to upate last message id")
	}
	return nil
}
