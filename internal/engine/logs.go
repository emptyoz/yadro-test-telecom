package engine

import (
	"fmt"
	"yadro-intern-test/internal/domain"
	"yadro-intern-test/internal/helpers"
)

func (e *Engine) logRegistered(event domain.Event) {
	e.stats.Logs = append(e.stats.Logs,
		fmt.Sprintf("[%s] Player [%d] registered",
			helpers.FormatTime(event.Time), event.PlayerID))
}

func (e *Engine) logEnteredDungeon(event domain.Event) {
	e.stats.Logs = append(e.stats.Logs,
		fmt.Sprintf("[%s] Player [%d] entered the dungeon",
			helpers.FormatTime(event.Time), event.PlayerID))
}

func (e *Engine) logDamage(event domain.Event, damage int) {
	e.stats.Logs = append(e.stats.Logs,
		fmt.Sprintf("[%s] Player [%d] recieved [%d] of damage",
			helpers.FormatTime(event.Time), event.PlayerID, damage))
}

func (e *Engine) logHeal(event domain.Event, heal int) {
	e.stats.Logs = append(e.stats.Logs,
		fmt.Sprintf("[%s] Player [%d] has restored [%d] of health",
			helpers.FormatTime(event.Time), event.PlayerID, heal))
}

func (e *Engine) logDead(event domain.Event) {
	e.stats.Logs = append(e.stats.Logs,
		fmt.Sprintf("[%s] Player [%d] is dead",
			helpers.FormatTime(event.Time), event.PlayerID))
}

func (e *Engine) logDisqualified(event domain.Event) {
	e.stats.Logs = append(e.stats.Logs,
		fmt.Sprintf("[%s] Player [%d] is disqualified",
			helpers.FormatTime(event.Time), event.PlayerID))
}

func (e *Engine) logImpossible(event domain.Event) {
	e.stats.Logs = append(e.stats.Logs,
		fmt.Sprintf("[%s] Player [%d] makes imposible move [%d]",
			helpers.FormatTime(event.Time), event.PlayerID, event.EventID))
}

func (e *Engine) logNextFloor(event domain.Event) {
	e.stats.Logs = append(e.stats.Logs,
		fmt.Sprintf("[%s] Player [%d] went to the next floor",
			helpers.FormatTime(event.Time), event.PlayerID))
}

func (e *Engine) logPreviousFloor(event domain.Event) {
	e.stats.Logs = append(e.stats.Logs,
		fmt.Sprintf("[%s] Player [%d] went to the previous floor",
			helpers.FormatTime(event.Time), event.PlayerID))
}

func (e *Engine) logKillMonster(event domain.Event) {
	e.stats.Logs = append(e.stats.Logs,
		fmt.Sprintf("[%s] Player [%d] killed the monster",
			helpers.FormatTime(event.Time), event.PlayerID))
}

func (e *Engine) logLeaveDungeon(event domain.Event) {
	e.stats.Logs = append(e.stats.Logs,
		fmt.Sprintf("[%s] Player [%d] left the dungeon",
			helpers.FormatTime(event.Time), event.PlayerID))
}

func (e *Engine) logEnterBossFloor(event domain.Event) {
	e.stats.Logs = append(e.stats.Logs,
		fmt.Sprintf("[%s] Player [%d] entered the boss's floor",
			helpers.FormatTime(event.Time), event.PlayerID))
}

func (e *Engine) logKillBoss(event domain.Event) {
	e.stats.Logs = append(e.stats.Logs,
		fmt.Sprintf("[%s] Player [%d] killed the boss",
			helpers.FormatTime(event.Time), event.PlayerID))
}

func (e *Engine) logCannotContinue(event domain.Event, reason string) {
	e.stats.Logs = append(e.stats.Logs,
		fmt.Sprintf("[%s] Player [%d] cannot continue due to [%s]",
			helpers.FormatTime(event.Time), event.PlayerID, reason))
}
