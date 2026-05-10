package helpers

import (
	"fmt"
	"strings"
	"time"
)

// Форматирование времени
func FormatTime(d time.Duration) string {
	h := int(d.Hours())
	m := int(d.Minutes()) % 60
	s := int(d.Seconds()) % 60

	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}

// Перевод строки типа: [14:00:00]
// В формат time.Duration
func TimeIntoDuration(date string) (time.Duration, error) {
	date = strings.Trim(date, "[]")

	t, err := time.Parse("15:04:05", date)
	if err != nil {
		return 0, fmt.Errorf("time must match HH:MM:SS, got %q: %w", date, err)
	}

	dur := time.Duration(t.Hour())*time.Hour +
		time.Duration(t.Minute())*time.Minute +
		time.Duration(t.Second())*time.Second

	return dur, nil
}
