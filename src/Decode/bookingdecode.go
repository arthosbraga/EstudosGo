package decode

import (
	"encoding/json"
	"errors"
	models "example/web-service-gin/Repository/Models"

	"net/http"

	"github.com/go-chi/chi/v5"
)

func DecodeStringIDFromURI(r *http.Request) (string, error) {
	id := chi.URLParam(r, "shippingOptionId")
	if id == "" {
		return "", errors.New("empty_id_error")
	}

	return id, nil
}

func DecodeTypeQueryString(r *http.Request) string {
	return r.URL.Query().Get("type")
}

func DecodeBookingFromBody(r *http.Request) (*models.Booking, error) {
	createProduct := &models.Booking{}
	err := json.NewDecoder(r.Body).Decode(&createProduct)
	if err != nil {
		return nil, err
	}

	return createProduct, nil
}
