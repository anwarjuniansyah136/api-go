package service

import (
	"api/helper"
	"api/modules/model"
	"api/modules/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserService interface {
	Create(ctx *gin.Context)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{
		repository: repository.NewUserRepository(db),
	}
}

func (u *userService) Create(ctx *gin.Context) {
	var input model.UserCreateRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := model.User{
		FullName: input.FullName,
		Email: input.Email,
		Password: helper.HashedPassword(input.Password),
	}

	result, err := u.repository.Save(user)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	createEmail(input.FullName, input.Email)

	ctx.JSON(http.StatusOK, result)
}

func createEmail(name, email string){
	body := `
		<h2>Selamat datang, ` + name + `!</h2>
		<p>Akun kamu berhasil dibuat dengan email: ` + email + `</p>
	`
	helper.SendEmail(email, "Success Create Account", body)
}
