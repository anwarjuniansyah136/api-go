package repository

import (
	"api/modules/model"

	"gorm.io/gorm"
)

type RoleRepository interface {
	Save(role model.Role) (*model.Role, error)
}

type roleRepository struct {
	conn *gorm.DB
}

func NewroleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{
		conn: db,
	}
}

func (r *roleRepository) Save(role model.Role) (*model.Role, error) {
	panic("unimplemented")
}