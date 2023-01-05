package prayer

import (
	"encoding/json"
	"fmt"
	"time"
)

type PrayerTimes struct {
	Date    string    `json:"date"`
	Fajr    time.Time `json:"fajr"`
	Sunrise time.Time `json:"sunrise"`
	Dhuhr   time.Time `json:"dhuhr"`
	Asr     time.Time `json:"asr"`
	Maghrib time.Time `json:"maghrib"`
	Isha    time.Time `json:"isha"`
}

func NewPrayerTimes(date string, fajr, sunrise, dhuhr, asr, maghrib, isha time.Time) PrayerTimes {
	return PrayerTimes{
		Date:    date,
		Fajr:    fajr,
		Sunrise: sunrise,
		Dhuhr:   dhuhr,
		Asr:     asr,
		Maghrib: maghrib,
		Isha:    isha,
	}
}

func (p PrayerTimes) MarshalBinary() ([]byte, error) {
	return json.Marshal(p)
}

func (p PrayerTimes) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &p)
}

func (p *PrayerTimes) EnString() string {
	return fmt.Sprintf(
		`
		Date    : %s
		Fajr    : %s
		Sunrise : %s
		Dhuhr   : %s
		Asr     : %s
		Maghrib : %s
		Isha    : %s
		`, p.Date, p.Fajr, p.Sunrise, p.Dhuhr, p.Asr, p.Maghrib, p.Isha,
	)
}

// HTML returns prayer times in HTML format
func (p *PrayerTimes) EnHTML() string {
	return fmt.Sprintf(
		`
		<b>Date</b>    : %s
		<b>Fajr</b>    : %s
		<b>Sunrise</b> : %s
		<b>Dhuhr</b>   : %s
		<b>Asr</b>     : %s
		<b>Maghrib</b> : %s
		<b>Isha</b>    : %s
		`, p.Date, p.Fajr, p.Sunrise, p.Dhuhr, p.Asr, p.Maghrib, p.Isha,
	)
}

func (p *PrayerTimes) ArString() string {
	return fmt.Sprintf(
		`
		التاريخ : %s
		الفجر  : %s
		الشروق : %s
		الظهر  : %s
		العصر  : %s
		المغرب : %s
		العشاء : %s
		`, p.Date, p.Fajr, p.Sunrise, p.Dhuhr, p.Asr, p.Maghrib, p.Isha,
	)
}

func (p *PrayerTimes) ArHTML() string {
	return fmt.Sprintf(
		`
		<b>التاريخ</b> : %s
		<b>الفجر</b>  : %s
		<b>الشروق</b> : %s
		<b>الظهر </b> : %s
		<b>العصر</b>  : %s
		<b>المغرب</b> : %s
		<b>العشاء</b> : %s
		
		`, p.Date, p.Fajr.Format("HH:MM"), p.Sunrise, p.Dhuhr, p.Asr, p.Maghrib, p.Isha,
	)
}

// Russian
func (p *PrayerTimes) RuString() string {
	return fmt.Sprintf(
		`
		Дата    : %s
		Фаджр  : %s
		Восход : %s
		Зухр   : %s
		Аср    : %s
		Магриб : %s
		Иша    : %s
		`, p.Date, p.Fajr, p.Sunrise, p.Dhuhr, p.Asr, p.Maghrib, p.Isha,
	)
}

func (p *PrayerTimes) RuHTML() string {
	return fmt.Sprintf(
		`
		%s
		<b>Фаджр</b>  : %s
		<b>Восход</b> : %s
		<b>Зухр </b>  : %s
		<b>Аср</b>    : %s
		<b>Магриб</b> : %s
		<b>Иша</b>    : %s
		`, p.Date, p.Fajr, p.Sunrise, p.Dhuhr, p.Asr, p.Maghrib, p.Isha,
	)
}
