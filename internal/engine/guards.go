package engine

import (
	"yadro-intern-test/internal/domain"
)

// Проверка обязательной регистрации
func (e *Engine) ensureRegistered(event domain.Event, ps *domain.PlayerState) bool {
	if ps.Disqualified || ps.Finished {
		return false
	}

	if ps.Registered {
		return true
	}

	ps.Disqualified = true
	ps.Finished = true
	e.logDisqualified(event)
	return false
}

// Проверка обязательных state
func (e *Engine) ensurePossible(event domain.Event, ok bool) bool {
	if ok {
		return true
	}

	e.logImpossible(event)
	return false
}
