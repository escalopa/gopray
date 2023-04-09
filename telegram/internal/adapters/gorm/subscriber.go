package mygorm

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type SubscriberRepository struct {
	db *gorm.DB
}

func NewSubscriberRepository(db *gorm.DB) *SubscriberRepository {
	return &SubscriberRepository{db}
}

func (r *SubscriberRepository) StoreSubscriber(ctx context.Context, id int) error {
	u := User{TelegramID: int64(id)}
	// Set user as subscribed
	err := r.db.WithContext(ctx).Model(&u).
		Where("telegram_id = ?", id).
		Update("is_subscribed", true).Error
	if err != nil {
		return errors.Wrapf(err, "failed to subscribe user %d", id)
	}
	return nil
}

func (r *SubscriberRepository) RemoveSubscribe(ctx context.Context, id int) error {
	u := User{TelegramID: int64(id)}
	// Set user as unsubscribed
	err := r.db.WithContext(ctx).Model(&u).
		Where("telegram_id = ?", id).
		Update("is_subscribed", false).Error
	if err != nil {
		return errors.Wrapf(err, "failed to unsubscribe user %d", id)
	}
	return nil
}

func (r *SubscriberRepository) GetSubscribers(ctx context.Context) ([]int, error) {
	var us []User
	err := r.db.WithContext(ctx).Where("is_subscribed = ?", true).Find(&us).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.Wrap(err, "failed to get subscribers")
	}
	// Convert to []int
	ids := make([]int, len(us))
	for i, s := range us {
		ids[i] = int(s.TelegramID)
	}
	return ids, nil
}
