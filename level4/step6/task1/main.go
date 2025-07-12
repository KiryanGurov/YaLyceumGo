package main

import "fmt"

type (
	Animal interface {
		MakeSound() string // возвращает строку со звуком животного
		GetName() string   // возвращает строку с именем животного
		GetInfo() string   // возвращает строку с информацией о животном в формате: Имя: *name*, Вид: *species*, Возраст: *age*
	}

	animal struct {
		name    string // имя животного
		species string // вид (например, "Тигр", "Пингвин")
		age     int    // возраст
		sound   string // звук, который издает животное (например, «Ррр» или «Кря»)
		Animal
	}
	ZooKeeper struct{}
)

func (a animal) MakeSound() string {
	return a.sound
}

func (a animal) GetName() string {
	return a.name
}

func (a animal) GetInfo() string {
	return fmt.Sprintf("Имя: %s, Вид: %s, Возраст: %d", a.name, a.species, a.age)
}

func NewAnimal(name, species string, age int, sound string) Animal {
	return &animal{
		name:    name,
		species: species,
		age:     age,
		sound:   sound,
	}
}

func ZooShow(animals []Animal) {
	for _, a := range animals {
		fmt.Println(a.GetInfo())
		fmt.Println(a.MakeSound())
	}
}

func (zf ZooKeeper) Feed(animal Animal) {
	fmt.Printf("Смотритель зоопарка кормит %s. %s!", animal.GetName(), animal.MakeSound())
}
