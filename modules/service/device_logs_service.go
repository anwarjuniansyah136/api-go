package service

import (
	"api/modules/model"
	"api/modules/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DeviceLogsService interface {
	Create(ctx *gin.Context)
	GetAllDevice(ctx *gin.Context)
	FindById(ctx *gin.Context)
}

type devivceLogsService struct {
	repository repository.DeviceLogsRepository
}

func NewDeviceLogService(db *gorm.DB) DeviceLogsService {
	return &devivceLogsService{
		repository: repository.NewDeviceLogsRepository(db),
	}
}

func (d *devivceLogsService) Create(ctx *gin.Context) {
	var deviceRequest model.DeviceLogCreateRequest

	if err := ctx.ShouldBindJSON(&deviceRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	device := model.DeviceLog{
		UserID:    deviceRequest.UserID,
		DeviceID:  deviceRequest.DeviceID,
		Platform:  deviceRequest.Platform,
		IPAddress: deviceRequest.IPAddress,
	}

	result, err := d.repository.Save(device)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (d *devivceLogsService) GetAllDevice(ctx *gin.Context) {
	result, err := d.repository.FindAll()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (d *devivceLogsService) FindById(ctx *gin.Context) {
	var id = ctx.Param("id")

	value, err := strconv.ParseUint(id, 10, 16)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : err,
		})
		return
	}

	result, err := d.repository.FindById(value)
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