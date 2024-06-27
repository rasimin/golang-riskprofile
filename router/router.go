// package router mengatur rute untuk aplikasi
package router

import (
	"project-riskprofile/handler"
	"project-riskprofile/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter menginisialisasi dan mengatur rute untuk aplikasi
func SetupRouter(r *gin.Engine, userHandler handler.IUserHandler, submissionHandlerHandler handler.IsubmissionHandler) {
	// Mengatur endpoint publik untuk pengguna
	usersPublicEndpoint := r.Group("/users")
	// Rute untuk mendapatkan pengguna berdasarkan ID
	usersPublicEndpoint.GET("/:id", userHandler.GetUser)
	usersPublicEndpoint.GET("/email/:email", userHandler.GetUserByEmail)
	// Rute untuk mendapatkan semua pengguna
	usersPublicEndpoint.GET("", userHandler.GetAllUsers)
	usersPublicEndpoint.GET("/", userHandler.GetAllUsers)

	// Mengatur endpoint privat untuk pengguna dengan middleware autentikasi
	usersPrivateEndpoint := r.Group("/users")
	// Menambahkan middleware autentikasi untuk endpoint privat
	usersPrivateEndpoint.Use(middleware.AuthMiddleware())
	// Rute untuk membuat pengguna baru
	usersPrivateEndpoint.POST("", userHandler.CreateUser)
	usersPrivateEndpoint.POST("/", userHandler.CreateUser)
	// Rute untuk memperbarui pengguna berdasarkan ID
	usersPrivateEndpoint.PUT("/:id", userHandler.UpdateUser)
	// Rute untuk menghapus pengguna berdasarkan ID
	usersPrivateEndpoint.DELETE("/:id", userHandler.DeleteUser)

	subsPrivateEndpoint := r.Group("/submissions")
	// Menambahkan middleware autentikasi untuk endpoint privat
	subsPrivateEndpoint.Use(middleware.AuthMiddleware())
	subsPrivateEndpoint.POST("", submissionHandlerHandler.CreateSubmi)
	// subsPrivateEndpoint.POST("/", submissionHandlerHandler.GetSubmi)

	subsPrivateEndpoint.GET("/:id", submissionHandlerHandler.GetSubmi)
	subsPrivateEndpoint.GET("/", submissionHandlerHandler.GetAllSubmi)
	subsPrivateEndpoint.DELETE("/:id", submissionHandlerHandler.DeleteSubmi)

	// subsPrivateEndpoint.PUT("/:id", submissionHandlerHandler.UpdateSubmi)
	// subsPrivateEndpoint.PUT("/:id", submissionHandlerHandler.)
}
