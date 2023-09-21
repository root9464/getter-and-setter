package main

import "fmt"

type user struct {
	id      int16
	name    string
	surname string
}

func debug(u user) {
	fmt.Printf("%+v\n", u)
}

// notify реализует метод вывода.
func (u user) notify() {
	fmt.Println("Система: добавлен", u.name, u.surname)
}

// реализация геттер свойств  .
func (u *user) changeSurname(surname string) {
	u.surname = surname
}

func (u *user) changeName(name string) {
	u.name = name
}

func (u *user) Send() {
	u.notify()
}

// main - это точка входа для приложения.
func main() {
	// Указатели типа user также могут использоваться для методов
	// объявлен с получателем значения.

	user1 := &user{1, "Egor", "Chervony"}
	user1.changeName("Егор")
	user1.changeSurname("Червоный")
	user1.notify()
	debug(user{1, "Egor", "Chervony"})
	fmt.Println("или же" + "{id:1 name:Egor surname:Chervony}") // не нашел ftm свойство чтобы читать много обьектов
}
