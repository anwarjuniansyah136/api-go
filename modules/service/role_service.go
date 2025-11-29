package service

import (
	"api/modules/model"
	"api/modules/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RoleService interface {
	Create(ctx *gin.Context)
	GetAllRole(ctx *gin.Context)
	FindById(ctx *gin.Context)
}

type roleService struct {
	repository repository.RoleRepository
}

func NewRoleService(db *gorm.DB) RoleService {
	return &roleService{
		repository: repository.NewroleRepository(db),
	}
}

func (r *roleService) Create(ctx *gin.Context) {
	var roleRequest model.RoleCreateRequest

	if err := ctx.ShouldBindJSON(&roleRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	role := model.Role{
		Name: roleRequest.Name,
	}

	result, err := r.repository.Save(role)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (r *roleService) GetAllRole(ctx *gin.Context) {
	result, err := r.repository.FindAll()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
	}

	ctx.JSON(http.StatusOK, result)
}

func (r *roleService) FindById(ctx *gin.Context) {
	var id = ctx.Param("id")

	value, err := strconv.ParseUint(id, 10, 16)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : err,
		})
		return
	}

	result, err := r.repository.FindById(value)
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