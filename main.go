package main

import (
	"log"
	"project-riskprofile/handler"
	"project-riskprofile/repository/postgres_gorm"
	"project-riskprofile/router"
	"project-riskprofile/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// setup database connection
	//dsn := "postgresql://postgres:postgres@localhost:5432/postgres"
	// setup pgx connection
	//pgxPool, err := connectDB(dsn)
	//if err != nil {
	//	log.Fatalln(err)
	//}

	// setup gorm connectoin
	dsn := "postgresql://postgres:admin@localhost:5432/tugasSmester1"
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		log.Fatalln(err)
	}
	// setup service

	// slice db is disabled. uncomment to enabled
	// var mockUserDBInSlice []entity.User
	// _ = slice.NewUserRepository(mockUserDBInSlice)

	// uncomment to use postgres pgx
	// userRepo := postgres_pgx.NewUserRepository(pgxPool)

	// uncomment to use postgres gorm
	userRepo := postgres_gorm.NewUserRepository(gormDB)

	// service and handler declaration
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	submissionRepo := postgres_gorm.NewsubmissionRepository(gormDB)

	// service and handler declaration
	submissionService := service.NewSubmissionService(submissionRepo)
	submissionHandler := handler.NewsubmissionHandler(submissionService)

	// Routes
	router.SetupRouter(r, userHandler, submissionHandler)

	// score := 28

	// // Get the risk profile based on the score
	// profile := model.GetRiskProfile(score)

	// // Print the result
	// fmt.Printf("Total Score: %d\n", score)
	// fmt.Printf("Risk Profile: %s\n", profile.Category)
	// fmt.Printf("Definition: %s\n", profile.Definition)

	// Run the server
	log.Println("Running server on port 8080")
	r.Run(":8080")
}
