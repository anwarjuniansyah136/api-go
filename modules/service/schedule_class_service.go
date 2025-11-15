package service

import (
	"api/modules/model"
	"api/modules/repository"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ScheduleClassService interface {
	Create(ctx *gin.Context)
	GetAllScheduleClass(ctx *gin.Context)
	FindById(ctx *gin.Context)
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
			"error": err,
		})
		return
	}

	layout := "2006-01-02 15:04:05"
	startTimeParse, err := time.ParseInLocation(layout, scheduleClassRequest.StartTime, time.Local)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid format start time",
		})
		return
	}

	endTimeParse, err := time.ParseInLocation(layout, scheduleClassRequest.EndTime, time.Local)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid format end timr",
		})
		return
	}

	scheduleClass := model.ScheduleClass{
		SchoolID:  scheduleClassRequest.SchoolID,
		ClassName: scheduleClassRequest.ClassName,
		SubjectID: scheduleClassRequest.SchoolID,
		TeacherID: scheduleClassRequest.TeacherID,
		StartTime: startTimeParse,
		EndTime:   endTimeParse,
		RoomID:    scheduleClassRequest.RoomID,
	}

	result, err := s.repository.Save(scheduleClass)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (s *scheduleClassService) GetAllScheduleClass(ctx *gin.Context) {
	result, err := s.repository.FindAll()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (s *scheduleClassService) FindById(ctx *gin.Context) {
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
