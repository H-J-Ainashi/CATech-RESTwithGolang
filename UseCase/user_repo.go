package usecase

import (
	"github.com/H-J-Ainashi/CATech-RESTwithGolang/domain"
)

type UserRepository interface {
	FindById(id int) (domain.User, error)
	Store(domain.User) (domain.User, error)
	Update(domain.User) (domain.User, error)
	DeleteById(domain.User) error
	FindAll() (domain.Users, error)
}
