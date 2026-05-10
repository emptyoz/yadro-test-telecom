package engine

// Фиксируем итог незавершенных игроков по времени закрытия данжа
func (e *Engine) finalizeByCloseTime() {
	closeAt := e.stats.Config.OpenAt + e.stats.Config.Duration

	for _, ps := range e.stats.Players {
		if ps.InDungeon && !ps.Finished {
			ps.Finished = true
			ps.InDungeon = false
			ps.EndTime = closeAt
		}
	}
}
