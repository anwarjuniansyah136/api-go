package service

import (
	"api/modules/model"
	"api/modules/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SubjectService interface {
	Create(ctx *gin.Context)
	GetAllSubject(ctx *gin.Context)
}

type subjectService struct {
	repository repository.SubjectRepository
}

func NewSubjectService(db *gorm.DB) SubjectService {
	return &subjectService{
		repository: repository.NewSubjectRepository(db),
	}
}

func (s *subjectService) Create(ctx *gin.Context) {
	var subjectRequest model.SubjectCreateRequest

	if err := ctx.ShouldBindJSON(&subjectRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : err,
		})
		return
	}

	subject := model.Subject{
		SubjectCode: subjectRequest.SubjectCode,
		SubjectName: subjectRequest.SubjectName,
		IsActive: subjectRequest.IsActive,
	}

	result, err := s.repository.Save(subject)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error" : err,
		})
		return
	}

	ctx.JSON(http.StatusOK,result)
}

func (s *subjectService) GetAllSubject(ctx *gin.Context) {
	result, err := s.repository.FindAll()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error" : err,
		})
		return
	}
	ctx.JSON(http.StatusOK, result)
}