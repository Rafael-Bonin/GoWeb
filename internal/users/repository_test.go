package users

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubStore struct{}

func (s *StubStore) Read(data interface{}) error {
	data = []User{
		{ 
			Id: 1,
			Nome: "Rafael",
			Sobrenome: "Bonin",
			Email: "rafalyf11@gmail.com",
			Idade: 21 ,
			Altura: 179 ,
			Ativo: true,
			DataDeCriacao: "11/01/2002",
		},
		{
			Id: 2,
			Nome: "A",
			Sobrenome: "B",
			Email: "C",
			Idade: 4,
			Altura: 5,
			Ativo: true,
			DataDeCriacao: "a", 
		},
	}

	return nil
}

func (s StubStore) Write(data interface{}) error {
	return nil
}

func TestGetAll(t *testing.T) {
	db := &StubStore{}
	repository := NewRepository(db)

	result, _ := repository.GetAll()
	fmt.Println(result)
	expectedResult := []User{
}

	assert.Equal(t, expectedResult, result, "devem ser iguais")
}