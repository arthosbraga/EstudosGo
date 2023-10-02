package bookrepositories

import (
	"context"
	models "example/web-service-gin/Repository/Models"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type BookingRepository interface {
	GetByID(_ context.Context, key map[string]*dynamodb.AttributeValue) (*models.Booking, error)
	Search(_ context.Context, id string) ([]*models.Booking, error)
	Create(_ context.Context, booking *models.Booking) error
	Update(_ context.Context, BookingToUpdate *models.Booking) error
	Delete(_ context.Context, id string) error
}

type bookingRepository struct {
	db        *dynamodb.DynamoDB
	tableName string
}

func New(db *dynamodb.DynamoDB, tableName string) *bookingRepository {
	return &bookingRepository{
		db:        db,
		tableName: tableName,
	}
}

func (p *bookingRepository) GetByID(_ context.Context, key map[string]*dynamodb.AttributeValue) (*models.Booking, error) {

	resp, err := p.db.GetItem(&dynamodb.GetItemInput{Key: key, TableName: aws.String(p.tableName)})

	if err != nil {
		return nil, err
	}

	var booking *models.Booking

	err = dynamodbattribute.UnmarshalMap(resp.Item, &booking)

	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	return booking, nil
}

func (p *bookingRepository) Search(_ context.Context, productType string) ([]*models.Booking, error) {
	panic("unimplemented")
}

func (p *bookingRepository) Create(_ context.Context, product *models.Booking) error {
	panic("unimplemented")
}

func (p *bookingRepository) Update(_ context.Context, productToUpdate *models.Booking) error {
	panic("unimplemented")
}

func (p *bookingRepository) Delete(_ context.Context, id string) error {
	panic("unimplemented")
}
