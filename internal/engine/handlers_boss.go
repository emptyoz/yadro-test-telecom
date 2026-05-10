package engine

import "yadro-intern-test/internal/domain"

func (e *Engine) handleEnterBossFloor(event domain.Event) {
	ps := e.getOrCreatePlayer(event.PlayerID)

	if !e.ensureRegistered(event, ps) {
		return
	}

	if !e.ensurePossible(event,
		!ps.Dead && ps.InDungeon && !ps.Finished &&
			ps.CurrentFloor == e.stats.Config.Floors-1) {
		return
	}

	ps.BossEntered = true
	ps.BossEnterTime = event.Time
	e.logEnterBossFloor(event)
}

func (e *Engine) handleKillBoss(event domain.Event) {
	ps := e.getOrCreatePlayer(event.PlayerID)

	if !e.ensureRegistered(event, ps) {
		return
	}

	if !e.ensurePossible(event,
		!ps.Dead && ps.InDungeon && !ps.Finished &&
			ps.BossEntered && !ps.BossKilled) {
		return
	}

	ps.BossKilled = true
	ps.BossKillDuration = event.Time - ps.BossEnterTime
	e.logKillBoss(event)
}
