package main

import (
	"errors"
	"strings"
	"time"
	"unicode/utf8"
)

var Weekdays = []string{"Воскресенье", "Понедельник", "Вторник", "Среда", "Четверг", "Пятница", "Суббота"}

func currentDayOfTheWeek() string {
	return Weekdays[TimeNow().Weekday()]
}

func dayOrNight() string {
	now := TimeNow().Hour()
	if now > 9 && now < 23 {
		return "День"
	}
	return "Ночь"
}

func nextFriday() int {
	return int(time.Friday) - int(TimeNow().Weekday())
}

func CheckCurrentDayOfTheWeek(answer string) bool {
	if strings.EqualFold(answer, currentDayOfTheWeek()) {
		return true
	}
	return false
}

func CheckNowDayOrNight(answer string) (bool, error) {
	if utf8.RuneCountInString(answer) != 4 {
		return false, errors.New("исправь свой ответ, а лучше ложись поспать")
	}
	if strings.EqualFold(answer, dayOrNight()) {
		return true, nil
	}
	return false, nil
}
