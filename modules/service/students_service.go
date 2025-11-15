package service

import (
	"api/modules/model"
	"api/modules/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type StudentsService interface {
	Create(ctx *gin.Context)
	GetAllStudents(ctx *gin.Context)
	FindById(ctx *gin.Context)
}

type studentsService struct {
	repository repository.StudentRepository
}

func NewStudentsService(db *gorm.DB) StudentsService {
	return &studentsService{
		repository: repository.NewStudentRepository(db),
	}
}

func (s *studentsService) Create(ctx *gin.Context) {
	var studentRequest model.StudentCreateRequest

	if err := ctx.ShouldBindJSON(&studentRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	student := model.Student{
		UserID:       studentRequest.UserID,
		NISN:         studentRequest.NISN,
		RoomID:       studentRequest.RoomID,
		AcademicYear: studentRequest.AcademicYear,
	}

	result, err := s.repository.Save(student)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (s *studentsService) GetAllStudents(ctx *gin.Context) {
	result, err := s.repository.FindAll()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (s *studentsService) FindById(ctx *gin.Context) {
	var id = ctx.Param("id")

	value, err := strconv.ParseUint(id, 10, 16)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : err,
		})
		return
	}

	result, err := s.repository.FindById(value)
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