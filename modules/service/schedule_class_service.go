package service

import (
	"api/modules/model"
	"api/modules/repository"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ScheduleClassService interface {
	Create(ctx *gin.Context)
	GetAllScheduleClass(ctx *gin.Context)
}

type scheduleClassService struct {
	repository repository.ScheduleClassRepository
}

func NewScheduleClassService(db *gorm.DB) ScheduleClassService {
	return &scheduleClassService{
		repository: repository.NewScheduleRepository(db),
	}
}

func (s *scheduleClassService) Create(ctx *gin.Context) {
	var scheduleClassRequest model.ScheduleClassCreateRequest

	if err := ctx.ShouldBindJSON(&scheduleClassRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : err,
		})
		return
	}

	layout := "2006-01-02 15:04:05"
	startTimeParse, err := time.ParseInLocation(layout, scheduleClassRequest.StartTime, time.Local)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : "invalid format start time",
		})
		return
	} 

	endTimeParse, err := time.ParseInLocation(layout, scheduleClassRequest.EndTime, time.Local)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : "invalid format end timr",
		})
		return
	}

	scheduleClass := model.ScheduleClass{
		SchoolID: scheduleClassRequest.SchoolID,
		ClassName: scheduleClassRequest.ClassName,
		SubjectID: scheduleClassRequest.SchoolID,
		TeacherID: scheduleClassRequest.TeacherID,
		StartTime: startTimeParse,
		EndTime: endTimeParse,
		RoomID: scheduleClassRequest.RoomID,
	}

	result, err := s.repository.Save(scheduleClass)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error" : err,
		})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (s *scheduleClassService) GetAllScheduleClass(ctx *gin.Context) {
	result, err := s.repository.FindAll()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error" : err,
		})
		return
	}

	ctx.JSON(http.StatusOK, result)
}