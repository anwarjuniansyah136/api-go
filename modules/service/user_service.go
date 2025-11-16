package service

import (
	"api/helper"
	"api/modules/jwt"
	"api/modules/model"
	"api/modules/repository"
	"api/modules/request"
	"api/modules/response"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserService interface {
	Create(ctx *gin.Context)
	GetAllUser(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Login(ctx *gin.Context)
	ForgotPassword(ctx *gin.Context)
	ConfirmationCode(ctx *gin.Context)
	ResetPassword(ctx *gin.Context)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{
		repository: repository.NewUserRepository(db),
	}
}

func (u *userService) ResetPassword(ctx *gin.Context) {
	panic("unimplemented")
}

func (u *userService) ConfirmationCode(ctx *gin.Context) {
	var inputRequest request.UserOTP

	if err := ctx.ShouldBindJSON(&inputRequest); err != nil {
		response.Error(ctx, 400, "Bad Request", err)
		return
	}

	result, err := u.repository.FindByEmail(inputRequest.Email)
	if err != nil {
		response.Error(ctx, 404, "Not Found", err)
		return
	}

	if !binding(inputRequest.Code, result.Code) {
		response.Error(ctx, 422, "Code False", err)
		return
	}

	if time.Now().After(result.ExpiresAt) {
		response.Error(ctx, 410, "Expired Code", err)
		return
	}

	response.Success(ctx, result)
}

func (u *userService) ForgotPassword(ctx *gin.Context) {
	var query = ctx.Param("query")
	result, err := u.repository.FindByEmail(query)
	if err != nil {
		response.Error(ctx, 404, "Not Found", err)
		return
	}

	code := helper.GenerateOTP()
	result.Code = code
	result.ExpiresAt = time.Now().Add(10 * time.Minute)

	helper.SendOTPByEmail(query, "Kode Reset Password Kamu : "+code)
	u.repository.Save(*result)

	response.Success(ctx, result)
}

func (u *userService) Login(ctx *gin.Context) {
	var loginRequest request.UserLogin

	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		response.Error(ctx, 400, "Bad Request", err)
		return
	}

	result, err := u.repository.FindByEmail(loginRequest.Username)
	if err != nil {
		response.Error(ctx, 404, "Not Found", err)
		return
	}

	if result == nil {
		response.Error(ctx, 404, "Not Found", err)
		return
	}

	if !helper.CheckPassword(result.Password, loginRequest.Password) {
		response.Error(ctx, 400, "Invalid Request", err)
		return
	}

	token, err := jwt.GenerateToken(result.ID, result.Email)
	if err != nil {
		response.Error(ctx, 500, "Failed To Create Token", err)
	}

	responseFromBE := response.UserResponse{
		ID: result.ID,
		Email: result.Email,
		Token: token,
	}

	response.Success(ctx, responseFromBE)
}

func (u *userService) Create(ctx *gin.Context) {
	var input request.UserCreateRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		response.Error(ctx, 400, "Bad Request", err)
		return
	}

	user := model.User{
		FullName: input.FullName,
		Email:    input.Email,
		Password: helper.HashedPassword(input.Password),
	}

	result, err := u.repository.Save(user)
	if err != nil {
		response.Error(ctx, 422, "Invalid Request", err)
		return
	}

	createEmail(input.FullName, input.Email)

	response.Success(ctx, result)
}

func (u *userService) GetAllUser(ctx *gin.Context) {
	result, err := u.repository.FindAll()
	if err != nil {
		response.Error(ctx, 422, "Invalid Request", err)
		return
	}

	response.Success(ctx, result)
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
			"error": err,
		})
		return
	}

	result, err := u.repository.FindById(value)
	if err == nil {
		if result == nil {
			response.Error(ctx, 404, "Not Found", err)
			return
		}
		response.Error(ctx, 500, "Internal Server Error", err)
		return
	}

	response.Success(ctx, result)
}

func binding(codeFromUser, codeFromData string) bool {
	return codeFromData == codeFromUser
}
