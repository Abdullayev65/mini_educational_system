package app

import "github.com/gin-gonic/gin"

func (a *App) initRouters() *gin.Engine {
	router := gin.Default()

	user := router.Group("/user")
	{
		user.POST("/", a.handler.CreateUser)
		user.POST("/sign-in", a.handler.SignIn)
		user.GET("/me", a.MW.UserIDFromToken, a.handler.UserMe)
		user.PUT("/", a.MW.UserIDFromToken, a.handler.PatchUser)
	}

	attendance := router.Group("/attendance")
	{
		attendance.POST("/",
			a.MW.UserIDFromToken, a.handler.AttendanceAdd)
		attendance.GET("/:userID",
			a.MW.SetIntFromParam("userID"), a.handler.AttendanceByUserID)
	}

	admin := router.Group("/admin", a.MW.UserIDFromToken, a.MW.CheckAdmin)
	{
		department := admin.Group("/department")
		{
			department.POST("/", a.handler.DepAdd)
			department.PUT("/:id",
				a.MW.SetIntFromParam("id"), a.handler.DepUpdate)
			department.GET("/all", a.handler.DepAll)
			department.DELETE("/:id",
				a.MW.SetIntFromParam("id"), a.handler.DepDelete)
		}

		subject := admin.Group("/subject")
		{
			subject.POST("/", a.handler.SubjectAdd)
			subject.GET("/all", a.handler.SubjectAll)
			subject.PUT("/:id",
				a.MW.SetIntFromParam("id"), a.handler.SubjectUpdate)
			subject.DELETE("/:id",
				a.MW.SetIntFromParam("id"), a.handler.SubjectDelete)
		}

		level := admin.Group("/level")
		{
			level.POST("/", a.handler.LvlAdd)
			level.GET("/all", a.handler.LvlAll)
			level.PUT("/:id",
				a.MW.SetIntFromParam("id"), a.handler.LvlUpdate)
			level.DELETE("/:id",
				a.MW.SetIntFromParam("id"), a.handler.LvlDelete)
		}

		users := admin.Group("users")
		{
			users.GET("/by-department/:departmentID",
				a.MW.SetIntFromParam("departmentID"), a.handler.UsersByDepartment)
		}

	}

	return router
}
