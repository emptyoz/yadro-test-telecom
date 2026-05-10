package config

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
	"yadro-intern-test/internal/domain"
	"yadro-intern-test/internal/helpers"
)

type rawConfig struct {
	Floors   int    `json:"Floors"`
	Monsters int    `json:"Monsters"`
	OpenAt   string `json:"OpenAt"`
	Duration int    `json:"Duration"`
}

func Load(path string) (domain.Config, error) {
	var raw rawConfig

	data, err := os.ReadFile(path)
	if err != nil {
		return domain.Config{}, fmt.Errorf("read config %q: %w", path, err)
	}

	if err := json.Unmarshal(data, &raw); err != nil {
		return domain.Config{}, fmt.Errorf("decode config %q: %w", path, err)
	}

	openAt, err := helpers.TimeIntoDuration(raw.OpenAt)
	if err != nil {
		return domain.Config{}, fmt.Errorf("parse OpenAt %q: %w", raw.OpenAt, err)
	}

	return domain.Config{
		Floors:   raw.Floors,
		Monsters: raw.Monsters,
		OpenAt:   openAt,
		Duration: time.Duration(raw.Duration) * time.Hour,
	}, nil
}
