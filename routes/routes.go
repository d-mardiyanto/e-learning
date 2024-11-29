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
	user := router.Group("/api/users")
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

		// Kelas
		protected.POST("/kelas", controllers.CreateKelas)
		protected.PUT("/kelas/update/:id", controllers.UpdateKelas)
		protected.DELETE("/kelas/delete:/id", controllers.DeleteKelas)
		protected.GET("/kelas/show", controllers.ShowKelas)
		protected.GET("/kelas/show/:id", controllers.GetKelas)

		// Prodi
		protected.POST("/prodi", controllers.CreateProdi)
		protected.PUT("/prodi/update/:id", controllers.UpdateProdi)
		protected.DELETE("/prodi/delete:/id", controllers.DeleteProdi)
		protected.GET("/prodi/show", controllers.ShowProdi)
		protected.GET("/prodi/show/:id", controllers.GetProdi)

		// Courses
		protected.POST("/courses", controllers.CreateCourse)
		protected.PUT("/courses/update/:id", controllers.UpdateCourse)
		protected.DELETE("/courses/delete:/id", controllers.DeleteCourse)
		protected.POST("/courses/upload-files", controllers.UploadCourseFile)
		protected.GET("/courses/show", controllers.ShowCourses)
		protected.GET("/courses/show/:id", controllers.GetCourse)

	}
}
