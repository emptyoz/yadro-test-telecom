package engine

import (
	"testing"
	"time"

	"yadro-intern-test/internal/domain"
)

func TestBuildFinalReport_Gold(t *testing.T) {
	stats := &domain.GameStats{
		Config: domain.Config{Floors: 2, Monsters: 2},
		Players: map[int]*domain.PlayerState{
			1: {
				Player:           domain.Player{ID: 1, Health: 35},
				Registered:       true,
				BossKilled:       true,
				FloorsCleared:    1, // Floors-1
				StartTime:        14*time.Hour + 40*time.Minute,
				EndTime:          15*time.Hour + 4*time.Minute,
				FloorsClearTotal: 5 * time.Minute,
				BossKillDuration: 11 * time.Minute,
			},
			2: {
				Player:     domain.Player{ID: 2, Health: 0},
				Registered: true,
				Dead:       true,
				Finished:   true,
				StartTime:  14*time.Hour + 10*time.Minute,
				EndTime:    14*time.Hour + 29*time.Minute,
			},
			3: {
				Player:       domain.Player{ID: 3, Health: 100},
				Disqualified: true,
				Finished:     true,
			},
		},
	}

	rows := New(stats).BuildFinalReport()
	if len(rows) != 3 {
		t.Fatalf("want 3 rows, got %d", len(rows))
	}

	if rows[0].State != "SUCCESS" || rows[0].TimeSpent != 24*time.Minute || rows[0].AvgFloor != 5*time.Minute || rows[0].BossTime != 11*time.Minute {
		t.Fatalf("player1 mismatch: %+v", rows[0])
	}
	if rows[1].State != "FAIL" || rows[1].TimeSpent != 19*time.Minute {
		t.Fatalf("player2 mismatch: %+v", rows[1])
	}
	if rows[2].State != "DISQUAL" {
		t.Fatalf("player3 mismatch: %+v", rows[2])
	}
}
