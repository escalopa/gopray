package memory

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/escalopa/gopray/telegram/internal/domain"
)

type PrayerRepository struct {
	prayers map[string]*domain.PrayerTime
	mu      sync.RWMutex
}

func NewPrayerRepository() *PrayerRepository {
	return &PrayerRepository{prayers: make(map[string]*domain.PrayerTime)}
}

func (pr *PrayerRepository) StorePrayer(_ context.Context, p *domain.PrayerTime) error {
	pr.mu.Lock()
	defer pr.mu.Unlock()

	key := formatKey(p.Day)
	pr.prayers[key] = p

	return nil
}

func (pr *PrayerRepository) GetPrayer(_ context.Context, day time.Time) (*domain.PrayerTime, error) {
	pr.mu.RLock()
	defer pr.mu.RUnlock()

	key := formatKey(day)
	val, ok := pr.prayers[key]
	if !ok || val == nil {
		return nil, domain.ErrNotFound
	}

	return val, nil
}

func formatKey(day time.Time) string {
	return fmt.Sprintf("%d/%d/%d", day.Day(), day.Month(), day.Year())
}
