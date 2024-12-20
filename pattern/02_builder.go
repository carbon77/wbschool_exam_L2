package pattern

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

// Структура пользователя
type User struct {
	firstName   string
	lastName    string
	age         int
	hasChildren bool
}

func (u User) String() string {
	return fmt.Sprintf("User{name=%s %s, age=%d, hasChildren=%v}", u.firstName, u.lastName, u.age, u.hasChildren)
}

// Структура строителя (Builder) пользователя, инкапсулирует создания пользователя
type UserBuilder struct {
	user *User
}

func NewUserBuilder() *UserBuilder {
	return &UserBuilder{&User{}}
}

func (b *UserBuilder) Build() *User {
	return b.user
}

func (b *UserBuilder) FirstName(firstName string) *UserBuilder {
	b.user.firstName = firstName
	return b
}

func (b *UserBuilder) LastName(lastName string) *UserBuilder {
	b.user.lastName = lastName
	return b
}

func (b *UserBuilder) Age(age int) *UserBuilder {
	b.user.age = age
	return b
}

func (b *UserBuilder) HasChildren(hasChildren bool) *UserBuilder {
	b.user.hasChildren = hasChildren
	return b
}

func TestBuilder() {
	user1 := NewUserBuilder().FirstName("Igor").LastName("Zakatov").Age(21).Build()
	user2 := NewUserBuilder().FirstName("Dima").Age(30).HasChildren(true).LastName("Ivanov").Build()
	fmt.Printf("Built user: %v\n", user1)
	fmt.Printf("Built user: %v\n", user2)
}
