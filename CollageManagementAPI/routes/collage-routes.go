package routes

import (
	"gorm-test/controllers"

	"github.com/gin-gonic/gin"
)

// var (
// 	UserInter           dao.UserRepo                    = dao.UserRepo{}
// 	UserService         service.UserService             = service.New(UserInter)
// 	ControllerInterFace controllers.ControllerInterFace = controllers.ControllerNew(UserService)
// )

func SetupRouter() *gin.Engine {

	r := gin.Default()

	r.POST("/users", controllers.CreateUser)
	r.GET("/users", controllers.GetUsers)
	r.GET("/users/:id", controllers.GetUser)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	grp1 := r.Group("/collage-api")
	{

		grp1.GET("collage", controllers.GetAllCollagesCon)
		grp1.POST("collage", controllers.CreateCollageCon)
		grp1.GET("collage/:collageid", controllers.GetCollageByIDCon)
		grp1.PUT("collage/:collageid", controllers.UpdateCollageCon)
		grp1.DELETE("collage/:collageid", controllers.DeleteCollageCon)

		grp1.GET("department", controllers.GetAllDepartments)
		grp1.POST("department", controllers.CreateDepartment)
		grp1.GET("department/:department_id", controllers.GetDepartmentByID)
		grp1.PUT("department/:department_id", controllers.UpadateDepartment)
		grp1.DELETE("department/:department_id", controllers.DeleteDepartmentByID)

		grp1.GET("staff", controllers.GetAllStaff)
		grp1.POST("staff", controllers.CreateStaff)
		grp1.GET("staff/:staff_id", controllers.GetStaffByID)
		grp1.PUT("staff/:staff_id", controllers.UpadateStaff)
		grp1.DELETE("staff/:staff_id", controllers.DeleteStaffByID)

		grp1.GET("student", controllers.GetAllStudents)
		grp1.POST("student", controllers.CreateStudent)
		grp1.GET("student/:student_id", controllers.GetStudentByID)
		grp1.PUT("student/:student_id", controllers.UpadateStudent)
		grp1.DELETE("student/:student_id", controllers.DeleteStudentByID)

	}

	return r
}
