package users

type Service interface {
	GetAll() ([]User, error)
	Create(nome, sobrenome, email string, idade, altura int, ativo bool, dataDeCriacao string) (User, error)
	Update(u User) (User, error)
	Delete(id int) error
	SoftUpdate(sobrenome string, idade, id int) (User, error)
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

func (s *service) Update(u User) (User, error) {
	result, err := s.repository.Update(u)
	if err != nil {
		return User{}, err
	}
	return result, nil
}

func (s *service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) SoftUpdate(sobrenome string, idade, id int) (User, error) {
	result, err := s.repository.SoftUpdate(sobrenome, idade, id)
	
	if err != nil {
		return User{}, err
	}

	return result, nil
}

func NewService(r Repository) Service {
	return &service{                   
		repository: r,
	}
}