package service

import (
	"api/modules/model"
	"api/modules/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TeacherService interface {
	Create(ctx *gin.Context)
	GetAllTeacher(ctx *gin.Context)
	GetById(ctx *gin.Context)
}

type teacherService struct {
	repository repository.TeacherRepository
}

func NewTeacherService(db *gorm.DB) TeacherService {
	return &teacherService{
		repository: repository.NewTeacherRepository(db),
	}
}

func (t *teacherService) Create(ctx *gin.Context) {
	var teacherRequest model.TeacherCreateRequest

	if err := ctx.ShouldBindJSON(&teacherRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	teacher := model.Teacher{
		Name:      teacherRequest.Name,
		Address:   teacherRequest.Address,
		Age:       teacherRequest.Age,
		SubjectID: teacherRequest.SubjectID,
	}

	result, err := t.repository.Save(teacher)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (t *teacherService) GetAllTeacher(ctx *gin.Context) {
	result, err := t.repository.FindAll()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
	}
	ctx.JSON(http.StatusOK, result)
}

func (t *teacherService) GetById(ctx *gin.Context) {
	var id = ctx.Param("id")

	value, err := strconv.ParseUint(id, 10, 64);
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : "invalid parameter",
		})
		return
	}
	result, err := t.repository.FindByID(uint64(value));

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