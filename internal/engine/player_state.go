package engine

import "yadro-intern-test/internal/domain"

func (e *Engine) getOrCreatePlayer(playerID int) *domain.PlayerState {
	// Инициализация map
	// Статистика по игрокам
	if e.stats.Players == nil {
		e.stats.Players = make(map[int]*domain.PlayerState)
	}

	// Если state игрока найден - вернем
	if ps, ok := e.stats.Players[playerID]; ok {
		return ps
	}

	// Не найдено - добавим
	ps := &domain.PlayerState{
		Player: domain.Player{
			ID:     playerID,
			Health: 100,
		},
	}

	// Возвращаем созданное state игрока
	e.stats.Players[playerID] = ps
	return ps
}
