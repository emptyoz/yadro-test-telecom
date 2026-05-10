package engine

import (
	"strings"
	"testing"
	"time"

	"yadro-intern-test/internal/domain"
)

func newEngineForTest() (*Engine, *domain.GameStats) {
	stats := &domain.GameStats{
		Config: domain.Config{
			Floors:   2,
			Monsters: 2,
			OpenAt:   14*time.Hour + 5*time.Minute,
			Duration: 2 * time.Hour,
		},
	}
	return New(stats), stats
}

func ev(h, m, s, pid, eid, intParam int) domain.Event {
	return domain.Event{
		Time:     time.Duration(h)*time.Hour + time.Duration(m)*time.Minute + time.Duration(s)*time.Second,
		PlayerID: pid,
		EventID:  eid,
		IntParam: intParam,
	}
}

func mustPlayer(t *testing.T, stats *domain.GameStats, id int) *domain.PlayerState {
	t.Helper()
	ps, ok := stats.Players[id]
	if !ok {
		t.Fatalf("player %d not found", id)
	}
	return ps
}

func mustLastLog(t *testing.T, stats *domain.GameStats) string {
	t.Helper()
	if len(stats.Logs) == 0 {
		t.Fatal("no logs")
	}
	return stats.Logs[len(stats.Logs)-1]
}

func TestHandlers_Gold(t *testing.T) {
	tests := []struct {
		name   string
		events []domain.Event
		check  func(*testing.T, *domain.GameStats)
	}{
		{
			name: "enter_without_register_disqual",
			events: []domain.Event{
				ev(14, 0, 0, 1, 2, 0),
			},
			check: func(t *testing.T, s *domain.GameStats) {
				ps := mustPlayer(t, s, 1)
				if !ps.Disqualified || !ps.Finished {
					t.Fatalf("want disqualified+finished")
				}
			},
		},
		{
			name: "next_floor_before_clear_impossible",
			events: []domain.Event{
				ev(14, 0, 0, 1, 1, 0),
				ev(14, 10, 0, 1, 2, 0),
				ev(14, 11, 0, 1, 4, 0),
			},
			check: func(t *testing.T, s *domain.GameStats) {
				if !strings.Contains(mustLastLog(t, s), "imposible move [4]") {
					t.Fatalf("unexpected log: %q", mustLastLog(t, s))
				}
			},
		},
		{
			name: "enter_boss_wrong_floor_impossible",
			events: []domain.Event{
				ev(14, 0, 0, 1, 1, 0),
				ev(14, 10, 0, 1, 2, 0),
				ev(14, 11, 0, 1, 6, 0),
			},
			check: func(t *testing.T, s *domain.GameStats) {
				if !strings.Contains(mustLastLog(t, s), "imposible move [6]") {
					t.Fatalf("unexpected log: %q", mustLastLog(t, s))
				}
			},
		},
		{
			name: "kill_boss_without_enter_impossible",
			events: []domain.Event{
				ev(14, 0, 0, 1, 1, 0),
				ev(14, 10, 0, 1, 2, 0),
				ev(14, 11, 0, 1, 7, 0),
			},
			check: func(t *testing.T, s *domain.GameStats) {
				if !strings.Contains(mustLastLog(t, s), "imposible move [7]") {
					t.Fatalf("unexpected log: %q", mustLastLog(t, s))
				}
			},
		},
		{
			name: "lethal_damage_dead_finished",
			events: []domain.Event{
				ev(14, 0, 0, 2, 1, 0),
				ev(14, 10, 0, 2, 2, 0),
				ev(14, 11, 0, 2, 11, 100),
			},
			check: func(t *testing.T, s *domain.GameStats) {
				ps := mustPlayer(t, s, 2)
				if !ps.Dead || !ps.Finished || ps.Health != 0 {
					t.Fatalf("want dead+finished+hp0")
				}
			},
		},
		{
			name: "heal_capped_at_100",
			events: []domain.Event{
				ev(14, 0, 0, 1, 1, 0),
				ev(14, 10, 0, 1, 2, 0),
				ev(14, 11, 0, 1, 11, 30),
				ev(14, 12, 0, 1, 10, 80),
			},
			check: func(t *testing.T, s *domain.GameStats) {
				ps := mustPlayer(t, s, 1)
				if ps.Health != 100 {
					t.Fatalf("want hp=100, got %d", ps.Health)
				}
			},
		},
		{
			name: "cannot_continue_disqual",
			events: []domain.Event{
				ev(14, 0, 0, 1, 1, 0),
				ev(14, 10, 0, 1, 2, 0),
				{Time: 14*time.Hour + 11*time.Minute, PlayerID: 1, EventID: 9, ExtraParam: "network issue"},
			},
			check: func(t *testing.T, s *domain.GameStats) {
				ps := mustPlayer(t, s, 1)
				if !ps.Disqualified || !ps.Finished || ps.InDungeon {
					t.Fatalf("unexpected state")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			eng, stats := newEngineForTest()
			eng.Process(tt.events)
			tt.check(t, stats)
		})
	}
}
