package service

import (
	"api/helper"
	"api/modules/model"
	"api/modules/repository"
	"net/http"
	"strconv"

	// "strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserService interface {
	Create(ctx *gin.Context)
	GetAllUser(ctx *gin.Context)
	FindById(ctx *gin.Context)
	// Update(ctx *gin.Context)
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
		Email:    input.Email,
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

func (u *userService) GetAllUser(ctx *gin.Context) {
	result, err := u.repository.FindAll()
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, result)
}

func createEmail(name, email string) {
	body := `
		<h2>Selamat datang, ` + name + `!</h2>
		<p>Akun kamu berhasil dibuat dengan email: ` + email + `</p>
	`
	helper.SendEmail(email, "Success Create Account", body)
}

func (u *userService) FindById(ctx *gin.Context) {
	var id = ctx.Param("id")
	
	value, err := strconv.ParseUint(id, 10, 16)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : err,
		})
		return
	}

	result, err := u.repository.FindById(value)
	if err == nil {
		if result == nil{
			ctx.JSON(http.StatusNotFound, gin.H{
				"error":"teacher not found",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error" : "something wrong in our server",
		})
	}

	ctx.JSON(http.StatusOK, result)
}