package repository

import (
	"core/db"
	"core/pkg/domain"
)

type UserRepository struct {
	userTable *db.TableAccessor
}

func NewUserRepository(userTable *db.TableAccessor) *UserRepository {
	return &UserRepository{userTable}
}

func (r *UserRepository) GetUserByID(id int) (*domain.User, error) {
	_, err := r.userTable.Query("condition", id)
	if err != nil {
		return nil, err
	}
	return &domain.User{ /*...*/ }, nil
}
