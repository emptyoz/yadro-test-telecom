package engine

import (
	"yadro-intern-test/internal/domain"
)

func (e *Engine) handleDamage(event domain.Event) {
	ps := e.getOrCreatePlayer(event.PlayerID)

	// Регистрация обязательна
	// Для получения урона
	if !e.ensureRegistered(event, ps) {
		return
	}

	// Обязательные состояния
	if !e.ensurePossible(event,
		!ps.Dead && ps.InDungeon && !ps.Finished) {
		return
	}

	// Проверка на смерть
	// Вычисление HP
	e.isDead(ps, event)
}

func (e *Engine) handleHeal(event domain.Event) {
	ps := e.getOrCreatePlayer(event.PlayerID)

	if !e.ensureRegistered(event, ps) {
		return
	}

	if !e.ensurePossible(event,
		!ps.Dead && ps.InDungeon && !ps.Finished) {
		return
	}

	heal := event.IntParam

	e.logHeal(event, heal)

	// Исцеление игрока
	// Если все состояния в норме
	ps.Health += heal
	if ps.Health > 100 {
		ps.Health = 100
	}
}

func (e *Engine) handleKillMonster(event domain.Event) {
	ps := e.getOrCreatePlayer(event.PlayerID)

	if !e.ensureRegistered(event, ps) {
		return
	}

	if !e.ensurePossible(event,
		!ps.Dead && !ps.BossEntered && ps.InDungeon && !ps.Finished &&
			ps.KilledOnFloor[ps.CurrentFloor] < e.stats.Config.Monsters) {
		return
	}

	ps.KilledOnFloor[ps.CurrentFloor]++
	e.logKillMonster(event)

	if ps.KilledOnFloor[ps.CurrentFloor] == e.stats.Config.Monsters {
		ps.FloorsClearTotal += event.Time - ps.FloorEnterTime
		ps.FloorsCleared++
	}
}
