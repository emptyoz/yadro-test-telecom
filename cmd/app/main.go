package main

import (
	"flag"
	"fmt"
	"log"
	"yadro-intern-test/internal/config"
	"yadro-intern-test/internal/domain"
	"yadro-intern-test/internal/engine"
	"yadro-intern-test/internal/helpers"
	"yadro-intern-test/internal/parser"
)

func main() {
	var cfgPath string
	var eventPath string
	flag.StringVar(&cfgPath, "cfg", "./config.json", "path to the config")
	flag.StringVar(&eventPath, "event", "./events", "path to the events file")
	flag.Parse()

	cfg, err := config.Load(cfgPath)
	if err != nil {
		log.Fatalf("config load error: %v", err)
	}

	stats := &domain.GameStats{
		Config: cfg,
	}
	eng := engine.New(stats)

	events, err := parser.ParseEvent(eventPath)
	if err != nil {
		log.Fatalf("parse events error: %v", err)
	}

	eng.Process(events)

	for _, log := range stats.Logs {
		fmt.Println(log)
	}

	fmt.Println()
	fmt.Println("Final report:")
	for _, row := range eng.BuildFinalReport() {
		fmt.Printf("[%s] %d [%s, %s, %s] HP:%d\n",
			row.State,
			row.PlayerID,
			helpers.FormatTime(row.TimeSpent),
			helpers.FormatTime(row.AvgFloor),
			helpers.FormatTime(row.BossTime),
			row.HP,
		)
	}

}
