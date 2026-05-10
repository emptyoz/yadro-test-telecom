package parser

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"yadro-intern-test/internal/domain"
	"yadro-intern-test/internal/helpers"
)

func ParseEvent(path string) ([]domain.Event, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open events %q: %w", path, err)
	}
	defer file.Close()

	var events []domain.Event
	scanner := bufio.NewScanner(file)

	lineNo := 0
	for scanner.Scan() {
		lineNo++
		line := scanner.Text()

		event, ok, err := parseLine(line, lineNo)
		if err != nil {
			return nil, fmt.Errorf("parse events %q: %w", path, err)
		}
		if !ok {
			continue
		}

		events = append(events, event)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scan events %q: %w", path, err)
	}

	return events, nil
}

// Парсинг файла event
// Построчно: [14:00:00] 1 1
func parseLine(line string, lineNo int) (domain.Event, bool, error) {
	fields := strings.Fields(line)

	if len(fields) < 3 {
		return domain.Event{}, false, nil
	}

	eventTime, err := helpers.TimeIntoDuration(fields[0])
	if err != nil {
		return domain.Event{}, false, fmt.Errorf("line %d: invalid time %q: %w", lineNo, fields[0], err)
	}

	playerID, err := strconv.Atoi(fields[1])
	if err != nil {
		return domain.Event{}, false, fmt.Errorf("line %d: invalid player id %q: %w", lineNo, fields[1], err)
	}

	eventID, err := strconv.Atoi(fields[2])
	if err != nil {
		return domain.Event{}, false, fmt.Errorf("line %d: invalid event id %q: %w", lineNo, fields[2], err)
	}

	intParam, strParam := parseExtra(fields)

	return domain.Event{
		Time:       eventTime,
		PlayerID:   playerID,
		EventID:    eventID,
		IntParam:   intParam,
		ExtraParam: strParam,
	}, true, nil
}

// Парсинг extra параметров
// В IntParams и ExtraParam
// ExtraParam может быть multistring
func parseExtra(fields []string) (int, string) {
	if len(fields) == 3 {
		return 0, ""
	}

	extra := strings.Join(fields[3:], " ")

	if n, err := strconv.Atoi(extra); err == nil {
		return n, ""
	}

	return 0, extra
}
