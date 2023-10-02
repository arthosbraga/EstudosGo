// controllers/books.go

package booksController

import (
	decode "example/web-service-gin/Decode"
	encode "example/web-service-gin/Encode"
	services "example/web-service-gin/Services"
	"net/http"
)

type BookingController interface {
}

type bookinController struct {
	bookedServices services.BookingService
}

func New(bookedServices services.BookingService) *bookinController {
	return &bookinController{
		bookedServices: bookedServices,
	}
}

func (p *bookinController) GetProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := decode.DecodeStringIDFromURI(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, err := p.bookedServices.GetByID(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	encode.WriteJsonResponse(w, product, http.StatusOK)
}
