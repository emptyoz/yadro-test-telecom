package engine

import (
	"yadro-intern-test/internal/domain"
)

type Engine struct {
	stats *domain.GameStats
}

func New(stats *domain.GameStats) *Engine {
	return &Engine{
		stats: stats,
	}
}

func (e *Engine) Process(events []domain.Event) {
	for _, event := range events {
		switch event.EventID {
		case 1:
			e.handleRegister(event)
		case 2:
			e.handleDungeonEnter(event)
		case 3:
			e.handleKillMonster(event)
		case 4:
			e.handleNextFloor(event)
		case 5:
			e.handlePreviousFloor(event)
		case 6:
			e.handleEnterBossFloor(event)
		case 7:
			e.handleKillBoss(event)
		case 8:
			e.handleLeaveDungeon(event)
		case 9:
			e.handleCannotContinue(event)
		case 10:
			e.handleHeal(event)
		case 11:
			e.handleDamage(event)
		default:
			e.logImpossible(event)
		}
	}

	e.finalizeByCloseTime()
}

// Проверка на смерть
func (e *Engine) isDead(ps *domain.PlayerState, event domain.Event) {
	damage := event.IntParam
	healthAfter := ps.Health - damage

	e.logDamage(event, damage)

	if healthAfter <= 0 {
		ps.Health = 0
		ps.Dead = true
		ps.Finished = true
		ps.EndTime = event.Time

		e.logDead(event)
		return
	}
	ps.Health = healthAfter
}
