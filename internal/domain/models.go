package domain

import (
	"time"
)

// Отделил состояния от данных игрока
type Player struct {
	ID     int
	Health int
}

// Сделал встраивание Player
type PlayerState struct {
	Player

	Registered   bool
	InDungeon    bool
	Finished     bool
	Dead         bool
	BossEntered  bool
	BossKilled   bool
	Disqualified bool

	CurrentFloor int

	// Индекс - это этаж
	// Элемент - количество убитых монстров
	KilledOnFloor []int

	StartTime      time.Duration
	EndTime        time.Duration
	FloorEnterTime time.Duration

	// Для итоговой статистики
	FloorsClearTotal time.Duration
	FloorsCleared    int
	BossEnterTime    time.Duration
	BossKillDuration time.Duration
}

type Config struct {
	Floors   int
	Monsters int
	OpenAt   time.Duration
	Duration time.Duration
}

// Статистика по игре
type GameStats struct {
	Config  Config
	Players map[int]*PlayerState
	Logs    []string
}

type Event struct {
	Time       time.Duration
	PlayerID   int
	EventID    int
	IntParam   int
	ExtraParam string
}

type FinalRow struct {
	State     string
	PlayerID  int
	TimeSpent time.Duration
	AvgFloor  time.Duration
	BossTime  time.Duration
	HP        int
}
