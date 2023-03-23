package memory

import (
	"context"
	"reflect"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/escalopa/gopray/pkg/language"
	"github.com/stretchr/testify/require"
)

func TestScripts(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	src := NewScriptRepository()

	tests := []struct {
		name   string
		lang   string
		script *language.Script
	}{
		{
			name:   "en",
			lang:   "en",
			script: randomScript(),
		},
		{
			name:   "es",
			lang:   "es",
			script: randomScript(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := src.StoreScript(ctx, tt.lang, tt.script)
			require.NoError(t, err)
			script, err := src.GetScript(ctx, tt.lang)
			require.NoError(t, err)
			require.True(t, reflect.DeepEqual(tt.script, script))
		})
	}

	cancel()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			script, err := src.GetScript(ctx, tt.lang)
			require.Error(t, err)
			require.Nil(t, script)
		})
	}
}

func randomScript() *language.Script {
	return &language.Script{
		DataPickerStart: gofakeit.InputName(),

		January:   gofakeit.InputName(),
		February:  gofakeit.InputName(),
		March:     gofakeit.InputName(),
		April:     gofakeit.InputName(),
		May:       gofakeit.InputName(),
		June:      gofakeit.InputName(),
		July:      gofakeit.InputName(),
		August:    gofakeit.InputName(),
		September: gofakeit.InputName(),
		October:   gofakeit.InputName(),
		November:  gofakeit.InputName(),
		December:  gofakeit.InputName(),

		LanguageSelectionStart:   gofakeit.InputName(),
		LanguageSelectionSuccess: gofakeit.InputName(),
		LanguageSelectionFail:    gofakeit.InputName(),

		Fajr:    gofakeit.InputName(),
		Sunrise: gofakeit.InputName(),
		Dhuhr:   gofakeit.InputName(),
		Asr:     gofakeit.InputName(),
		Maghrib: gofakeit.InputName(),
		Isha:    gofakeit.InputName(),

		PrayrifyTableDay:    gofakeit.InputName(),
		PrayrifyTablePrayer: gofakeit.InputName(),
		PrayrifyTableTime:   gofakeit.InputName(),
		PrayerFail:          gofakeit.InputName(),

		SubscriptionSuccess: gofakeit.InputName(),
		SubscriptionError:   gofakeit.InputName(),

		UnsubscriptionSuccess: gofakeit.InputName(),
		UnsubscriptionError:   gofakeit.InputName(),

		PrayerSoon:    gofakeit.InputName(),
		PrayerArrived: gofakeit.InputName(),
		GomaaDay:      gofakeit.InputName(),

		Help: gofakeit.InputName(),

		FeedbackStart:   gofakeit.InputName(),
		FeedbackSuccess: gofakeit.InputName(),
		FeedbackFail:    gofakeit.InputName(),

		BugReportStart:   gofakeit.InputName(),
		BugReportSuccess: gofakeit.InputName(),
		BugReportFail:    gofakeit.InputName(),
	}
}
