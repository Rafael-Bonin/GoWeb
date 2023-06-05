package users

type Service interface {
	GetAll() ([]User, error)
	Create(nome, sobrenome, email string, idade, altura int, ativo bool, dataDeCriacao string) (User, error)
	Update(nome, sobrenome, email string, idade, altura int, ativo bool, dataDeCriacao string)
}

type service struct {
	repository Repository
}

func (s *service) GetAll() ([]User, error) {
	result, err := s.repository.GetAll()
	return result, err
}

func (s *service) Create(nome, sobrenome, email string, idade, altura int, ativo bool, dataDeCriacao string) (User, error) {
	user, err := s.repository.Create(nome, sobrenome, email, idade, altura, ativo, dataDeCriacao)
	return user, err
}

func (s *service) Update(nome, sobrenome, email string, idade, altura int, ativo bool, dataDeCriacao string) {

}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}