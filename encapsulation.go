package user

import "fmt"

type user struct {
	name string
	Age  int
}

func (u user) getName() string {
	return u.name
}

func (u user) GetUserAge() string {
	return fmt.Sprintf("%s is %d years old", u.getName(), u.Age)
}

func NewUser(name string, age int) user {
	return user{name: name, Age: age}
}
