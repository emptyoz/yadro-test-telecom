package engine

import (
	"sort"
	"time"
	"yadro-intern-test/internal/domain"
)

func (e *Engine) BuildFinalReport() []domain.FinalRow {
	ids := make([]int, 0, len(e.stats.Players))
	for id := range e.stats.Players {
		ids = append(ids, id)
	}
	sort.Ints(ids)

	rows := make([]domain.FinalRow, 0, len(ids))
	for _, id := range ids {
		ps := e.stats.Players[id]

		state := "FAIL"
		if !ps.Registered || ps.Disqualified {
			state = "DISQUAL"
		} else if ps.BossKilled && ps.FloorsCleared == e.stats.Config.Floors-1 {
			state = "SUCCESS"
		}

		var timeSpent, avgFloor, bossTime time.Duration
		if ps.StartTime > 0 || ps.EndTime > 0 {
			timeSpent = ps.EndTime - ps.StartTime
		}
		if ps.FloorsCleared > 0 {
			avgFloor = ps.FloorsClearTotal / time.Duration(ps.FloorsCleared)
		}
		if ps.BossKilled {
			bossTime = ps.BossKillDuration
		}

		rows = append(rows, domain.FinalRow{
			State: state, PlayerID: id,
			TimeSpent: timeSpent, AvgFloor: avgFloor, BossTime: bossTime, HP: ps.Health,
		})
	}

	return rows
}
