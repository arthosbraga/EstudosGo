package main

import (
	booksController "example/web-service-gin/Controllers"
	bookrepositories "example/web-service-gin/Repository"
	services "example/web-service-gin/Services"
	"net/http"

	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Profile:           "mpms-dev",
		Config: aws.Config{
			Region: aws.String("eu-west-3"),
		},
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	bookingRepository := bookrepositories.New(svc, "mpms-booking-dev")
	bookingService := services.New(bookingRepository)
	booksController := booksController.New(bookingService)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/booking/{shippingOptionId}", booksController.GetProductByIDHandler)

	http.ListenAndServe(":"+port, r)

}
