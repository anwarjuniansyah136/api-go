package repository

import (
	"api/modules/model"
	"time"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user model.User) (*model.User, error)
	FindById(id uint64) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	FindAll() (*[]model.User, error)
	DeleteByID(id uint64) error
	Delete(user *model.User) error
}

type userRepository struct {
	conn *gorm.DB
}


func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		conn: db,
	}
}

func (u *userRepository) Save(user model.User) (*model.User, error) {
	user.UpdateAt = time.Now()

	if user.ID == 0 {
		user.CreateAt = time.Now()
	}

	if err := u.conn.Save(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) FindById(id uint64) (*model.User, error) {
	var user model.User
	err := u.conn.Where("id = ?", id).First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := u.conn.Where("LOWER(email) = LOWER(?)", email).First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) FindAll() (*[]model.User, error) {
	var users []model.User
	err := u.conn.Find(&users).Error

	if err != nil {
		return nil, err
	}

	return &users, nil
}

func (u *userRepository) DeleteByID(id uint64) error {
	result := u.conn.Delete(&model.User{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (u *userRepository) Delete(user *model.User) error {
	result := u.conn.Delete(user)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0{
		return gorm.ErrRecordNotFound
	}

	return nil
}