package services

import (
	// "net/http"

	"fmt"

	"github.com/Igor-Kreshchenko/shipment-test/models"
	"github.com/Igor-Kreshchenko/shipment-test/repositories"

	"github.com/go-playground/validator"
)

type ShipmentRequest struct {
	SenderName           string  `json:"senderName"`
	SenderEmail          string  `json:"senderEmail"`
	SenderAddress        string  `json:"senderAddress"`
	SenderCountryCode    string  `json:"senderCountryCode"`
	RecipientName        string  `json:"recipientName"`
	RecipientEmail       string  `json:"recipientEmail"`
	RecipientAddress     string  `json:"recipientAddress"`
	RecipientCountryCode string  `json:"recipientCountryCode"`
	Weight               float64 `json:"weight"`
}

type ApiResponse struct {
	Region string
}

type ShipmentService interface {
	GetAllShipments() ([]models.Shipment, error)
	CreateNewShipment(shipment *ShipmentRequest) (*models.Shipment, error)
	GetShipmentByID(id uint) (*models.Shipment, error)
}

type shipmentService struct {
	shipmentRepository repositories.ShipmentRepository
}

func NewShipmentService(shipmentRepo repositories.ShipmentRepository) ShipmentService {
	return &shipmentService{shipmentRepository: shipmentRepo}
}

func (s *shipmentService) GetAllShipments() ([]models.Shipment, error) {
	res, err := s.shipmentRepository.GetAllShipments()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *shipmentService) CreateNewShipment(shipmentReq *ShipmentRequest) (*models.Shipment, error) {
	// weightClass := getWeightClass(shipmentReq.Weight)
	// price := calculatePrice(weightClass)

	shipment := models.Shipment{
		SenderName:           shipmentReq.SenderName,
		SenderEmail:          shipmentReq.SenderEmail,
		SenderAddress:        shipmentReq.SenderAddress,
		SenderCountryCode:    shipmentReq.SenderCountryCode,
		RecipientName:        shipmentReq.RecipientName,
		RecipientEmail:       shipmentReq.RecipientEmail,
		RecipientAddress:     shipmentReq.RecipientAddress,
		RecipientCountryCode: shipmentReq.RecipientAddress,
		Weight:               shipmentReq.Weight,
		Price:                100,
	}

	err := validate(&shipment)
	if err != nil {
		return nil, err
	}

	res, err := s.shipmentRepository.CreateNewShipment(&shipment)
	if err != nil {
		return nil, err
	}

	return res, err
}

func (s *shipmentService) GetShipmentByID(id uint) (*models.Shipment, error) {
	shipment, err := s.shipmentRepository.GetShipmentByID(id)
	if err != nil {
		return nil, err
	}

	return shipment, nil
}

func calculatePrice(region string, weightClass float64) float64 {
	index := 1

	price := weightClass * float64(index)

	return price
}

func getWeightClass(weight float64) float64 {
	weightClass := 0

	if weight > 0 && weight <= 10 {
		weightClass = 100
	} else if weight > 10 && weight <= 25 {
		weightClass = 300
	} else if weight > 25 && weight <= 50 {
		weightClass = 500
	} else if weight > 50 && weight <= 1000 {
		weightClass = 2000
	}

	return float64(weightClass)
}

// func getIsEU(code string) (bool, err) {
// 	var country *ApiResponse

// 	url := fmt.Sprintf("https://restcountries.com/v2/alpha/%v", code)

// 	res, err := http.Get(url)
// 	if err != nil {
// 		return false, err
// 	}

// }

func validate(s *models.Shipment) error {
	var validate *validator.Validate = validator.New()

	err := validate.Struct(s)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return err
		}

		fmt.Println("------ List of tag fields with error ---------")

		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.StructField())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println("---------------")
		}
		return err
	}

	return nil
}
