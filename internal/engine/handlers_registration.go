package engine

import (
	"yadro-intern-test/internal/domain"
)

func (e *Engine) handleRegister(event domain.Event) {
	// Создание или получение state игрока
	ps := e.getOrCreatePlayer(event.PlayerID)

	// Если игрок уже зарегестрирован
	// Записываем ошибку в логи
	if !e.ensurePossible(event, !ps.Registered) {
		return
	}

	// Регистрируем игрока
	// Записываем в логи
	ps.Registered = true
	e.logRegistered(event)
}

func (e *Engine) handleDungeonEnter(event domain.Event) {
	ps := e.getOrCreatePlayer(event.PlayerID)

	// Регистрация игрока обязательна
	// Для вхождения в Dungeon
	if !e.ensureRegistered(event, ps) {
		return
	}

	// Проверка обязательных состояний
	if !e.ensurePossible(event, !ps.InDungeon && !ps.Finished && !ps.Disqualified) {
		return
	}

	// Инициализация slice
	// Этажи - индекс
	ps.KilledOnFloor = make([]int, e.stats.Config.Floors)

	// Бизнес-правило для входа
	// Стартовый этаж 0
	ps.CurrentFloor = 0
	ps.StartTime = event.Time
	ps.FloorEnterTime = event.Time
	ps.InDungeon = true
	e.logEnteredDungeon(event)
}
