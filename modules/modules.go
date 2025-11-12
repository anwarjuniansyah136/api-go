package modules

import (
	"api/helper"
	"api/modules/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Versions interface {
	Run()
}

type versions struct {
	configEnv  helper.Config
	mainServer *gin.Engine
	db         *gorm.DB
}

func NewVersions(configEnv helper.Config, mainServer *gin.Engine, db *gorm.DB) Versions {
	return &versions{
		configEnv:  configEnv,
		mainServer: mainServer,
		db:         db,
	}
}


func (v *versions) Run() {
	apiRoutes := v.mainServer.Group("/api")
	db_v1 := helper.OpenDB(v.configEnv.DB, v.configEnv.SCHEMA, "v1")

	attendanceRecordController := controller.NewAttendanceRecordController(apiRoutes, db_v1, "v1")
	attendanceSessionController := controller.NewAttendanceSessionController(apiRoutes, db_v1, "v1")
	deviceLogsController := controller.NewDeviceLogController(apiRoutes, db_v1, "v1")
	roleController := controller.NewRoleController(apiRoutes, db_v1, "v1")
	roomController := controller.NewRoomController(apiRoutes, db_v1, "v1")
	scheduleClassController := controller.NewScheduleClassController(apiRoutes, db_v1, "v1")
	schoolController := controller.NewSchoolController(apiRoutes, db_v1, "v1")
	studentsController := controller.NewStudentsController(apiRoutes, db_v1, "v1")
	subjectController := controller.NewSubjectController(apiRoutes, db_v1, "v1")
	teacherController := controller.NewTeacherController(apiRoutes, db_v1, "v1")
	userController := controller.NewUserController(apiRoutes, db_v1, "v1")

	attendanceRecordController.Init()
	attendanceSessionController.Init()
	deviceLogsController.Init()
	roleController.Init()
	roomController.Init()
	scheduleClassController.Init()
	schoolController.Init()
	studentsController.Init()
	subjectController.Init()
	teacherController.Init()
	userController.Init()
}