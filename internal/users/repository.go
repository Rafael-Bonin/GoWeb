package users

import "errors"

type User struct {
	Id int `json:"id"`
	Nome string `json:"nome"`
	Sobrenome string `json:"sobrenome"`
	Email string `json:"email"`
	Idade int `json:"idade"`
	Altura int `json:"altura"`
	Ativo bool `json:"ativo"`
	DataDeCriacao string `json:"dataDeCriacao"`
}

type Repository interface {
	GetAll() ([]User, error)
	Create(nome, sobrenome, email string, idade, altura int, ativo bool, dataDeCriacao string) (User, error)
	Update(u User) (User, error)
	Delete(id int) error
	SoftUpdate(sobrenome string, idade, id int) (User, error)
}

type repository struct {}

var AllUsers []User = []User{}
var id int = 0

func (r *repository) GetAll() ([]User, error) {
	return AllUsers, nil
}

func (r *repository) Create(nome, sobrenome, email string, idade, altura int, ativo bool, dataDeCriacao string) (User, error) {
	id++
	newUser := User{Id: id, Nome: nome, Sobrenome: sobrenome, Email: email, Idade: idade, Altura: altura, Ativo: ativo, DataDeCriacao: dataDeCriacao}
	AllUsers = append(AllUsers, newUser)

	return newUser, nil
}

func (r *repository) Update(u User) (User, error) {
	var found bool = false
	var updatedUser User

	for i, value := range AllUsers {
		if value.Id == u.Id {
			AllUsers[i] = User{Id: u.Id, Nome: u.Nome, Sobrenome: u.Sobrenome, Email: u.Email, Idade: u.Idade, Altura: u.Altura, Ativo: u.Ativo, DataDeCriacao: u.DataDeCriacao}
			found = true
			updatedUser = AllUsers[i]
			break
		}
	}

	if !found {
		return User{}, errors.New("Nao foram encontrados usuarios com esse id")
	}

	return updatedUser, nil
}

func (r *repository) Delete(id int) error {
	deleted := false
	var index int

	for i, value := range AllUsers {
		if value.Id == id {
			index = i
			deleted = true

		}
	}

	if !deleted {
		return errors.New("nao foi encontrado um usuario com este id")
	}

	AllUsers = append(AllUsers[:index], AllUsers[index + 1:]...)
	return nil
} 

func (r *repository) SoftUpdate(sobrenome string, idade, id int) (User, error) {
	var found bool = false
	var updatedUser User

	for i, value := range AllUsers {
		if value.Id == id {
			AllUsers[i].Sobrenome = sobrenome
			AllUsers[i].Idade = idade
			found = true
			updatedUser = AllUsers[i]
			break
		}
	}

	if !found {
		return User{}, errors.New("Nao foram encontrados usuarios com esse id")
	}

	return updatedUser, nil
}

func NewRepository() Repository {
	return &repository{}
}