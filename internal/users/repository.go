package users

import (
	"errors"

	"github.com/Rafael-Bonin/GoWeb/pkg/store"
)

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

type repository struct {
	db store.Store
}

func (r *repository) GetAll() ([]User, error) {
	var allUsers []User
	r.db.Read(&allUsers)
	return allUsers, nil
}

func (r *repository) Create(nome, sobrenome, email string, idade, altura int, ativo bool, dataDeCriacao string) (User, error) {
	var allUsers []User
	r.db.Read(&allUsers)
	lastId, err := r.lastId()
	if err != nil {
		return User{}, err
	}
	newUser := User{Id: lastId, Nome: nome, Sobrenome: sobrenome, Email: email, Idade: idade, Altura: altura, Ativo: ativo, DataDeCriacao: dataDeCriacao}
	allUsers = append(allUsers, newUser)
	if er := r.db.Write(allUsers); er != nil {
		return User{}, er
	}

	return newUser, nil
}

func (r *repository) Update(u User) (User, error) {
	var found bool = false
	var updatedUser User
	var allUsers []User
	r.db.Read(&allUsers)

	for i, value := range allUsers {
		if value.Id == u.Id {
			allUsers[i] = User{Id: u.Id, Nome: u.Nome, Sobrenome: u.Sobrenome, Email: u.Email, Idade: u.Idade, Altura: u.Altura, Ativo: u.Ativo, DataDeCriacao: u.DataDeCriacao}
			found = true
			updatedUser = allUsers[i]
			break
		}
	}

	if !found {
		return User{}, errors.New("Nao foram encontrados usuarios com esse id")
	}

	if err := r.db.Write(allUsers); err != nil {
		return User{}, err
	}

	return updatedUser, nil
}

func (r *repository) Delete(id int) error {
	deleted := false
	var index int
	var allUsers []User
	r.db.Read(&allUsers)

	for i, value := range allUsers {
		if value.Id == id {
			index = i
			deleted = true

		}
	}

	if !deleted {
		return errors.New("nao foi encontrado um usuario com este id")
	}

	allUsers = append(allUsers[:index], allUsers[index + 1:]...)
	
	if err := r.db.Write(allUsers); err != nil {
		return err
	}

	return nil
} 

func (r *repository) SoftUpdate(sobrenome string, idade, id int) (User, error) {
	var found bool = false
	var updatedUser User
	var allUsers []User
	r.db.Read(&allUsers)

	for i, value := range allUsers {
		if value.Id == id {
			allUsers[i].Sobrenome = sobrenome
			allUsers[i].Idade = idade
			found = true
			updatedUser = allUsers[i]
			break
		}
	}

	if !found {
		return User{}, errors.New("Nao foram encontrados usuarios com esse id")
	}

	if err := r.db.Write(allUsers); err != nil {
		return User{}, err
	}

	return updatedUser, nil
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) lastId() (int, error) {
	var allUsers []User
	if err := r.db.Read(&allUsers); err != nil {
		return 1, nil
	}
	if len(allUsers) < 1 {
		return 1, nil
	}
	return allUsers[len(allUsers) - 1].Id + 1, nil
}