package service

import (
	"api/modules/model"
	"api/modules/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SchoolService interface {
	Create(ctx *gin.Context)
	GetAllUser(ctx *gin.Context)
}

type schoolService struct {
	repository repository.SchoolRepository
}

func NewSchoolServie(db *gorm.DB) SchoolService {
	return &schoolService{
		repository: repository.NewSchoolRepository(db),
	}
}

func (s *schoolService) Create(ctx *gin.Context) {
	var schoolRequest model.SchoolCreateRequest

	if err := ctx.ShouldBindJSON(&schoolRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : err,
		})
		return
	}

	school := model.School{
		SchoolName: schoolRequest.SchoolName,
		Address: schoolRequest.Address,
		Latitude: schoolRequest.Latitude,
		Longitude: schoolRequest.Longitude,
		RadiusMeter: schoolRequest.RadiusMeter,
	}

	result, err := s.repository.Save(school)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error" : err,
		})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (s *schoolService) GetAllUser(ctx *gin.Context) {
	result, err := s.repository.FindAll()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error" : err,
		})
		return
	}
	ctx.JSON(http.StatusOK, result)
}