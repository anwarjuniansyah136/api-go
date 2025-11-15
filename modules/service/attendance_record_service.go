package service

import (
	"api/modules/model"
	"api/modules/repository"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AttendanceRecordService interface {
	Create(ctx *gin.Context)
	GetAllAttendanceRecord(ctx *gin.Context)
}

type attendanceRecordService struct {
	repository repository.AttendanceRecordRepository
}

func NewAttendanceRecordService(db *gorm.DB) AttendanceRecordService {
	return &attendanceRecordService{
		repository: repository.NewAttendanceRecordRepository(db),
	}
}

func (a *attendanceRecordService) Create(ctx *gin.Context) {
	var attendanceRecordRequest model.AttendanceRecordCreateRequest
	
	if err := ctx.ShouldBindJSON(&attendanceRecordRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : err,
		})
		return
	}

	layout := "2006-01-02 15:04:05"
	checkinTimeParse, err := time.ParseInLocation(layout, attendanceRecordRequest.CheckinTime, time.Local)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : "invalid field format checkin time",
		})
		return
	}

	verifiedAtParse, err := time.ParseInLocation(layout, attendanceRecordRequest.VerifiedAt, time.Local)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : "invalid field format verified at",
		})
		return
	}

	attendaceRecord := model.AttendanceRecord{
		SessionID: attendanceRecordRequest.SessionID,
		StudentID: attendanceRecordRequest.StudentID,
		CheckinTime: checkinTimeParse,
		Latitude: attendanceRecordRequest.Latitude,
		Longitude: attendanceRecordRequest.Longitude,
		SelfieURL: attendanceRecordRequest.SelfieURL,
		DistanceFrom: attendanceRecordRequest.DistanceFrom,
		SchoolID: attendanceRecordRequest.SchoolID,
		VerifiedAt: verifiedAtParse,
	}

	result, err := a.repository.Save(attendaceRecord)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error" : err,
		})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (a *attendanceRecordService) GetAllAttendanceRecord(ctx *gin.Context) {
	result, err := a.repository.FindAll()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error" : err,
		})
		return
	}

	ctx.JSON(http.StatusOK, result)
}