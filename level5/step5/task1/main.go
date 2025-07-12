package main

import (
	"slices"
	"strings"
)

type (
	CalcRating interface {
		calculateRating()
		calculateRatio() float64
	}

	Player struct {
		Name    string
		Goals   int
		Misses  int
		Assists int
		Rating  float64
	}
)

func (p *Player) calculateRating() {
	if p.Misses != 0 {
		p.Rating = (float64(p.Goals) + float64(p.Assists)/2) / float64(p.Misses)
	} else {
		p.Rating = float64(p.Goals) + float64(p.Assists)/2
	}
}

func (p *Player) calculateRatio() float64 {
	if p.Misses != 0 {
		return float64(p.Goals) / float64(p.Misses)
	} else {
		return float64(p.Goals)
	}
}

func NewPlayer(name string, goals, misses, assists int) Player {
	p := Player{Name: name, Goals: goals, Misses: misses, Assists: assists, Rating: 0.0}
	p.calculateRating()
	return p
}

// Убыванию количества голов
func goalsSort(players []Player) []Player {
	slices.SortFunc(players, func(a, b Player) int {
		if a.Goals != b.Goals {
			return b.Goals - a.Goals
		}
		return strings.Compare(a.Name, b.Name)
	})
	return players
}

// Убыванию рейтинга
func ratingSort(players []Player) []Player {
	slices.SortFunc(players, func(a, b Player) int {
		if a.Rating != b.Rating {
			if a.Rating < b.Rating {
				return 1
			}
			return -1
		}
		return strings.Compare(a.Name, b.Name)

	})
	return players
}

// Убыванию отношения голов к промахам
func gmSort(players []Player) []Player {
	slices.SortFunc(players, func(a, b Player) int {
		a_gm, b_gm := a.calculateRatio(), b.calculateRatio()
		if a_gm != b_gm {

			if a_gm < b_gm {
				return 1
			}
			return -1
		}
		return strings.Compare(a.Name, b.Name)
	})
	return players
}
