package main

import (
	"fmt"
)

type user struct {
	name string
	Age  int
}

// Contoh interface untuk kelas user
type UserService interface {
	GetUserAge() string
}

func (u user) getName() string {
	return u.name
}

func (u user) GetUserAge() string {
	return fmt.Sprintf("%s berumur %d Tahun", u.getName(), u.Age)
}

// Implementasikan interface UserService
func NewUser(name string, age int) UserService {
	return user{name: name, Age: age}
}

func main() {
	user := NewUser("Elang Muhammad Fikhri", 29)
	fmt.Println(user.GetUserAge())
}
