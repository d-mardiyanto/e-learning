package routes

import (
	"e-learning/controllers"
	"e-learning/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.Use(middleware.APIKeyMiddleware())

	// Welcome route
	router.GET("/", controllers.WelcomeMessage)

	// Public routes
	user := router.Group("/api/auth")
	{
		user.POST("/register", controllers.RegisterUser)
		user.POST("/login", controllers.LoginUser)
	}
	// Protected routes
	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		// User
		protected.GET("/profile", controllers.GetUserProfile)

		// Classes
		protected.POST("/classes", controllers.CreateClasses)
		protected.PUT("/classes/update/:id", controllers.UpdateClasses)
		protected.DELETE("/classes/delete/:id", controllers.DeleteClasses)
		protected.GET("/classes/show", controllers.ShowClasses)
		protected.GET("/classes/show/:id", controllers.GetClasses)

		// Prodi
		protected.POST("/prody", controllers.CreateStudyProgram)
		protected.PUT("/prody/update/:id", controllers.UpdateStudyProgram)
		protected.DELETE("/prody/delete/:id", controllers.DeleteStudyProgram)
		protected.GET("/prody/show", controllers.ShowStudyProgram)
		protected.GET("/prody/show/:id", controllers.GetStudyProgram)

		// Instructor
		protected.POST("/instructor", controllers.CreateInstructor)
		protected.PUT("/instructor/update/:id", controllers.UpdateInstructor)
		protected.DELETE("/instructor/delete/:id", controllers.DeleteInstructor)
		protected.GET("/instructor/show", controllers.ShowInstructors)
		protected.GET("/instructor/show/:id", controllers.GetInstructor)

		// Students
		protected.POST("/student", controllers.CreateStudent)
		protected.PUT("/student/update/:id", controllers.UpdateStudent)
		protected.DELETE("/student/delete/:id", controllers.DeleteStudent)
		protected.GET("/student/show", controllers.ShowStudents)
		protected.GET("/student/show/:id", controllers.GetStudent)

		// Courses
		protected.POST("/courses", controllers.CreateCourse)
		protected.PUT("/courses/update/:id", controllers.UpdateCourse)
		protected.DELETE("/courses/delete/:id", controllers.DeleteCourse)
		protected.POST("/courses/upload-files", controllers.UploadCourseFile)
		protected.GET("/courses/show", controllers.ShowCourses)
		protected.GET("/courses/show/:id", controllers.GetCourse)

	}
}
