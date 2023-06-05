package users

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
	Update(nome, sobrenome, email string, idade, altura int, ativo bool, dataDeCriacao string) 
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

func (r *repository) Update(nome, sobrenome, email string, idade, altura int, ativo bool, dataDeCriacao string) {

}

func NewRepository() Repository {
	return &repository{}
}