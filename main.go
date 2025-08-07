package main

import (
	"exams/internal/auth"
	"exams/internal/config"
	"exams/internal/db"
	"exams/internal/modules/exams"
	examschoice "exams/internal/modules/examsChoice"
	"exams/internal/modules/examsQuestions"
	examssession "exams/internal/modules/examsSession"
	examssessionexam "exams/internal/modules/examsSessionExam"
	examssubmission "exams/internal/modules/examsSubmission"
	examssubmissionanswer "exams/internal/modules/examsSubmissionAnswer"
	"exams/internal/modules/students"
	"exams/internal/modules/subjects"
	"exams/internal/modules/teachers"
	"exams/internal/modules/users"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	config.LoadConfig()
	dbConn := db.NewPostgresConnection()

	users.RegisterRoutes(app, dbConn)
	teachers.RegisterRoutes(app, dbConn)
	auth.RegisterRoutes(app, dbConn)
	exams.RegisterRoutes(app, dbConn)
	examsQuestions.RegisterRoutes(app, dbConn)
	examschoice.RegisterRoutes(app, dbConn)
	examssession.RegisterRoutes(app, dbConn)
	examssessionexam.RegisterRoutes(app, dbConn)
	examssubmission.RegisterRoutes(app, dbConn)
	examssubmissionanswer.RegisterRoutes(app, dbConn)
	students.RegisterRoutes(app, dbConn)
	subjects.RegisterRoutes(app, dbConn)
	app.Listen(":8000")
}
