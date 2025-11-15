package service

import (
	"api/modules/model"
	"api/modules/repository"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AttendanceSessionService interface {
	Create(ctx *gin.Context)
	GetAllAttendanceSession(ctx *gin.Context)
}

type attendanceSessionService struct {
	repository repository.AttendanceSessionRepository
}

func NewAttendanceSessionService(db *gorm.DB) AttendanceSessionService {
	return &attendanceSessionService{
		repository: repository.NewAttendanceSessionRepository(db),
	}
}

func (a *attendanceSessionService) Create(ctx *gin.Context) {
	var attendanceSessionRequest model.AttendanceSessionCreateRequest
	
	if err := ctx.ShouldBindJSON(&attendanceSessionRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : err,
		})
		return
	}

	layout := "2006-01-02 15:04:05"
	startTimeParse, err := time.ParseInLocation(layout, attendanceSessionRequest.StartTime, time.Local)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : "invalid format start time",
		})
		return
	}

	endTimeParse, err := time.ParseInLocation(layout, attendanceSessionRequest.EndTime, time.Local)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : "invalid format end time",
		})
		return
	}

	attendaceSession := model.AttendanceSession{
		SchoolID: attendanceSessionRequest.SchoolID,
		Name: attendanceSessionRequest.Name,
		StartTime: startTimeParse,
		EndTime: endTimeParse,
	}

	result, err := a.repository.Save(attendaceSession)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error" : err,
		})
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (a *attendanceSessionService) GetAllAttendanceSession(ctx *gin.Context) {
	result, err := a.repository.FindAll()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error" : err,
		})
		return
	}

	ctx.JSON(http.StatusOK, result)
}