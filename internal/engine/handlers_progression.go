package engine

import "yadro-intern-test/internal/domain"

func (e *Engine) handleNextFloor(event domain.Event) {
	ps := e.getOrCreatePlayer(event.PlayerID)

	if !e.ensureRegistered(event, ps) {
		return
	}

	if !e.ensurePossible(event,
		!ps.Dead && ps.InDungeon && !ps.Finished &&
			ps.CurrentFloor < e.stats.Config.Floors-1 &&
			ps.KilledOnFloor[ps.CurrentFloor] == e.stats.Config.Monsters) {
		return
	}

	// Прибавляем этаж при
	// Переходе на следующий
	ps.CurrentFloor++

	// Сохраняем время входа на этаж
	ps.FloorEnterTime = event.Time
	e.logNextFloor(event)
}

func (e *Engine) handlePreviousFloor(event domain.Event) {
	ps := e.getOrCreatePlayer(event.PlayerID)

	if !e.ensureRegistered(event, ps) {
		return
	}

	if !e.ensurePossible(event, !ps.Dead && ps.InDungeon && !ps.Finished) {
		return
	}

	// Если этаж больше первого(нулевого)
	// Тогда можем перейти на previous этаж
	if ps.CurrentFloor > 0 {
		ps.CurrentFloor--
		ps.FloorEnterTime = event.Time
		e.logPreviousFloor(event)
		return
	}

	e.logImpossible(event)
}
