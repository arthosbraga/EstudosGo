package models

import "time"

type Booking struct {
	ShippingOptionID  string        `json:"shippingOptionId"`
	AvailableServices []interface{} `json:"availableServices"`
	BookedServices    []struct {
		ID         string `json:"id"`
		BookLineID string `json:"bookLineId"`
	} `json:"bookedServices"`
	BookID           string    `json:"bookId"`
	DeliveryDateFrom time.Time `json:"deliveryDateFrom"`
	DeliveryDateTo   time.Time `json:"deliveryDateTo"`
	DeliveryMethod   string    `json:"deliveryMethod"`
	DispatchDate     time.Time `json:"dispatchDate"`
	JourneyID        string    `json:"journeyId"`
	Mot              string    `json:"mot"`
	Motive           string    `json:"motive"`
	Responsibility   string    `json:"responsibility"`
	Outbound         bool      `json:"outbound"`
	ProductInfo      []struct {
		UnitWeight                 string `json:"unitWeight"`
		UnitWidth                  string `json:"unitWidth"`
		Item                       string `json:"item"`
		Quantity                   string `json:"quantity"`
		UnitVolume                 string `json:"unitVolume"`
		UnitWeightUnitOfMeasure    string `json:"unitWeightUnitOfMeasure"`
		BookLineID                 string `json:"bookLineId"`
		UnitHeight                 string `json:"unitHeight"`
		UnitLength                 string `json:"unitLength"`
		QuantityUnitOfMeasure      string `json:"quantityUnitOfMeasure"`
		ItemDescription            string `json:"itemDescription"`
		UnitDimensionUnitOfMeasure string `json:"unitDimensionUnitOfMeasure"`
		UnitVolumeUnitOfMeasure    string `json:"unitVolumeUnitOfMeasure"`
	} `json:"productInfo"`
	ScheduleCounter    string `json:"scheduleCounter"`
	SelectionID        string `json:"selectionId"`
	SenderLocation     string `json:"senderLocation"`
	ServiceLevelDetail string `json:"serviceLevelDetail"`
	ShipmentIdentifier string `json:"shipmentIdentifier"`
	ShipmentStatus     string `json:"shipmentStatus"`
	TimeSlotID         string `json:"timeSlotId"`
	TTL                string `json:"TTL"`
	DestinationAddress struct {
		AdditionalAddressInformation1 string `json:"additionalAddressInformation1"`
		AdditionalAddressInformation2 string `json:"additionalAddressInformation2"`
		AddressType                   string `json:"addressType"`
		City                          string `json:"city"`
		ContactName                   string `json:"contactName"`
		Country                       string `json:"country"`
		Email                         string `json:"email"`
		ExternalAddressIdentifier     string `json:"externalAddressIdentifier"`
		Fax                           string `json:"fax"`
		HouseNumber                   string `json:"houseNumber"`
		Name                          string `json:"name"`
		Phone                         string `json:"phone"`
		PostalCode                    string `json:"postalCode"`
		Residential                   bool   `json:"residential"`
		State                         string `json:"state"`
		Street                        string `json:"street"`
		VatNumber                     string `json:"vatNumber"`
	} `json:"destinationAddress"`
}
