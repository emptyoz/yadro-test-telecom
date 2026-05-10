package engine

import "yadro-intern-test/internal/domain"

func (e *Engine) handleLeaveDungeon(event domain.Event) {
	ps := e.getOrCreatePlayer(event.PlayerID)

	if !e.ensureRegistered(event, ps) {
		return
	}

	if !e.ensurePossible(event,
		ps.InDungeon && !ps.Finished) {
		return
	}

	ps.InDungeon = false
	ps.Finished = true
	ps.EndTime = event.Time

	e.logLeaveDungeon(event)
}

func (e *Engine) handleCannotContinue(event domain.Event) {
	ps := e.getOrCreatePlayer(event.PlayerID)

	if !e.ensureRegistered(event, ps) {
		return
	}

	if !e.ensurePossible(event, ps.InDungeon && !ps.Finished) {
		return
	}

	ps.InDungeon = false
	ps.Finished = true
	ps.Disqualified = true
	ps.EndTime = event.Time

	e.logCannotContinue(event, event.ExtraParam)
}
