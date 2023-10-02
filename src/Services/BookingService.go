package bookservices

import (
	"context"
	bookrepositories "example/web-service-gin/Repository"
	models "example/web-service-gin/Repository/Models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type BookingService interface {
	GetByID(ctx context.Context, id string) (*models.Booking, error)
	Search(ctx context.Context, bookingType string) ([]*models.Booking, error)
	Create(ctx context.Context, booking *models.Booking) error
	Update(ctx context.Context, booking *models.Booking) error
	Delete(ctx context.Context, id string) error
}

type bookingService struct {
	bookingRepository bookrepositories.BookingRepository
}

func New(bookingRepository bookrepositories.BookingRepository) *bookingService {
	return &bookingService{
		bookingRepository: bookingRepository,
	}
}

func (p *bookingService) GetByID(ctx context.Context, id string) (*models.Booking, error) {

	Key := map[string]*dynamodb.AttributeValue{
		"shippingOptionId": {
			S: aws.String(id),
		},
	}

	return p.bookingRepository.GetByID(ctx, Key)
}
func (p *bookingService) Search(ctx context.Context, bookingType string) ([]*models.Booking, error) {
	return p.bookingRepository.Search(ctx, bookingType)
}
func (p *bookingService) Create(ctx context.Context, booking *models.Booking) error {
	return p.bookingRepository.Create(ctx, booking)
}
func (p *bookingService) Update(ctx context.Context, booking *models.Booking) error {
	return p.bookingRepository.Update(ctx, booking)
}
func (p *bookingService) Delete(ctx context.Context, id string) error {
	return p.bookingRepository.Delete(ctx, id)
}
