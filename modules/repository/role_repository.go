package repository

import (
	"api/modules/model"
	"time"

	"gorm.io/gorm"
)

type RoleRepository interface {
	Save(role model.Role) (*model.Role, error)
	FindAll() (*[]model.Role, error)
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
	role.CreateAt = time.Now()
	role.UpdateAt = time.Now()

	if err := r.conn.Create(&role).Error; err != nil {
		return nil, err
	}

	return &role, nil
}

func (r *roleRepository) FindAll() (*[]model.Role, error) {
	var roles []model.Role

	err := r.conn.Find(&roles).Error
	if err != nil {
		return nil, err
	}

	return &roles, err
}