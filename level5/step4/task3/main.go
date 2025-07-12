package main

import (
	"errors"
	"fmt"
	"slices"
)

type (
	Worker struct {
		name       string
		position   string
		salary     uint
		experience uint
	}

	Company struct {
		Workers []Worker
	}

	CompanyInterface interface {
		AddWorkerInfo(name, position string, salary, experience uint) error
		SortWorkers() ([]string, error)
	}
)

var (
	positions = []string{"директор", "зам. директора", "начальник цеха", "мастер", "рабочий"}
)

func getPosition(position string) int {
	for i, p := range positions {
		if p == position {
			return i
		}
	}
	return -1
}

func (c *Company) AddWorkerInfo(name, position string, salary, experience uint) error {
	if getPosition(position) == -1 {
		return errors.New("ERROR: POSITION IS INCORRECT!")
	}
	if salary < 1 {
		return errors.New("ERROR: SALARY IS WRONG!")
	}
	c.Workers = append(c.Workers, Worker{name, position, salary, experience})
	return nil
}

func (c *Company) SortWorkers() ([]string, error) {
	// сортирует список сотрудников по доходу за время работы на предприятии (по убыванию),
	// должности (от высокой до низкой) и возвращает слайс формата: *имя* — *доход* — *должность*
	slices.SortFunc(c.Workers, func(a, b Worker) int {
		incomeA := a.salary * a.experience
		incomeB := b.salary * b.experience
		rankA := getPosition(a.position)
		rankB := getPosition(b.position)
		if incomeA != incomeB {
			if incomeA > incomeB {
				return -1
			}
			return 1
		}
		if rankA != rankB {
			if rankA < rankB {
				return -1
			}
			return 1
		}
		return 0
	})

	var res []string
	for _, w := range c.Workers {
		res = append(res, fmt.Sprintf("%s — %d — %s", w.name, w.experience*w.salary, w.position))
	}
	return res, nil
}
