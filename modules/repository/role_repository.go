package repository

import (
	"api/modules/model"
	"time"

	"gorm.io/gorm"
)

type RoleRepository interface {
	Save(role model.Role) (*model.Role, error)
	FindAll() (*[]model.Role, error)
	FindById(id uint64) (*model.Role, error)
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
	role.UpdateAt = time.Now()
	if role.ID == 0 {
		role.CreateAt = time.Now()
	}

	if err := r.conn.Save(&role).Error; err != nil {
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

func (r *roleRepository) FindById(id uint64) (*model.Role, error) {
	var role model.Role

	err := r.conn.Where("id = ?", id).First(&role).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &role, nil
}
