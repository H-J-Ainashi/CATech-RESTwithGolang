package database

import (
	"github.com/H-J-Ainashi/CATech-RESTwithGolang/domain"
)

type UserRepository struct {
	SqlHandler
}

func (userRepos *UserRepository) FindById(id int) (user domain.User, err error) {
	err := userRepos.Find(&user, id).Error
	if err != nil {
		return
	}
	return
}

func (userRepos *UserRepository) Store(u domain.User) (user domain.User, err error) {
	err := userRepos.Create(&u).Error
	if err != nil {
		return
	}
	user = u
	return
}

func (userRepos *UserRepository) Update(u domain.User) (user domain.User, err error) {
	err := userRepos.Update(&u).Error
	if err != nil {
		return
	}
	user = u
	return
}

func (userRepos *UserRepository) DeleteById(user domain.User) (err error) {
	err := userRepos.Delete(&user).Error
	if err != nil {
		return
	}
	return
}

func (userRepos *UserRepository) FindAll() (users domain.Users, err error) {
	err := userRepos.Find(&users).Error
	if err != nil {
		return
	}
	return
}
